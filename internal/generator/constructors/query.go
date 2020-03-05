package constructors

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/constructors/fields"
	"github.com/kepkin/gorest/internal/generator/translator"
)

// MakeQueryParamsConstructor receive a query params struct definition and generate corresponding constructor
func MakeQueryParamsConstructor(wr io.Writer, def translator.TypeDef) error {
	return queryParamsConstructorTemplate.Execute(wr, def)
}

var queryParamsConstructorTemplate = template.Must(template.New("queryParamsConstructor").Funcs(fields.BaseConstructor).Parse(`
func Make{{ .Name }}(c *gin.Context) (result {{ .Name }}, errors []FieldError) {
	{{- if .HasNoStringFields }}
	var err error
	{{ end }}

	{{- range $, $field := .Fields }}
	{{- with $field }}
		{{- if .IsString }}
			{{- if .CheckDefault}}
				result.{{ .Name }}, ok = c.GetQuery("{{ .Parameter }}")
				if !ok {
					result.{{ .Name }} = "{{ .Schema.Default }}"
				}
			{{ else }}
				result.{{ .Name }}, _ = c.GetQuery("{{ .Parameter }}")
			{{- end }}
		{{- else if or (.IsCustom)  (.IsInteger)  (.IsFloat)  (.IsDate)  (.IsDateTime)  (.IsUnixTime)}}
			 {{- if .CheckDefault}}
				{{ .StrVarName }}, ok := c.GetQuery("{{ .Parameter }}")
				if !ok {
				   {{ .StrVarName }} = "{{ .Schema.Default }}"
				}
			 {{ else }}
				{{ .StrVarName }}, _ := c.GetQuery("{{ .Parameter }}")
			 {{- end }}
		{{- end }}

		{{- BaseValueFieldConstructor . "InQuery" }}

	{{- end -}}
	{{ end -}}
	return
}
`))
