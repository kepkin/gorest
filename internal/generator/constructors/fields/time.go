package fields

import (
	"fmt"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var dateFieldTemplate = template.Must(template.New("dateField").Parse(
	`result.{{ .Name }}, err = time.Parse("2006-01-02", {{ .StrVarName }})
if err != nil {
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as RFC3339 date", err))
}`))

func MakeDateFieldConstructor(f translator.Field, place string) (string, error) {
	if !(f.Type == translator.DateField) {
		return "", fmt.Errorf("%v isn't date field", f)
	}
	return executeFieldConstructorTemplate(dateFieldTemplate, f, place)
}

var dateTimeFieldTemplate = template.Must(template.New("timeField").Parse(
	`result.{{ .Name }}, err = time.Parse(time.RFC3339, {{ .StrVarName }})
if err != nil {
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as RFC3339 time", err))
}`))

func MakeDateTimeFieldConstructor(f translator.Field, place string) (string, error) {
	if !(f.Type == translator.DateTimeField) {
		return "", fmt.Errorf("%v isn't datetime field", f)
	}
	return executeFieldConstructorTemplate(dateTimeFieldTemplate, f, place)
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
