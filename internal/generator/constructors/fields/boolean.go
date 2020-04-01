package fields

import (
	"fmt"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var booleanFieldTemplate = template.Must(template.New("booleanField").Parse(
	`switch strings.ToLower({{ .StrVarName }}) {
    case "1", "true", "t":
        result.{{ .Name }} = true
    case "0", "false", "f":
        result.{{ .Name }} = false
    default:
        errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as boolean", nil))
}`))

func MakeBooleanFieldConstructor(f translator.Field, place string) (string, error) {
	if f.Type != translator.BooleanField {
		return "", fmt.Errorf("%v isn't boolean field", f)
	}
	return executeFieldConstructorTemplate(booleanFieldTemplate, f, place)
}
