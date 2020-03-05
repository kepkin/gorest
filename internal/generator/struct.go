package generator

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/translator"
)

var primitiveTypeTemplate = template.Must(template.New("primitiveType").Parse(`
type {{ .Name }} {{ .GoType }}
`))

var structTemplate = template.Must(template.New("struct").Parse(`
type {{ .Name }} struct {
	{{- range $, $field := .Fields2 -}}
	{{ $field.FieldImpl.Name2 }} {{ $field.FieldImpl.GoType }}
	{{ end -}}
}
`))

var structWithJSONTagsTemplate = template.Must(template.New("structWithJSON").Parse(`
type {{ .Name }} struct {
	{{- range $, $field := .Fields2 -}}
	{{ $field.FieldImpl.Name2 }} {{ $field.FieldImpl.GoType }}  ` + "`json:\"" + "{{ $field.FieldImpl.Parameter }}\"`" + `
	{{ end -}}
}
`))

// TODO(a.telyshev): Test me
func (g *Generator) makeStruct(wr io.Writer, def translator.TypeDef, withJSONTags bool) error {
	if def.GoType != "struct" {
		return primitiveTypeTemplate.Execute(wr, def)
	}

	for _, f := range def.Fields2 {
		if f.FieldImpl.IsCustom() {
			g.customFields[f.Schema.Name] = f
		}
	}
	if withJSONTags {
		return structWithJSONTagsTemplate.Execute(wr, def)
	}
	return structTemplate.Execute(wr, def)
}
