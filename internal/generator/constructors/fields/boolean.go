package fields

import (
	"fmt"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var booleanFieldTemplate = template.Must(template.New("booleanField").Parse(
	`{{ .StrVarName }} = strings.strings.ToLower({{ .StrVarName }})
switch {{ .StrVarName }} {
    case "1", "true", "t":
        result.{{ .Name }} = true
    default:
        result.{{ .Name }} = false
}`))

func MakeBooleanFieldConstructor(f translator.Field, place string) (string, error) {
	if f.Type != translator.BooleanField {
		return "", fmt.Errorf("%v isn't boolean field", f)
	}
	return executeFieldConstructorTemplate(booleanFieldTemplate, f, place)
}
