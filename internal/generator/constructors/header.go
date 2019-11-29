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

var headerParamsConstructorTemplate = template.Must(template.New("headerParamsConstructor").Funcs(fields.Constructors).Parse(`
func Make{{ .Name }}(c *gin.Context) (result {{ .Name }}, errors []FieldError) {
	{{- if .HasNoStringFields }}
	var err error
	{{ end }}

	{{- range $, $field := .Fields }}
	{{- with $field }}
		
		{{- if .IsString }}
			result.{{ .Name }} = c.Request.Header.Get("{{ .Parameter }}")
		{{- end }}

		{{- if .IsCustom }}
			{{ .StrVarName }} := c.Request.Header.Get("{{ .Parameter }}")
			{{ CustomFieldConstructor . "InHeader" }}
		{{- end }}

		{{- if .IsInteger }}
			{{ .StrVarName }} := c.Request.Header.Get("{{ .Parameter }}")
			{{ IntConstructor . "InHeader" }}
		{{- end }}

		{{- if .IsFloat }}
			{{ .StrVarName }} := c.Request.Header.Get("{{ .Parameter }}")
			{{ FloatConstructor . "InHeader" }}
		{{- end }}

		{{- if or .IsDate .IsDateTime }}
			{{ .StrVarName }}, _ := c.GetQuery("{{ .Parameter }}")
			{{ DateTimeConstructor . "InHeader" }}
		{{- end }}

        {{- if .IsDateTime }}
			{{ .StrVarName }}, _ := c.GetQuery("{{ .Parameter }}")
			{{ DateTimeConstructor . "InHeader" }}
		{{- end }}

		{{- if .IsUnixTime }}
			{{ .StrVarName }}, _ := c.GetQuery("{{ .Parameter }}")
			{{ UnixTimeConstructor . "InHeader" }}
		{{- end }}

	{{- end }}
	{{ end -}}
	return
}
`))
