package translator

import (
	"github.com/kepkin/gorest/internal/spec/openapi3"
	"strings"
	"text/template"
)

type integerField struct {
	BaseField
	BitSize int  // as for strconv.ParseInt
}

const integerGlobalTpl = `
func integerInt64Converter(input []string) (int64, error) {
	if len(input) != 1 {
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

func numberFloatConverter(input []string) (float64, error) {
	if len(input) != 1 {
		return 0, fmt.Errorf("got array '%v' instead of integer", input)
	}

	return strconv.ParseFloat(input[0], 32)
}
`

type IntegerFieldConstructor struct {
}

func (IntegerFieldConstructor) BuildGlobalCode() (string, error) {
	tpl := template.Must(template.New("integerGlobalTpl").Parse(integerGlobalTpl))
	res := strings.Builder{}
	err := tpl.Execute(&res,
		struct {
			GlobalIdentifier string
		}{""})

	return res.String(), err
}

func (IntegerFieldConstructor) ImportsRequired() []string {
	return []string{}
}

func (IntegerFieldConstructor) RegisterAllFormats(res Translator) {

	res.RegisterObjectFieldConstructor(openapi3.IntegerType, openapi3.None, func(field BaseField, parentName string) Field {
		field.GoType = "int64"
		return &integerField{BaseField: field}
	})

	res.RegisterObjectFieldConstructor(openapi3.IntegerType, openapi3.Integer32bit, func(field BaseField, parentName string) Field {
		field.GoType = "int32"
		return &integerField{BaseField: field}
	})

	res.RegisterObjectFieldConstructor(openapi3.IntegerType, openapi3.Integer64bit, func(field BaseField, parentName string) Field {
		field.GoType = "int64"
		return &integerField{BaseField: field}
	})

	res.RegisterObjectFieldConstructor(openapi3.NumberType, openapi3.None, func(field BaseField, parentName string) Field {
		field.GoType = "int64"
		return &integerField{BaseField: field}
	})

	res.RegisterObjectFieldConstructor(openapi3.NumberType, openapi3.NumberFloat, func(field BaseField, parentName string) Field {
		field.GoType = "float64"
		return &integerField{BaseField: field}
	})
}