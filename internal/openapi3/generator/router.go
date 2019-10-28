package generator

import (
	"io"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/kepkin/gorest/internal/openapi3/spec"
	"github.com/kepkin/gorest/internal/openapi3/translator"
)

var routerTemplate = template.Must(template.New("router").Funcs(sprig.GenericFuncMap()).Funcs(template.FuncMap{
	"ConvertUrl": convertUrl,
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
}`))

func MakeRouter(wr io.Writer, sp spec.Spec) error {
	return routerTemplate.Execute(wr, struct {
		InterfaceName string
		Paths         spec.PathMap
	}{
		InterfaceName: translator.MakeIdentifier(sp.Info.Title),
		Paths:         sp.Paths,
	})
}

func convertUrl(url string) string {
	return strings.ReplaceAll(strings.ReplaceAll(url, "{", ":"), "}", "")
}
