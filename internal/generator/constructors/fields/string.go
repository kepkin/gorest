package fields

import (
	"fmt"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var stringFieldTemplate = template.Must(template.New("intField").Parse(
	`result.{{ .Name }} = {{ .StrVarName }}
`))

func MakeStringFieldConstructor(f translator.Field, place string) (string, error) {
	if f.Type != translator.StringField {
		return "", fmt.Errorf("%v isn't string field", f)
	}
	return executeFieldConstructorTemplate(stringFieldTemplate, f, place)
}
