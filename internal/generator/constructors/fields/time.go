package fields

import (
	"fmt"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var timeFieldTemplate = template.Must(template.New("timeField").Parse(
	`result.{{ .Name }}, err = time.Parse(time.RFC3339, {{ .StrVarName }})
if err != nil {
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as RFC3339 time", err))
}`))

func MakeTimeFieldConstructor(f translator.Field, place string) (string, error) {
	if !(f.Type == translator.DateField || f.Type == translator.DateTimeField) {
		return "", fmt.Errorf("%v isn't date[time] field", f)
	}
	return executeFieldConstructorTemplate(timeFieldTemplate, f, place)
}

var unixTimeFieldTemplate = template.Must(template.New("timeField").Parse(
	`{{ .SecondsVarName }}, err := strconv.ParseInt({{ .StrVarName }}, 10, 64)
if err != nil {
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as 64 bit integer", err))
} else {
	result.{{ .Name }} = time.Unix({{ .SecondsVarName }}, 0)
}`))

func MakeUnixTimeFieldConstructor(f translator.Field, place string) (string, error) {
	if f.Type != translator.UnixTimeField {
		return "", fmt.Errorf("%v isn't unix time field", f)
	}
	return executeFieldConstructorTemplate(unixTimeFieldTemplate, f, place)
}
