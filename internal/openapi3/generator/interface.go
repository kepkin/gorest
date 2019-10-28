package generator

import (
	"html/template"
	"io"

	"github.com/Masterminds/sprig"

	"github.com/kepkin/gorest/internal/openapi3/spec"
	"github.com/kepkin/gorest/internal/openapi3/translator"
)

var interfaceTemplate = template.Must(template.New("interfaceTmpl").Funcs(sprig.GenericFuncMap()).Parse(`
{{- define "interfaceMethod" -}}
    {{ .OperationID }}(in {{ .OperationID }}Request, c *gin.Context)
{{- end}}

type {{ .InterfaceName }} interface {
    {{- range $url, $m := .Paths -}}
        {{- range $methodName, $method := (dict "GET" $m.Get "POST" $m.Post "PATCH" $m.Patch "DELETE" $m.Delete "PUT" $m.Put "OPTIONS" $m.Options) }}
			{{- with $method }}
				// {{ print $methodName " " $url }}
				{{ template "interfaceMethod" $method }}
            {{- end }}
        {{- end }}
	{{ end -}}
}

type {{ .InterfaceName }}Server struct {
	Srv {{ .InterfaceName }}
}
`))

func MakeInterface(wr io.Writer, sp spec.Spec) error {
	// TODO(a.telyshev): Check not null and uniq OperationId
	return interfaceTemplate.Execute(wr, struct {
		InterfaceName string
		Paths         spec.PathMap
	}{
		InterfaceName: translator.MakeIdentifier(sp.Info.Title),
		Paths:         sp.Paths,
	})
}
