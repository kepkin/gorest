package constructors

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

// MakeHeaderParamsConstructor receive a header params struct definition and generate corresponding constructor
func MakeHeaderParamsConstructor(wr io.Writer, def translator.TypeDef) error {
	return headerParamsConstructorTemplate.Execute(wr, def)
}

var headerParamsConstructorTemplate = template.Must(template.New("headerParamsConstructor").Funcs(template.FuncMap{
	"IntConstructor":   makeIntConstructor,
	"FloatConstructor": makeFloatConstructor,
}).Parse(`
func Make{{ .Name }}(c *gin.Context) (result {{ .Name }}, errors []FieldError) {
	{{- range $, $field := .Fields }}
	{{- with $field }}
		
		{{- if .IsString }}
			result.{{ .Name }} = c.Request.Header.Get("{{ .Parameter }}")
		{{- end }}

		{{- if .IsInteger }}
			{{ .StrVarName }} = c.Request.Header.Get("{{ .Parameter }}")
			{{ IntConstructor . "InHeader" }}
		{{- end }}

		{{- if .IsFloat }}
			{{ .StrVarName }} = c.Request.Header.Get("{{ .Parameter }}")
			{{ FloatConstructor . "InHeader" }}
		{{- end }}

	{{- end }}
	{{ end -}}
	return
}
`))
