package generator

import (
	"html/template"
	"io"

	"github.com/Masterminds/sprig"

	"github.com/kepkin/gorest/internal/spec/openapi3"
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
	{{ end }}
	// Service methods
	ProcessMakeRequestErrors(c *gin.Context, errors []FieldError)
	ProcessValidateErrors(c *gin.Context, errors []FieldError)
}

type {{ .InterfaceName }}Server struct {
	Srv {{ .InterfaceName }}
}
`))

func (Generator) makeInterface(wr io.Writer, sp openapi3.Spec) error {
	// TODO(a.telyshev): Check not null and uniq OperationId
	return interfaceTemplate.Execute(wr, struct {
		InterfaceName string
		Paths         openapi3.PathMap
	}{
		InterfaceName: MakeIdentifier(sp.Info.Title),
		Paths:         sp.Paths,
	})
}
