package generator

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var structTemplate = template.Must(template.New("struct").Parse(`
type {{ .Name }} struct {
	{{- range $, $field := .Fields }}
	{{ $field.Name }} {{ $field.GoType -}}
	{{ end }}
}
`))

func (g *Generator) makeStruct(wr io.Writer, def translator.TypeDef) error {
	for _, f := range def.Fields {
		if f.IsCustom() {
			g.customFields[f.GoType] = f
		}
	}
	return structTemplate.Execute(wr, def)
}
