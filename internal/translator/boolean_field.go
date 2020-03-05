package translator

import (
	"strings"
	"text/template"
)

type BooleanFieldImpl struct {
	Field

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

func (c *BooleanFieldImpl) BuildGlobalCode() (string, error) {
	tpl := template.Must(template.New("booleanGlobalTpl").Parse(booleanGlobalTpl))
	res := strings.Builder{}
	err := tpl.Execute(&res,
		struct {
			GlobalIdentifier string
		}{c.ImplIdentifier})

	return res.String(), err
}

func (c *BooleanFieldImpl) ContextErrorRequired() bool {
	return false
}
