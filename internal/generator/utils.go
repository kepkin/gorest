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
func (g *Generator) MakeAllTypeDefsFromOpenAPIObject(schema openapi3.SchemaType) ([]translator.FieldI, error) {
	dependantOpenAPIObjects := list.New()
	dependantOpenAPIObjects.PushBack(schema)

	result := make([]translator.FieldI, 0)
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
func (g *Generator) MakeTypeDefFromOpenAPIObject(schema openapi3.SchemaType, queue *list.List) (def translator.FieldI, err error) {
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

func MakeTranslator() (res translator.Translator) {
	res.RefResolver = GetNameFromRef

	res.RegisterObjectFieldConstructor(openapi3.BooleanType, openapi3.Format(""), func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "bool"
		return &translator.BooleanFieldImpl{field, "a"}
	})

	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.None,  func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "string"
		return &translator.StringFieldImpl{field}
	})

	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.Date,  func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "time.Time"
		return &translator.StringFieldImpl{field}
	})

	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.DateTime,  func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "time.Time"
		return &translator.StringFieldImpl{field}
	})

	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.UnixTime,  func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "time.Time"
		return &translator.StringFieldImpl{field}
	})

	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.Format("yyyy-mm-dd"),  func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "time.Time"
		return &translator.StringFieldImpl{field}
	})

	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.Format("email"),  func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "string"
		return &translator.StringFieldImpl{field}
	})
	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.Binary,  func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "[]byte"
		return &translator.StringFieldImpl{field}
	})

	res.RegisterObjectFieldConstructor(openapi3.ObjectType, openapi3.None, func(field translator.Field, parentName string) translator.FieldI {
		if field.Schema.Ref != "" {
			field.GoType = GetNameFromRef(field.Schema.Ref)
		} else if field.Schema.AdditionalProperties != nil && field.Schema.AdditionalProperties.Type == openapi3.ObjectType {
			 field.GoType = "json.RawMessage"
		} else if field.Schema.AdditionalProperties != nil && field.Schema.AdditionalProperties.Type != openapi3.ObjectType {
			panic(fmt.Sprintf("not implemented `AdditionalProperties` for %v: %v", field.Schema.AdditionalProperties.Type, field.Name2()))
		} else {
			field.GoType = parentName + MakeTitledIdentifier(field.Name)
		}
		return &translator.ObjectFieldImpl{Field: field, Translator: res, MakeIdentifier: MakeIdentifier}
	})

	res.RegisterObjectFieldConstructor(openapi3.ArrayType, openapi3.None, func(field translator.Field, parentName string) translator.FieldI {
		if field.Schema.Items.Ref != "" {
			field.GoType = "[]" + GetNameFromRef(field.Schema.Items.Ref)
		} else if field.Schema.Items.Type == openapi3.ObjectType {
			field.GoType = "[]" + parentName + MakeTitledIdentifier(field.Name) + "Item"
		} else {
			field.GoType = "[]" + string(field.Schema.Items.Type)
		}

		return &translator.ArrayFieldImpl{Field: field, Translator: res, MakeIdentifier: MakeIdentifier}
	})

	res.RegisterObjectFieldConstructor(openapi3.IntegerType, openapi3.None, func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "int64"
		return &translator.IntegerFieldImpl{Field: field}
	})

	res.RegisterObjectFieldConstructor(openapi3.IntegerType, openapi3.Integer32bit, func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "int32"
		return &translator.IntegerFieldImpl{Field: field}
	})

	res.RegisterObjectFieldConstructor(openapi3.IntegerType, openapi3.Integer64bit, func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "int64"
		return &translator.IntegerFieldImpl{Field: field}
	})

	res.RegisterObjectFieldConstructor(openapi3.NumberType, openapi3.None, func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "int64"
		return &translator.IntegerFieldImpl{Field: field}
	})

	res.RegisterObjectFieldConstructor(openapi3.StringType, openapi3.Format("decimal"), func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "decimal.Decimal"
		return &translator.DecimalFieldImpl{Field: field}
	})

	res.RegisterObjectFieldConstructor(openapi3.NumberType, openapi3.Format("decimal"), func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "decimal.Decimal"
		return &translator.DecimalFieldImpl{Field: field}
	})

	res.RegisterObjectFieldConstructor(openapi3.IntegerType, openapi3.Format("decimal"), func(field translator.Field, parentName string) translator.FieldI {
		field.GoType = "decimal.Decimal"
		return &translator.DecimalFieldImpl{Field: field}
	})


	return
}
