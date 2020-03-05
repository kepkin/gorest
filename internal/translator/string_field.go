package translator

import (
	"strings"
	"text/template"
)

type StringFieldImpl struct {
	Field
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

func (c *StringFieldImpl) BuildGlobalCode() (string, error) {
	tpl := template.Must(template.New("stringGlobalTpl").Parse(stringGlobalTpl))
	res := strings.Builder{}
	err := tpl.Execute(&res,
		struct {
			GlobalIdentifier string
		}{""})

	return res.String(), err
}

func (c *StringFieldImpl) ContextErrorRequired() bool {
	return false
}

func (c *StringFieldImpl) ImportsRequired() []string {
	return []string{
		"strconv",
	}
}
