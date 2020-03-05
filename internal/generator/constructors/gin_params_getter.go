package constructors

import (
	"strings"
	"text/template"

	"github.com/kepkin/gorest/internal/spec/openapi3"
)

type RequestParamsGetter interface {
	BuildGlobalCode() (string, error)

	MakeCookieGetter(parameter string) (string, error)
	MakeQueryGetter(parameter string) (string, error)
	MakeHeaderGetter(parameter string) (string, error)
	MakePathGetter(parameter string) (string, error)
	MakeBodyGetter(parameter string) (string, error)

	MakeHandler() (string, error)
 }


type GinParamsGetter struct {}

func (g *GinParamsGetter) BuildGlobalCode() (string, error) {
	return `
ginGetCookie := func(c *gin.Context, param string) (string, bool) {
		cookie, err := c.Request.Cookie(param)
		if err == http.ErrNoCookie {
			return "", false
		}
		return cookie.Value, true
}
`, nil
}


func (g *GinParamsGetter) MakeCookieGetter(parameter string) (string, error) {
	return "", nil
}

func (g *GinParamsGetter) MakeQueryGetter(parameter string) (string, error) {
	return "", nil
}

func (g *GinParamsGetter) MakeHeaderGetter(parameter string) (string, error) {
	return "", nil
}

func (g *GinParamsGetter) MakePathGetter(parameter string) (string, error) {
	return "", nil
}

func (g *GinParamsGetter) MakeBodyGetter(parameter string) (string, error) {
	return "", nil
}


var handlerTemplate = template.Must(template.New("handler").Parse(`
func (server {{ .InterfaceName }}Server) _{{ .InterfaceName }}_{{ .Path.OperationID }}_Handler(c *gin.Context) {
	c.Set(handlerNameKey, "{{ .Path.OperationID }}")
	
	{{ with .Request.Properties }}
		{{- range $, $field := .Fields }}
			result.{{ .Name }}, errors = Make{{ .GoType }}(c)
			if errors != nil {
				server.Srv.ProcessMakeRequestErrors(c, errors)
				return
			}
		{{ end -}}


    //TODO
	//errors = req.Validate()
	//if len(errors) > 0 {
	//	server.Srv.ProcessValidateErrors(c, errors)
	//	return
	//}

	server.Srv.{{ $.Path.OperationID }}(req, c)
	{{- else -}}
	server.Srv.{{ $.Path.OperationID }}({{ $.Path.OperationID }}Request{}, c)
	{{- end }}
}
`))

func (g *GinParamsGetter) MakeHandler(interfaceName string, path *openapi3.PathSpec, request openapi3.SchemaType) (string, error) {
	res := strings.Builder{}

	err := handlerTemplate.Execute(&res, struct {
		InterfaceName string
		Path          *openapi3.PathSpec
		Request       openapi3.SchemaType
	}{
		InterfaceName: interfaceName,
		Path:          path,
		Request:       request,
	})

	return res.String(), err
}