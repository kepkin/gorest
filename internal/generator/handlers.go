package generator

import (
	"container/list"
	"fmt"
	"io"
	"sort"
	"strings"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/constructors"

	"github.com/kepkin/gorest/internal/generator/translator"
	"github.com/kepkin/gorest/internal/spec/openapi3"
)

func (g *Generator) makeHandlers(wr io.Writer, sp openapi3.Spec) error {
	interfaceName := translator.MakeIdentifier(sp.Info.Title)

	sortedPaths := make([]string, 0, len(sp.Paths))
	for k := range sp.Paths {
		sortedPaths = append(sortedPaths, k)
	}

	sort.Strings(sortedPaths)

	for _, pathKey := range sortedPaths {
		path := sp.Paths[pathKey]
		for _, method := range path.Methods() {
			if method != nil {
				root, err := g.makeRequest(wr, interfaceName, method)
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

type s = openapi3.SchemaType

func (g *Generator) makeRequest(wr io.Writer, interfaceName string, method *openapi3.PathSpec) (s, error) { //nolint:gocyclo,gocognit
	if _, err := fmt.Fprintf(wr, "\n// _%s_%s_Handler\n", interfaceName, method.OperationID); err != nil {
		return s{}, err
	}

	// Make and fill root request schema
	request := openapi3.NewObjectSchema(method.OperationID + "Request")

	queryParams := openapi3.NewObjectSchema("")
	pathParams := openapi3.NewObjectSchema("")
	cookieParams := openapi3.NewObjectSchema("")
	headerParams := openapi3.NewObjectSchema("")
	body := openapi3.NewObjectSchema("")

	for _, param := range method.Parameters {
		switch param.In {
		case "query":
			queryParams.Properties[param.Name] = param.Schema
		case "path":
			pathParams.Properties[param.Name] = param.Schema
		case "cookie":
			cookieParams.Properties[param.Name] = param.Schema
		case "header":
			headerParams.Properties[param.Name] = param.Schema
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
		request.Properties["Headers"] = &headerParams
	}

	if method.RequestBody != nil {
		for mimeType, content := range method.RequestBody.Content {
			switch mimeType {
			case "application/json":
				body.Properties["JSON"] = content.Schema
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
	defs, err := translator.ProcessRootSchema(request)
	if err != nil {
		return s{}, err
	}

	// Make structs
	for _, d := range defs {
		if err := g.makeStruct(wr, d, false); err != nil {
			return s{}, err
		}

		if len(d.Fields) != 0 {
			if err = g.makeValidateFunc(wr, d); err != nil {
				return s{}, err
			}
		}
	}

	rootDef := defs[0]
	if err := constructors.MakeRequestConstructor(wr, rootDef); err != nil {
		return s{}, err
	}

	var formDataDef *translator.TypeDef
	for i, d := range defs {
		if strings.HasSuffix(d.Name, "RequestBodyForm") {
			formDataDef = &defs[i]
			break
		}
	}

	type SchemaConstructorPair struct {
		Params      *openapi3.SchemaType
		Constructor func(io.Writer, translator.TypeDef) error
	}

	for i, e := range []SchemaConstructorPair{
		{&body, constructors.MakeBodyConstructor},
		{&headerParams, constructors.MakeHeaderParamsConstructor},
		{&cookieParams, constructors.MakeCookieParamsConstructor},
		{&pathParams, constructors.MakePathParamsConstructor},
		{&queryParams, constructors.MakeQueryParamsConstructor},
	} {
		if len(e.Params.Properties) == 0 {
			continue
		}

		def, err := translator.ProcessObjSchema(*e.Params, list.New())
		if err != nil {
			return s{}, err
		}

		def.Name = defs[0].Name + def.Name
		if err := e.Constructor(wr, def); err != nil {
			return s{}, err
		}

		// MakeBodyConstructor
		if i == 0 && formDataDef != nil {
			if err := constructors.MakeFormDataConstructor(wr, *formDataDef); err != nil {
				return s{}, err
			}
		}
	}

	return request, nil
}

var handlerTemplate = template.Must(template.New("handler").Parse(`
func (server {{ .InterfaceName }}Server) _{{ .InterfaceName }}_{{ .Path.OperationID }}_Handler(c *gin.Context) {
	c.Set(handlerNameKey, "{{ .Path.OperationID }}")
	
	{{ with .Request.Properties }}
	req, errors := Make{{ $.Path.OperationID }}Request(c)
	if len(errors) > 0 {
		server.Srv.ProcessMakeRequestErrors(c, errors)
		return
	}

	errors = req.Validate()
	if len(errors) > 0 {
		server.Srv.ProcessValidateErrors(c, errors)
		return
	}

	server.Srv.{{ $.Path.OperationID }}(req, c)
	{{- else -}}
	server.Srv.{{ $.Path.OperationID }}({{ $.Path.OperationID }}Request{}, c)
	{{- end }}
}
`))

func (Generator) makeHandler(wr io.Writer, interfaceName string, path *openapi3.PathSpec, request openapi3.SchemaType) error {
	return handlerTemplate.Execute(wr, struct {
		InterfaceName string
		Path          *openapi3.PathSpec
		Request       openapi3.SchemaType
	}{
		InterfaceName: interfaceName,
		Path:          path,
		Request:       request,
	})
}
