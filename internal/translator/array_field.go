package translator

import (
	"fmt"
	"strings"

	"github.com/kepkin/gorest/internal/spec/openapi3"
)

type arrayField struct {
	BaseField

	Translator Translator
	MakeIdentifier func(string) string
}

func (c *arrayField) BuildDefinition() (string, error) {
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

func (c *arrayField) RegisterAllFormats(translator Translator) {
	translator.RegisterObjectFieldConstructor(openapi3.BooleanType, openapi3.Format(""), func(field BaseField, parentName string) Field {
		field.GoType = "bool"
		return &booleanField{field, "a"}
	})
}


type ArrayFieldConstructor struct {
}

func (ArrayFieldConstructor) BuildGlobalCode() (string, error) {
	return "", nil
}

func (ArrayFieldConstructor) ImportsRequired() []string {
	return []string{}
}

func (ArrayFieldConstructor) RegisterAllFormats(res Translator) {
	res.RegisterObjectFieldConstructor(openapi3.ArrayType, openapi3.None, func(field BaseField, parentName string) Field {
		if field.Schema.Items.Ref != "" {
			field.GoType = "[]" + res.RefResolver(field.Schema.Items.Ref)
		} else if field.Schema.Items.Type == openapi3.ObjectType {
			field.GoType = "[]" + parentName + res.MakeTitledIdentifier(field.Name) + "Item"
		} else {
			field.GoType = "[]" + string(field.Schema.Items.Type)
		}

		return &arrayField{BaseField: field, Translator: res, MakeIdentifier: res.MakeIdentifier}
	})
}