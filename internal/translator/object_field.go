package translator

import (
	"fmt"
	"github.com/kepkin/gorest/internal/spec/openapi3"
	"sort"
	"strings"
	"text/template"
)

type ObjectFieldImpl struct {
	Field

	Translator Translator
	MakeIdentifier func(string) string

	fields []FieldI
}

func (c *ObjectFieldImpl) BuildGlobalCode() (string, error) {
	return "", nil
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

func (c *ObjectFieldImpl) Fields() []FieldI {
	if c.fields == nil {
		c.fields = make([]FieldI, 0)

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
				panic(err)
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
// One is for struct Field
// Other one is for defining a Type
func (c *ObjectFieldImpl) BuildDefinition() (string, error) {
	if c.GoType == "json.RawMessage" {
		return "", nil
	}

	res := strings.Builder{}
	err := structWithJSONTagsTemplate.Execute(&res,
		struct {
			Name string
			Fields []FieldI
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

func (c *ObjectFieldImpl) ContextErrorRequired() bool {
	return false
}
