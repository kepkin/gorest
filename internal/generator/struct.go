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

var structWithJSONTagsTemplate = template.Must(template.New("struct").Parse(`
type {{ .Name }} struct {
	{{- range $, $field := .Fields -}}
	{{ $field.Name }} {{ $field.GoType -}} ` + "`json:\"" + "{{ $field.Parameter }}\"`" + `
	{{ end -}}
}
`))

// TODO(a.telyshev): Test me
func (g *Generator) makeStruct(wr io.Writer, def translator.TypeDef, withJSONTags bool) error {
	for _, f := range def.Fields {
		if f.IsCustom() {
			g.customFields[f.GoType] = f
		}
	}
	if withJSONTags {
		return structWithJSONTagsTemplate.Execute(wr, def)
	}
	return structTemplate.Execute(wr, def)
}
