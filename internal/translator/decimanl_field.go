package translator

import (
	"strings"
	"text/template"
)

type DecimalFieldImpl struct {
	Field
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

func (c *DecimalFieldImpl) BuildGlobalCode() (string, error) {
	tpl := template.Must(template.New("decimalGlobalTpl").Parse(decimalGlobalTpl))
	res := strings.Builder{}
	err := tpl.Execute(&res,
		struct {
			GlobalIdentifier string
		}{""})

	return res.String(), err
}

func (c *DecimalFieldImpl) ContextErrorRequired() bool {
	return false
}

func (c *DecimalFieldImpl) ImportsRequired() []string {
	return []string{
		"github.com/shopspring/decimal",
	}
}
