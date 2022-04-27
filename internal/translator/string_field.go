package translator

import (
	"github.com/kepkin/gorest/internal/spec/openapi3"
	"strings"
	"text/template"
)

type stringField struct {
	BaseField
}

//TODO: separate const stringDateGlobalTpl =
const stringGlobalTpl = `
func stringDateConverter(input []string) (time.Time, error) {
	if len(input) != 1 {
		return time.Time{}, fmt.Errorf("got array '%v' instead of string", input)
	}

	return time.Parse("2006-01-02", input[0])
}

func stringDateTimeConverter(input []string) (time.Time, error) {
	return stringDateConverter(input)
}

func stringYyyyMmDdConverter(input []string) (time.Time, error) {
	if len(input) != 1 {
		return time.Time{}, fmt.Errorf("got array '%v' instead of string", input)
	}

	return time.Parse("2006-01-02", input[0])
}

func stringUnixTimeConverter(input []string) (time.Time, error) {
	if len(input) != 1 {
		return time.Time{}, fmt.Errorf("got array '%v' instead of unix-time", input)
	}

	fromDateSec, err := strconv.ParseInt(input[0], 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(fromDateSec, 0), err
}

func stringEmailConverter(input []string) (string, error) {
	if len(input) != 1 {
		return "", fmt.Errorf("got array '%v' instead of string", input)
	}

	//TODO: write email check

	return input[0], nil
}

func stringBinaryConverter(input []string) ([]byte, error) {
	if len(input) != 1 {
		return []byte{}, fmt.Errorf("got array '%v' instead of string", input)
	}

	return []byte(input[0]), nil
}

func stringConverter(input []string) (string, error) {
	if len(input) != 1 {
		return "", fmt.Errorf("got array '%v' instead of string", input)
	}

	return input[0], nil
}
`

type StringFieldConstructor struct {
}

func (StringFieldConstructor) BuildGlobalCode() (string, error) {
	tpl := template.Must(template.New("stringGlobalTpl").Parse(stringGlobalTpl))
	res := strings.Builder{}
	err := tpl.Execute(&res,
		struct {
			GlobalIdentifier string
		}{""})

	return res.String(), err
}

func (StringFieldConstructor) ImportsRequired() []string {
	return []string{
		"strconv",
	}
}

func (StringFieldConstructor) RegisterAllFormats(res Translator) {

	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.None, func(field BaseField, parentName string) Field {
		field.GoType = "string"
		return &stringField{field}
	})

	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.Date, func(field BaseField, parentName string) Field {
		field.GoType = "time.Time"
		return &stringField{field}
	})

	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.DateTime, func(field BaseField, parentName string) Field {
		field.GoType = "time.Time"
		return &stringField{field}
	})

	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.UnixTime, func(field BaseField, parentName string) Field {
		field.GoType = "time.Time"
		return &stringField{field}
	})

	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.Format("yyyy-mm-dd"), func(field BaseField, parentName string) Field {
		field.GoType = "time.Time"
		return &stringField{field}
	})

	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.Format("email"), func(field BaseField, parentName string) Field {
		field.GoType = "string"
		return &stringField{field}
	})

	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.Binary, func(field BaseField, parentName string) Field {
		field.GoType = "[]byte"
		return &stringField{field}
	})
}