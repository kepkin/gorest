package fields

import (
	"fmt"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var customFieldTemplate = template.Must(template.New("customField").Parse(
	`result.{{ .Name }} = {{ .GoType }}{}
if err = result.{{ .Name }}.SetFromString({{ .StrVarName }}); err != nil {
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", fmt.Sprintf("can't create from string '%s'", {{ .StrVarName }}), err))
}`))

func MakeCustomFieldConstructor(f translator.Field, place string) (string, error) {
	if f.Type != translator.CustomField {
		return "", fmt.Errorf("%v isn't field of user defined type", f)
	}
	return executeFieldConstructorTemplate(customFieldTemplate, f, place)
}
