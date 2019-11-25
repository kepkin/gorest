package constructors

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

// MakeQueryParamsConstructor receive a query params struct definition and generate corresponding constructor
func MakeQueryParamsConstructor(wr io.Writer, def translator.TypeDef) error {
	return queryParamsConstructorTemplate.Execute(wr, def)
}

var queryParamsConstructorTemplate = template.Must(template.New("queryParamsConstructor").Funcs(template.FuncMap{
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
			result.{{ .Name }}, _ = c.GetQuery("{{ .Parameter }}")
		{{- end }}

		{{- if .IsCustom }}
			{{ .StrVarName }}, _ := c.GetQuery("{{ .Parameter }}")
			{{ CustomFieldConstructor . "InQuery" }}
		{{- end }}

		{{- if .IsInteger }}
			{{ .StrVarName }}, _ := c.GetQuery("{{ .Parameter }}")
			{{ IntConstructor . "InQuery" }}
		{{- end }}

		{{- if .IsFloat }}
			{{ .StrVarName }}, _ := c.GetQuery("{{ .Parameter }}")
			{{ FloatConstructor . "InQuery" }}
		{{- end }}

	{{- end }}
	{{ end -}}
	return
}
`))
