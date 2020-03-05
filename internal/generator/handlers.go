package generator

import (
	"container/list"
	"fmt"
	"io"
	"sort"
	"strings"
	"text/template"

	"github.com/kepkin/gorest/internal/spec/openapi3"
)

func (g *Generator) makeHandlers(wr io.Writer, sp openapi3.Spec) error {
	interfaceName := MakeIdentifier(sp.Info.Title)

	sortedPaths := make([]string, 0, len(sp.Paths))
	for k := range sp.Paths {
		sortedPaths = append(sortedPaths, k)
	}

	sort.Strings(sortedPaths)

	for _, pathKey := range sortedPaths {
		path := sp.Paths[pathKey]
		for _, method := range path.Methods() {
			if method != nil {
				root, err := g.makeRequestObject(wr, interfaceName, method)
				if err != nil {
					return err
				}
				if err := g.makeHandler(wr, interfaceName, method, root); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// From PathSpec generates openapi schema for creating a request struct
func (g *Generator) makeRequestObject(wr io.Writer, interfaceName string, method *openapi3.PathSpec) (openapi3.SchemaType, error) { //nolint:gocyclo,gocognit
	if _, err := fmt.Fprintf(wr, "\n// _%s_%s_Handler\n", interfaceName, method.OperationID); err != nil {
		return openapi3.SchemaType{}, err
	}

	// Make and fill root request schema
	request := openapi3.NewObjectSchema(method.OperationID + "Request")

	queryParams := openapi3.NewObjectSchema(method.OperationID + "RequestQuery")
	pathParams := openapi3.NewObjectSchema(method.OperationID + "RequestPath")
	cookieParams := openapi3.NewObjectSchema(method.OperationID + "RequestCookie")
	headerParams := openapi3.NewObjectSchema(method.OperationID + "RequestHeader")
	body := openapi3.NewObjectSchema(method.OperationID + "RequestBody")

	for _, param := range method.Parameters {
		addProperty := func(obj *openapi3.SchemaType, parameter openapi3.ParameterType) {
			obj.Properties[parameter.Name] = parameter.Schema
			if parameter.Required  {
				obj.Required = append(obj.Required, parameter.Name)
			}
		}

		switch param.In {
		case "query":
			addProperty(&queryParams, param)
		case "path":
			addProperty(&pathParams, param)
		case "cookie":
			addProperty(&cookieParams, param)
		case "header":
			addProperty(&headerParams, param)
		default:
			return request, fmt.Errorf("unknown param place: %s", param.In)
		}
	}

	if len(queryParams.Properties) > 0 {
		request.Properties["Query"] = &queryParams
	}
	if len(pathParams.Properties) > 0 {
		request.Properties["Path"] = &pathParams
	}
	if len(cookieParams.Properties) > 0 {
		request.Properties["Cookie"] = &cookieParams
	}
	if len(headerParams.Properties) > 0 {
		request.Properties["Header"] = &headerParams
	}

	if method.RequestBody != nil {
		for mimeType, content := range method.RequestBody.Content {
			switch mimeType {
			case "application/json":
				schema := *content.Schema
				schema.Type = openapi3.ObjectType
				body.Properties["JSON"] = &schema
			case "application/xml":
				body.Properties["XML"] = content.Schema
			case "multipart/form-data":
				schema := *content.Schema
				schema.Type = openapi3.ObjectType
				body.Properties["Form"] = &schema
			default:
				return request, fmt.Errorf("unsupported content type: %s", mimeType)
			}
		}
		// TODO(a.telyshev): Dirty hack
		body.Properties["Type"] = &openapi3.SchemaType{
			Name: "Type",
			Ref:  "#/components/schemas/ContentType",
		}
		request.Properties["Body"] = &body
	}

	// Determine definitions
	for _, obj := range []openapi3.SchemaType{
		request,
		//queryParams,
		//cookieParams,
		//pathParams,
		//headerParams,
		//body,
	} {
		d1, err := g.MakeTypeDefFromOpenAPIObject(obj, list.New())
		if err != nil {
			return openapi3.SchemaType{}, err
		}

		//Make struct
		definition, err := d1.BuildDefinition()
		if err != nil {
			return openapi3.SchemaType{}, err
		}
		wr.Write([]byte(definition))
	}

	return request, nil
}

var funcMap = template.FuncMap{
"ToUpper": strings.ToUpper,
"Title": strings.Title,
"MakeIdentifier": MakeIdentifier,
"MakeFormatIdentifier": MakeFormatIdentifier,
}

var ginGlobalTemplate = template.Must(template.New("handler").Funcs(funcMap).Parse(`
func ginGetCookie(c *gin.Context, param string) (string, bool) {
	cookie, err := c.Request.Cookie(param)
	if err == http.ErrNoCookie {
		return "", false
	}
	return cookie.Value, true
}

func __gin_get_parameter(c *gin.Context, in string, parameterName string) ([]string, bool, error) {
	if in == "cookie" {
		data, existed := ginGetCookie(c, parameterName)
		return []string{data}, existed, nil
	} else if in == "header" {
		data := c.Request.Header.Get(parameterName)
		return []string{data}, len(data) != 0, nil
	} else if in == "path" {
		data, existed := c.Params.Get(parameterName)
		return []string{data}, existed, nil
	} else if in == "query" {
		data, existed := c.Request.URL.Query()[parameterName]
		return data, existed, nil
	} else if in == "body" {
		data, existed := c.GetPostFormArray(parameterName)
		if existed == false && c.Request.MultipartForm != nil && c.Request.MultipartForm.File != nil { 
			fhs, ok := c.Request.MultipartForm.File[parameterName]
			if !ok {
				return []string{}, ok, nil
			}
			
			return []string{fhs[0].Filename}, ok, nil
		}

		return data, existed, nil
	}

	return []string{}, false, fmt.Errorf("Unsupported 'in': %v", in)
}
`))

//TODO: support for getting body from ginGenerator
var handlerTemplate = template.Must(template.New("handler").Funcs(funcMap).Parse(`
func (server {{ .InterfaceName }}Server) _{{ .InterfaceName }}_{{ .Path.OperationID }}_Handler(c *gin.Context) {
	c.Set(handlerNameKey, "{{ .Path.OperationID }}")
	
	var req {{ .Path.OperationID }}Request
	{{- range .Path.Parameters }}
		data{{ Title .In }}{{ MakeIdentifier .Name }}, _, err := __gin_get_parameter(c, "{{.In}}", "{{.Name}}")
		if err != nil {
			server.Srv.ProcessMakeRequestErrors(c, []FieldError{NewFieldError(In{{Title .In}}, "-", "", err)})
			return
  		}
		req.{{ Title .In }}.{{ MakeIdentifier .Name }}, err = {{ .Schema.Type }}{{ MakeFormatIdentifier .Schema.Format }}Converter(data{{ Title .In }}{{ MakeIdentifier .Name }})
		if err != nil {
			server.Srv.ProcessMakeRequestErrors(c, []FieldError{NewFieldError(In{{Title .In}}, "-", "", err)})
			return
  		}
	{{- end}}

`))

var handlerTemplateMimeJsonTypeReqBody = template.Must(template.New("handler").Funcs(funcMap).Parse(`
case "application/json":
	if err := json.NewDecoder(c.Request.Body).Decode(&req.Body.JSON); err != nil {
		server.Srv.ProcessMakeRequestErrors(c, []FieldError{NewFieldError(InBody, "-", "can't decode body from JSON", err)})
		return
	}
`))

var handlerTemplateMimeMultiPartFormTypeReqBody = template.Must(template.New("handler").Funcs(funcMap).Parse(`
case "multipart/form-data":
{{- with .ContentSpec.Schema }}
{{ range $propName, $propSchema := .Properties -}}
	dataBody{{ MakeIdentifier $propName }}, _, err := __gin_get_parameter(c, "body", "{{$propName}}")
	if err != nil {
		server.Srv.ProcessMakeRequestErrors(c, []FieldError{NewFieldError(InBody, "-", "", err)})
		return
	}
	req.Body.Form.{{ MakeIdentifier $propName }}, err = {{ $propSchema.Type }}{{ MakeFormatIdentifier $propSchema.Format }}Converter(dataBody{{ MakeIdentifier $propName }})
	if err != nil {
		server.Srv.ProcessMakeRequestErrors(c, []FieldError{NewFieldError(InBody, "-", "", err)})
		return
	}
{{ end }}
{{- end }}
`))

var handlerTemplateReqBody = template.Must(template.New("handler").Funcs(funcMap).Parse(`
	{{ with .Path.RequestBody }}
	contentType := c.Request.Header.Get("Content-Type")
	
	if contentType == "" {
		server.Srv.ProcessMakeRequestErrors(c, []FieldError{NewFieldError(InBody, "-", "unsupported Content-type", nil)})
		return
	}
	contentType = strings.Split(contentType, ";")[0]

	//TODO: refactor
	switch contentType  {
	{{- end}}
`))

func (Generator) makeHandler(wr io.Writer, interfaceName string, path *openapi3.PathSpec, request openapi3.SchemaType) error {
	err := handlerTemplate.Execute(wr, struct {
		InterfaceName string
		Path          *openapi3.PathSpec
		Request       openapi3.SchemaType
	}{
		InterfaceName: interfaceName,
		Path:          path,
		Request:       request,
	})

	if err != nil {
		return err
	}
	if path.RequestBody != nil {
		err = handlerTemplateReqBody.Execute(wr, struct {
			InterfaceName string
			Path          *openapi3.PathSpec
			Request       openapi3.SchemaType
		}{
			InterfaceName: interfaceName,
			Path:          path,
			Request:       request,
		})

		if err != nil {
			return err
		}


	if contentSpec, ok := path.RequestBody.Content[openapi3.MultiPartFormDataMimeType]; ok == true {
		err := handlerTemplateMimeMultiPartFormTypeReqBody.Execute(wr, struct {
			InterfaceName string
			Path          *openapi3.PathSpec
			Request       openapi3.SchemaType
			ContentSpec   openapi3.ContentType
		}{
			InterfaceName: interfaceName,
			Path:          path,
			Request:       request,
			ContentSpec:   contentSpec,
		})
		if err != nil {
			return err
		}
	}
	if contentSpec, ok := path.RequestBody.Content[openapi3.ApplicationJsonMimeType]; ok == true {
		err := handlerTemplateMimeJsonTypeReqBody.Execute(wr, struct {
			InterfaceName string
			Path          *openapi3.PathSpec
			Request       openapi3.SchemaType
			ContentSpec   openapi3.ContentType
		}{
			InterfaceName: interfaceName,
			Path:          path,
			Request:       request,
			ContentSpec:   contentSpec,
		})
		if err != nil {
			return err
		}
	}


		_, err = wr.Write([]byte("}\n"))

		if err != nil {
			return err
		}
	}

	_, err = wr.Write([]byte(fmt.Sprintf("server.Srv.%v(req, c)\n", path.OperationID)))

	_, err = wr.Write([]byte("}\n"))
	return err
}
