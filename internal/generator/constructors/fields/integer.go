package fields

import (
	"fmt"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var intFieldTemplate = template.Must(template.New("intField").Parse(
	`result.{{ .Name }}, err = strconv.ParseInt({{ .StrVarName }}, 10, {{ .Schema.BitSize }})
if err != nil {
	{{- if eq .Schema.BitSize 0 }}
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as integer", err))
	{{- else }}
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as {{ .Schema.BitSize }} bit integer", err))
	{{- end }}
}`))

func MakeIntFieldConstructor(f translator.Field, place string) (string, error) {
	if f.Type != translator.IntegerField {
		return "", fmt.Errorf("%v isn't integer field", f)
	}
	return executeFieldConstructorTemplate(intFieldTemplate, f, place)
}
