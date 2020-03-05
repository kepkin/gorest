package translator

import (
	"strings"
	"text/template"
)

type IntegerFieldImpl struct {
	Field
	BitSize int  // as for strconv.ParseInt
}

const integerGlobalTpl = `
func integerInt64Converter(input []string) (int64, error) {
	if len(input) > 1 {
		return 0, fmt.Errorf("got array '%v' instead of integer", input)
	}

	return strconv.ParseInt(input[0], 10, 64)
}

func integerInt32Converter(input []string) (int32, error) {
	if len(input) > 1 {
		return 0, fmt.Errorf("got array '%v' instead of integer", input)
	}

	v, err := strconv.ParseInt(input[0], 10, 32)
	return int32(v), err
}

func integerConverter(input []string) (int64, error) {
	return integerInt64Converter(input)
}
`

func (c *IntegerFieldImpl) BuildGlobalCode() (string, error) {
	tpl := template.Must(template.New("integerGlobalTpl").Parse(integerGlobalTpl))
	res := strings.Builder{}
	err := tpl.Execute(&res,
		struct {
			GlobalIdentifier string
		}{""})

	return res.String(), err
}

func (c *IntegerFieldImpl) ContextErrorRequired() bool {
	return false
}
