package constructors

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

// MakePathParamsConstructor receive a path params struct definition and generate corresponding constructor
func MakePathParamsConstructor(wr io.Writer, def translator.TypeDef) error {
	return pathParamsConstructorTemplate.Execute(wr, def)
}

var pathParamsConstructorTemplate = template.Must(template.New("pathParamsConstructor").Funcs(template.FuncMap{
	"CustomFieldConstructor": makeCustomFieldConstructor,
	"IntConstructor":         makeIntFieldConstructor,
	"FloatConstructor":       makeFloatFieldConstructor,
}).Parse(`
func Make{{ .Name }}(c *gin.Context) (result {{ .Name }}, errors []FieldError) {
	{{- if .HasNoStringFields }}
	var err error
	{{ end }}

	{{- range $, $field := .Fields }}
	{{- with $field }}
		
		{{- if .IsString }}
			result.{{ .Name }}, _ = c.Params.Get("{{ .Parameter }}")
		{{- end }}

		{{- if .IsCustom }}
			{{ .StrVarName }}, _ := c.Params.Get("{{ .Parameter }}")
			{{ CustomFieldConstructor . "InPath" }}
		{{- end }}

		{{- if .IsInteger }}
			{{ .StrVarName }}, _ := c.Params.Get("{{ .Parameter }}")
			{{ IntConstructor . "InPath" }}
		{{- end }}

		{{- if .IsFloat }}
			{{ .StrVarName }}, _ := c.Params.Get("{{ .Parameter }}")
			{{ FloatConstructor . "InPath" }}
		{{- end }}

	{{- end }}
	{{ end -}}
	return
}
`))
