package translator

import (
	"fmt"
	"github.com/kepkin/gorest/internal/spec/openapi3"
	"sort"
	"strings"
	"text/template"
)

type objectField struct {
	BaseField

	Translator Translator
	MakeIdentifier func(string) string

	fields []Field
}

var structTemplate = template.Must(template.New("struct").Parse(`
type {{ .Name }} struct {
	{{ range $, $field := .Fields -}}
	{{ $field.FieldImpl.Name2 }} {{ $field.FieldImpl.GoType }}
	{{ end }}
}
`))

var structWithJSONTagsTemplate = template.Must(template.New("structWithJSON").Parse(`
type {{ .Name }} struct {
	{{ range $, $field := .Fields -}}
	{{ $field.Name2 }} {{ $field.GoType }}  ` + "`json:\"" + "{{ $field.ParameterName }}\"`" + `
	{{ end }}
}
`))

func (c *objectField) Fields() []Field {
	if c.fields == nil {
		c.fields = make([]Field, 0)

		requiredMap := make(map[string]bool)
		for _, propName := range c.Schema.Required {
			requiredMap[propName] = true
		}

		for propName, propSchema := range c.Schema.Properties {
			propID := c.MakeIdentifier(propName)
			//propSchema.Name = propID // TODO: do I need this?

			isRequired := requiredMap[propName]

			field, err := c.Translator.MakeObjectField(c.GoTypeString(), propID, *propSchema, propName, isRequired)
			if err != nil {
				panic(fmt.Errorf("error occured while building type %v.%v: %w", c.GoTypeString(), propName, err))
			}

			//TODO delme
			if field.FieldImpl != nil {
				c.fields = append(c.fields, field.FieldImpl)
			}
		}

		sort.Slice(c.fields, func(i, j int) bool {
			return c.fields[i].Name2() < c.fields[j].Name2()
		})
	}

	return c.fields
}

//TODO: Think about different definitions
// One is for struct BaseField
// Other one is for defining a Type
func (c *objectField) BuildDefinition() (string, error) {
	if c.GoType == "json.RawMessage" {
		return "", nil
	}

	res := strings.Builder{}
	err := structWithJSONTagsTemplate.Execute(&res,
		struct {
			Name string
			Fields []Field
	}{
		c.GoType,
		c.Fields(),
		})
	if err != nil {
		return "", err
	}

	for _, innerField := range c.Fields() {
		if innerField.SchemaType().Ref != "" {
			continue
		}

		if innerField.SchemaType().Type != openapi3.ObjectType && innerField.SchemaType().Type != openapi3.ArrayType {
			continue
		}

		if innerField.SchemaType().Type == openapi3.ArrayType && innerField.SchemaType().Items.Ref != "" {
			continue
		}

		innerFieldStr, err := innerField.BuildDefinition()
		if err != nil {
			return "", fmt.Errorf("Can't build definition for `%v` cause of field `%v`: %w", c.Name2(), innerField.Name2(), err)
		}
		res.WriteString(innerFieldStr)
	}

	return res.String(), err
}

type ObjectFieldConstructor struct {
}

func (ObjectFieldConstructor) RegisterAllFormats(translator Translator) {
	translator.RegisterObjectFieldConstructor(openapi3.ObjectType, openapi3.None, func(field BaseField, parentName string) Field {
		if field.Schema.Ref != "" {
			field.GoType = translator.RefResolver(field.Schema.Ref)
		} else if field.Schema.AdditionalProperties != nil && field.Schema.AdditionalProperties.Type == openapi3.ObjectType {
			field.GoType = "json.RawMessage"
		} else if field.Schema.AdditionalProperties != nil && field.Schema.AdditionalProperties.Type != openapi3.ObjectType {
			panic(fmt.Sprintf("not implemented `AdditionalProperties` for %v: %v", field.Schema.AdditionalProperties.Type, field.Name2()))
		} else {
			field.GoType = parentName + translator.MakeTitledIdentifier(field.Name)
		}
		return &objectField{BaseField: field, Translator: translator, MakeIdentifier: translator.MakeIdentifier}
	})
}

func (ObjectFieldConstructor) BuildGlobalCode() (string, error) {
	return "", nil
}

func (ObjectFieldConstructor) ImportsRequired() []string {
	return []string{}
}