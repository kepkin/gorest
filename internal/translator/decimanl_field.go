package translator

import (
	"github.com/kepkin/gorest/internal/spec/openapi3"
	"strings"
	"text/template"
)

type decimalField struct {
	BaseField
}

const decimalGlobalTpl = `
func integerDecimalConverter(input []string) (decimal.Decimal, error) {
	if len(input) > 1 {
		return decimal.Decimal{}, fmt.Errorf("got array '%v' instead of decimal", input)
	}

	return decimal.NewFromString(input[0])
}

func stringDecimalConverter(input []string) (decimal.Decimal, error) {
	if len(input) > 1 {
		return decimal.Decimal{}, fmt.Errorf("got array '%v' instead of decimal", input)
	}

	return decimal.NewFromString(input[0])
}
`

type DecimalFieldConstructor struct {
}

func (DecimalFieldConstructor) BuildGlobalCode() (string, error) {
	tpl := template.Must(template.New("decimalGlobalTpl").Parse(decimalGlobalTpl))
	res := strings.Builder{}
	err := tpl.Execute(&res,
		struct {
			GlobalIdentifier string
		}{""})

	return res.String(), err
}

func (DecimalFieldConstructor) ImportsRequired() []string {
	return []string{}
}

func (DecimalFieldConstructor) RegisterAllFormats(res Translator) {
	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.Format("decimal"), func(field BaseField, parentName string) Field {
		field.GoType = "decimal.Decimal"
		return &decimalField{BaseField: field}
	})

	res.RegisterObjectFieldConstructor(openapi3.NumberType, openapi3.Format("decimal"), func(field BaseField, parentName string) Field {
		field.GoType = "decimal.Decimal"
		return &decimalField{BaseField: field}
	})

	res.RegisterObjectFieldConstructor(openapi3.IntegerType, openapi3.Format("decimal"), func(field BaseField, parentName string) Field {
		field.GoType = "decimal.Decimal"
		return &decimalField{BaseField: field}
	})
}