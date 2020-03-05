package constructors

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/constructors/fields"
	"github.com/kepkin/gorest/internal/generator/translator"
)

// MakeHeaderParamsConstructor receive a header params struct definition and generate corresponding constructor
func MakeHeaderParamsConstructor(wr io.Writer, def translator.TypeDef) error {
	return headerParamsConstructorTemplate.Execute(wr, def)
}

var headerParamsConstructorTemplate = template.Must(template.New("headerParamsConstructor").Funcs(fields.BaseConstructor).Parse(`
func Make{{ .Name }}(c *gin.Context) (result {{ .Name }}, errors []FieldError) {
	{{- if .HasNoStringFields }}
	var err error
	{{ end }}

	{{- range $, $field := .Fields }}
	{{- with $field }}
		{{- if .IsString }}
			{{- if .CheckDefault}}
				result.{{ .Name }} = c.Request.Header.Get("{{ .Parameter }}")
				if result.{{ .Name }} == "" {
					result.{{ .Name }} = "{{ .Schema.Default }}"
				}
			{{ else }}
				result.{{ .Name }} = c.Request.Header.Get("{{ .Parameter }}")
			{{- end }}
		{{- else if or (.IsCustom)  (.IsInteger)  (.IsFloat)  (.IsDate)  (.IsDateTime)  (.IsUnixTime)}}
			 {{- if .CheckDefault}}
				{{ .StrVarName }} := c.Request.Header.Get("{{ .Parameter }}")
				if {{ .StrVarName }} != "" {
				   {{ .StrVarName }} = "{{ .Schema.Default }}"
				}
			 {{ else }}
				{{ .StrVarName }} := c.Request.Header.Get("{{ .Parameter }}")
			 {{- end }}
		{{- end }}

		{{- BaseValueFieldConstructor . "InHeader" }}

	{{- end -}}
	{{ end -}}
	return
}
`))
