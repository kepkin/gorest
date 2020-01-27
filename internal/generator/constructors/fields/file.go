package fields

import (
	"fmt"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var fileFieldTemplate = template.Must(template.New("fileField").Parse(
	`result.{{ .Name }}, err = c.FormFile("{{ .Parameter }}")
if err != nil {
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't extract file from form-data", err))
}`))

func MakeFileFieldConstructor(f translator.Field, place string) (string, error) {
	if f.Type != translator.FileField {
		return "", fmt.Errorf("%v isn't file field", f)
	}
	return executeFieldConstructorTemplate(fileFieldTemplate, f, place)
}
