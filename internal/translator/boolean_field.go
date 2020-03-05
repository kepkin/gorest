package translator

import (
	"github.com/kepkin/gorest/internal/spec/openapi3"
	"strings"
	"text/template"
)

type booleanField struct {
	BaseField

	ImplIdentifier string
}

const booleanGlobalTpl = `
func booleanConverter(input []string) (bool, error) {
	if len(input) == 0 {
        return false, fmt.Errorf("got empty value instead of boolean", input)
	}

	if len(input) > 1 {
        return false, fmt.Errorf("got array '%v' instead of boolean", input)
	}

	switch strings.ToLower(input[0]) {
		case "1", "true", "t":
			return true, nil
		case "0", "false", "f":
			return false, nil
		default:
			return false, fmt.Errorf("can't parse '%v' as boolean", input[0])
	}
}
`

type BooleanFieldConstructor struct {
}

func (BooleanFieldConstructor) RegisterAllFormats(translator Translator) {
	translator.RegisterObjectFieldConstructor(openapi3.BooleanType, openapi3.Format(""), func(field BaseField, parentName string) Field {
		field.GoType = "bool"
		return &booleanField{field, "a"}
	})
}

func (BooleanFieldConstructor) BuildGlobalCode() (string, error) {
	tpl := template.Must(template.New("booleanGlobalTpl").Parse(booleanGlobalTpl))
	res := strings.Builder{}
	err := tpl.Execute(&res,
		struct {
			GlobalIdentifier string
		}{"a"})  //TODO: it's a prefix for converter functions. Might be obsolete

	return res.String(), err
}

func (BooleanFieldConstructor) ImportsRequired() []string {
	return []string{}
}