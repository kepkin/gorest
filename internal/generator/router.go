package generator

import (
	"io"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"

	"github.com/kepkin/gorest/internal/spec/openapi3"
)

var routerTemplate = template.Must(template.New("router").Funcs(sprig.GenericFuncMap()).Funcs(template.FuncMap{
	"ConvertUrl": convertURL,
}).Parse(`
// Router
func RegisterRoutes(r *gin.Engine, api {{ .InterfaceName }}) {
	e := &{{ .InterfaceName }}Server{api}
	{{ range $url, $m := .Paths -}}
        {{- range $methodName, $method := (dict "GET" $m.Get "POST" $m.Post "PATCH" $m.Patch "DELETE" $m.Delete "PUT" $m.Put "OPTIONS" $m.Options) }}
			{{- with $method }}
				r.Handle("{{ $methodName }}", "{{ ConvertUrl $url }}", e._{{ $.InterfaceName }}_{{ .OperationID }}_Handler)
            {{- end }}
        {{- end }}
	{{ end -}}
}

type HandlerRegister func(operationID, httpMethod, relativePath string, handler gin.HandlerFunc)

// CustomRouter
func RegisterRoutesCustom(handlerRegister HandlerRegister, api {{ .InterfaceName }}) {
	e := &{{ .InterfaceName }}Server{api}
	{{ range $url, $m := .Paths -}}
        {{- range $methodName, $method := (dict "GET" $m.Get "POST" $m.Post "PATCH" $m.Patch "DELETE" $m.Delete "PUT" $m.Put "OPTIONS" $m.Options) }}
			{{- with $method }}
				handlerRegister("{{ .OperationID }}", "{{ $methodName }}", "{{ ConvertUrl $url }}", e._{{ $.InterfaceName }}_{{ .OperationID }}_Handler)
            {{- end }}
        {{- end }}
	{{ end -}}
}
`))

func (Generator) makeRouter(wr io.Writer, sp openapi3.Spec) error {
	return routerTemplate.Execute(wr, struct {
		InterfaceName string
		Paths         openapi3.PathMap
	}{
		InterfaceName: MakeIdentifier(sp.Info.Title),
		Paths:         sp.Paths,
	})
}

func convertURL(url string) string {
	return strings.ReplaceAll(strings.ReplaceAll(url, "{", ":"), "}", "")
}
