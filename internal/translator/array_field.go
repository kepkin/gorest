package translator

import (
	"fmt"
	"github.com/kepkin/gorest/internal/spec/openapi3"
	"strings"
)

type ArrayFieldImpl struct {
	Field

	Translator Translator
	MakeIdentifier func(string) string
}

func (c *ArrayFieldImpl) BuildDefinition() (string, error) {
	res := strings.Builder{}

	if c.Schema.Items.Ref != "" {
		res.WriteString(fmt.Sprintf("type %v %v\n", c.Name2(), c.GoType))
		return res.String(), nil
	}
	if c.Schema.Items.Type == openapi3.ObjectType {
		typeNameWithoutBrackets := c.GoType[2:]  // GoType always looks like `[]type`
		itemType, err := c.Translator.MakeObjectField(typeNameWithoutBrackets, "", *c.Schema.Items, "parameter?", false)

		itemTypeStr, err := itemType.FieldImpl.BuildDefinition()
		if err != nil {
			return "", fmt.Errorf("сan't build definition for `%v` sub items: %w", c.Name2(), err)
		}
		res.WriteString(itemTypeStr)
	} else {
		res.WriteString(fmt.Sprintf("type %v = %v\n", c.Name2(), c.GoType))
	}

	return res.String(), nil
}

func (c *ArrayFieldImpl) ContextErrorRequired() bool {
	return false
}
