package translator

import (
	"container/list"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/kepkin/gorest/internal/spec/openapi3"
)

type FieldType int

const (
	UnknownField FieldType = iota
	ArrayField
	BooleanField
	ComponentField
	CustomField
	FloatField
	FreeFormObject
	IntegerField
	StringField
	StructField
)

type TypeDef struct {
	Name   string
	GoType string
	Fields []Field

	Level int
	Place string
}

func (d TypeDef) HasNoStringFields() bool {
	for _, f := range d.Fields {
		if f.Type != StringField {
			return true
		}
	}
	return false
}

// Field represents struct field
type Field struct {
	Name      string    // UserID
	GoType    string    // int64
	Parameter string    // user_id
	Type      FieldType // IntegerField
	Schema    openapi3.SchemaType
}

func (f Field) StrVarName() string {
	return strcase.ToLowerCamel(f.Parameter) + "Str"
}

func (f Field) IsString() bool {
	return f.Type == StringField
}

func (f Field) IsInteger() bool {
	return f.Type == IntegerField
}

func (f Field) IsFloat() bool {
	return f.Type == FloatField
}

func (f Field) IsCustom() bool {
	return f.Type == CustomField
}

func ProcessRootSchema(schema openapi3.SchemaType) ([]TypeDef, error) {
	queue := list.New()
	queue.PushBack(schema)

	result := make([]TypeDef, 0)
	for {
		el := queue.Back()
		if el == nil {
			break
		}
		queue.Remove(el)

		if sch, ok := el.Value.(openapi3.SchemaType); ok {
			def, err := ProcessObjSchema(sch, queue)
			if err != nil {
				return nil, err
			}
			result = append(result, def)

		} else {
			return nil, fmt.Errorf("unprocessible entity: %v", el.Value)
		}
	}
	return result, nil
}

func ProcessObjSchema(schema openapi3.SchemaType, queue *list.List) (def TypeDef, err error) {
	if schema.Type != openapi3.ObjectType {
		err = fmt.Errorf("schema must be `object`, got: %s", schema.Type)
		return
	}

	def.Name = MakeIdentifier(schema.Name)
	def.GoType = "struct" // For debug usage only

	for propName, propSchema := range schema.Properties {
		propID := MakeIdentifier(propName)
		propSchema.Name = propID

		var field Field
		field, err = determineType(def.Name, *propSchema, propName, queue)
		if err != nil {
			return
		}
		def.Fields = append(def.Fields, field)
	}
	return
}

func determineType(parentName string, schema openapi3.SchemaType, parameter string, queue *list.List) (Field, error) {
	if schema.Ref != "" {
		return Field{
			Type:      ComponentField,
			Name:      schema.Name,
			Parameter: parameter,
			GoType:    GetNameFromRef(schema.Ref),
		}, nil
	}

	switch schema.Type {
	case openapi3.ArrayType:
		childName := parentName + MakeTitledIdentifier(schema.Name)
		t, err := determineType(childName, *schema.Items, parameter, queue)
		if err != nil {
			return Field{}, err
		}
		return Field{
			Type:      ArrayField,
			Name:      schema.Name,
			Parameter: parameter,
			GoType:    "[]" + t.GoType,
		}, nil

	case openapi3.BooleanType:
		return Field{
			Type:      BooleanField,
			Name:      schema.Name,
			Parameter: parameter,
			GoType:    "bool",
		}, nil

	case openapi3.IntegerType:
		switch schema.Format {
		case "":
			schema.BitSize = 0

		case openapi3.Integer32bit:
			schema.BitSize = 32

		case openapi3.Integer64bit:
			schema.BitSize = 64

		default:
			type_ := MakeTitledIdentifier(schema.Format)
			return Field{
				Type:      CustomField,
				Name:      schema.Name,
				Parameter: parameter,
				GoType:    type_,
				Schema:    schema,
			}, nil
		}
		return Field{
			Type:      IntegerField,
			Name:      schema.Name,
			Parameter: parameter,
			GoType:    "int64",
			Schema:    schema,
		}, nil

	case openapi3.NumberType:
		switch schema.Format {
		case "":
			schema.BitSize = 0

		case openapi3.NumberFloat:
			schema.BitSize = 32

		case openapi3.NumberDouble:
			schema.BitSize = 64

		default:
			type_ := MakeTitledIdentifier(schema.Format)
			return Field{
				Type:      CustomField,
				Name:      schema.Name,
				Parameter: parameter,
				GoType:    type_,
				Schema:    schema,
			}, nil
		}
		return Field{
			Type:      FloatField,
			Name:      schema.Name,
			Parameter: parameter,
			GoType:    "int64",
			Schema:    schema,
		}, nil

	case openapi3.ObjectType:
		if len(schema.ObjectSchema.Properties) == 0 &&
			(schema.AdditionalProperties == nil || len(schema.AdditionalProperties.Properties) == 0) {

			return Field{
				Type:   FreeFormObject,
				Name:   schema.Name,
				GoType: "json.RawMessage",
			}, nil
		}

		name := schema.Name
		type_ := parentName + MakeTitledIdentifier(schema.Name)

		schema.Name = type_

		queue.PushBack(schema)
		return Field{
			Type:   StructField,
			Name:   name,
			GoType: type_,
		}, nil

	case openapi3.StringType:
		return Field{
			Type:      StringField,
			GoType:    "string",
			Name:      schema.Name,
			Parameter: parameter,
		}, nil

	default:
		return Field{}, fmt.Errorf("unknown data type: %v", schema.Type)
	}
}

func MakeIdentifier(s string) string {
	result := strcase.ToCamel(strings.ReplaceAll(s, " ", "_"))
	if strings.HasSuffix(result, "Id") {
		result = result[:len(result)-2] + "ID"
	}
	return result
}

func MakeTitledIdentifier(s string) string {
	return strings.Title(MakeIdentifier(s))
}

func GetNameFromRef(s string) string {
	return s[len("#/components/schemas/"):]
}
