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
func Make{{ .Name }}(ctx *gin.Context) (result {{ .Name }}, errors []FieldError) {
	{{- range $, $field := .Fields }}
		{{- MakeFieldConstructor $field }}
	{{ end -}}
	return
}`))

var intFieldConstructorTemplate = template.Must(template.New("intFieldConstructor").Parse(`
	result.{{ .Name }}, err = strconv.ParseInt(ExtractParameter(ctx, "{{ .Parameter }}", {{ .Place }}), 10, {{ .BitSize }})
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
	result.{{ .Name }} = ExtractParameter(ctx, "{{ .Parameter }}", {{ .Place }})`))

var structFieldConstructorTemplate = template.Must(template.New("structFieldConstructor").Parse(`
	{{- if lt .Level 2 }}
	result.{{ .Name }}, nestedErrors = Make{{ .Type }}(ctx)
	{{- else }}
	result.{{ .Name }}, nestedErrors = Make{{ .Type }}(ExtractParameter(ctx, "{{ .Parameter }}", {{ .Place }}))
	{{- end }}
	if nestedErrors != nil {
		errors = append(errors, nestedErrors)
	}`))

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
