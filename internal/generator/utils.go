package generator

import (
	"container/list"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/kepkin/gorest/internal/spec/openapi3"
	"github.com/kepkin/gorest/internal/translator"
	"sort"
	"strings"
)

// Giving an open API object, it will return several GO struct definitions
func (g *Generator) MakeAllTypeDefsFromOpenAPIObject(schema openapi3.SchemaType) ([]translator.Field, error) {
	dependantOpenAPIObjects := list.New()
	dependantOpenAPIObjects.PushBack(schema)

	result := make([]translator.Field, 0)
	for {
		el := dependantOpenAPIObjects.Back()
		if el == nil {
			break
		}
		dependantOpenAPIObjects.Remove(el)

		if sch, ok := el.Value.(openapi3.SchemaType); ok {
			def, err := g.MakeTypeDefFromOpenAPIObject(sch, dependantOpenAPIObjects)
			if err != nil {
				return nil, err
			}


			result = append(result, def)
		} else {
			return nil, fmt.Errorf("unprocessible entity: %v", el.Value)
		}
	}

	sort.Slice(result, func(i, j int) bool { return result[i].Name2() < result[j].Name2() })
	return result, nil
}

// Will build internal type definition by reading an open api Schema object. Queue will have all schemas, on which this
// scheme depends
func (g *Generator) MakeTypeDefFromOpenAPIObject(schema openapi3.SchemaType, queue *list.List) (def translator.Field, err error) {
	fieldPair, err := g.translator.MakeObjectField("", MakeIdentifier(schema.Name), schema, "", false)

	return fieldPair.FieldImpl, err
}

func MakeFormatIdentifier(s openapi3.Format) string {
	return MakeIdentifier(string(s))
}

func MakeIdentifier(s string) string {
	result := strcase.ToCamel(strings.ReplaceAll(s, " ", "_"))

	for _, suff := range [...]string{
		"Api",
		"Edo",
		"Db",
		"Http",
		"Id",
		"Inn",
		"Json",
		"Sql",
		"Uid",
		"Url",
	} {
		if strings.HasSuffix(result, suff) {
			result = result[:len(result)-len(suff)] + strings.ToUpper(suff)
			break
		}
	}
	return result
}

func MakeTitledIdentifier(s string) string {
	return strings.Title(MakeIdentifier(s))
}

func GetNameFromRef(s string) string {
	return s[len("#/components/schemas/"):]
}
