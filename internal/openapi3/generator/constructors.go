package generator

import (
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/kepkin/gorest/internal/openapi3/translator"
)

var constructorTemplate = template.Must(template.New("constructor").Funcs(template.FuncMap{
	"MakeFieldConstructor": makeFieldConstructor,
}).Parse(`
func Make{{ .Name }}(c *gin.Context) (result {{ .Name }}, errors []FieldError) {
	{{- with .Fields }}	
		var err error
		_ = err
	{{ end -}}
	{{- range $, $field := .Fields }}
		{{- MakeFieldConstructor $field }}
	{{ end -}}
	return
}
`))

var intFieldConstructorTemplate = template.Must(template.New("intFieldConstructor").Parse(`
	result.{{ .Name }}, err = strconv.ParseInt(ExtractParameter(c, "{{ .Parameter }}", {{ .Place }}), 10, {{ .BitSize }})
	if err != nil {
		errors = append(errors, FieldError{
			Field:   "{{ .Parameter }}",
			{{- if eq .BitSize 0 }}
			Message: "can't parse as integer",
			{{- else }}
			Message: "can't parse as {{ .BitSize }} bit integer",
			{{- end }}
			Reason:  err.Error(),
		})
	}`))

var stringFieldConstructorTemplate = template.Must(template.New("stringFieldConstructor").Parse(`
	result.{{ .Name }} = ExtractParameter(c, "{{ .Parameter }}", {{ .Place }})`))

var structFieldConstructorTemplate = template.Must(template.New("structFieldConstructor").Funcs(template.FuncMap{
	"ToLower": strings.ToLower,
}).Parse(`
	{{- if lt .Level 2 }}
		{{ .Name | ToLower }}, nestedErrors := Make{{ .Type }}(c)
	{{- else }}
		{{ .Name | ToLower }}, nestedErrors := Make{{ .Type }}(ExtractParameter(c, "{{ .Parameter }}", {{ .Place }}))
	{{- end }}
	if nestedErrors != nil {
		errors = append(errors, nestedErrors...)
	}
	result.{{ .Name }} = {{ .Name | ToLower }}`))

func MakeConstructor(wr io.Writer, def translator.TypeDef) error {
	return constructorTemplate.Execute(wr, def)
}

func makeFieldConstructor(field interface{}) (result string, err error) {
	wr := new(strings.Builder)

	switch f := field.(type) {
	case translator.IntegerField:
		err = intFieldConstructorTemplate.Execute(wr, f)

	case translator.StringField:
		err = stringFieldConstructorTemplate.Execute(wr, f)

	case translator.StructField:
		err = structFieldConstructorTemplate.Execute(wr, f)

	default:
		err = fmt.Errorf("unknown field type: %T%+v", f, f)
	}

	if err != nil {
		return
	}
	result = wr.String()
	return
}
