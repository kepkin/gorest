package fields

import (
	"fmt"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var floatFieldTemplate = template.Must(template.New("floatField").Parse(
	`result.{{ .Name }}, err = strconv.ParseFloat({{ .StrVarName }}, 10, {{ .Schema.BitSize }})
if err != nil {
	{{- if eq .Schema.BitSize 0 }}
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as float", err))
	{{- else }}
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as {{ .Schema.BitSize }} bit float", err))
	{{- end }}
}`))

func MakeFloatFieldConstructor(f translator.Field, place string) (string, error) {
	if f.Type != translator.FloatField {
		return "", fmt.Errorf("%v isn't float field", f)
	}
	return executeFieldConstructorTemplate(floatFieldTemplate, f, place)
}
