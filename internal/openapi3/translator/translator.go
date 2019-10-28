package translator

import (
	"container/list"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/kepkin/gorest/internal/openapi3/spec"
)

type ParameterPlace int

const (
	Path ParameterPlace = iota
	Query
	Header
	Cookie
)

type TypeDef struct {
	Name   string
	Type   string
	Fields []interface{}

	Level int
	Place string
}

type InterfaceCheckerDef struct {
	TypeName      string
	InterfaceName string
}

type Field struct {
	Name      string
	Type      string
	Parameter string

	Level int
	Place string
}

type ArrayField struct {
	Field
}

type BooleanField struct {
	Field
}

type StructField struct {
	Field
}

type IntegerField struct {
	Field
	BitSize int
}

type FloatField struct {
	Field
	BitSize int
}

type StringField struct {
	Field
	Format string
}

func ProcessRootSchema(schema spec.SchemaType) ([]interface{}, error) {
	queue := list.New()

	schema.Level = 0
	queue.PushBack(schema)

	result := make([]interface{}, 0)
	for {
		el := queue.Back()
		if el == nil {
			break
		}
		queue.Remove(el)

		switch t := el.Value.(type) {
		case spec.SchemaType:
			def, err := processObjSchema(t, queue)
			if err != nil {
				return nil, err
			}
			result = append(result, def)

		default:
			result = append(result, t)
		}
	}
	return result, nil
}

func processObjSchema(schema spec.SchemaType, queue *list.List) (def TypeDef, err error) {
	if schema.Type != spec.ObjectType {
		err = fmt.Errorf("schema must be `object`, got: %s", schema.Type)
		return
	}

	def.Name = MakeIdentifier(schema.Name)
	def.Level = schema.Level
	def.Place = schema.Place

	for propName, propSchema := range schema.Properties {
		propID := MakeIdentifier(propName)
		propSchema.Name = propID

		var field interface{}
		field, err = determineType(def.Name, schema.Level, *propSchema, propName, schema.Place, queue)
		if err != nil {
			return
		}
		def.Fields = append(def.Fields, field)
	}
	return
}

func determineType(parentName string, currentLevel int, schema spec.SchemaType, parameter string, place string, queue *list.List) (interface{}, error) {
	if schema.Ref != "" {
		return StructField{Field{
			Name:      schema.Name,
			Parameter: parameter,
			Type:      GetNameFromRef(schema.Ref),
			Level:     currentLevel,
		}}, nil
	}

	switch schema.Type {
	case spec.ArrayType:
		childName := parentName + MakeTitledIdentifier(schema.Name)
		t, err := determineType(childName, currentLevel+1, *schema.Items, parameter, place, queue)
		if err != nil {
			return "", err
		}
		return ArrayField{Field{
			Name:      schema.Name,
			Parameter: parameter,
			Type:      "[]" + t.(Field).Name,
			Level:     currentLevel,
			Place:     place,
		}}, nil

	case spec.BooleanType:
		return BooleanField{Field{
			Name:      schema.Name,
			Parameter: parameter,
			Type:      "bool",
			Level:     currentLevel,
			Place:     place,
		}}, nil

	case spec.IntegerType:
		var type_ string
		var bitSize int

		switch schema.Format {
		case "":
			type_ = "int"
			bitSize = 0

		case spec.Integer32bit:
			type_ = "int32"
			bitSize = 32

		case spec.Integer64bit:
			type_ = "int64"
			bitSize = 64

		default:
			type_ = MakeTitledIdentifier(schema.Format)
			queue.PushBack(InterfaceCheckerDef{
				TypeName:      type_,
				InterfaceName: "json.Marshaller",
			})
			fmt.Printf("please implement own integer type `%s`\n", type_)
		}
		return IntegerField{
			Field: Field{
				Name:      schema.Name,
				Parameter: parameter,
				Type:      type_,
				Level:     currentLevel,
				Place:     place,
			},
			BitSize: bitSize,
		}, nil

	case spec.NumberType:
		var type_ string
		var bitSize int

		switch schema.Format {
		case "":
			type_ = "float"
			bitSize = 0

		case spec.NumberFloat:
			type_ = "float32"

		case spec.NumberDouble:
			type_ = "float64"

		default:
			type_ = MakeTitledIdentifier(schema.Format)
			queue.PushBack(InterfaceCheckerDef{
				TypeName:      type_,
				InterfaceName: "json.Marshaller",
			})
			fmt.Printf("please implement own number type `%s`\n", type_)
		}
		return FloatField{
			Field: Field{
				Name:      schema.Name,
				Parameter: parameter,
				Type:      type_,
				Level:     currentLevel,
				Place:     place,
			},
			BitSize: bitSize,
		}, nil

	case spec.ObjectType:
		name := schema.Name
		type_ := parentName + MakeTitledIdentifier(schema.Name)

		schema.Name = type_
		schema.Level = currentLevel + 1
		queue.PushBack(schema)

		return StructField{Field{
			Name:  name,
			Type:  type_,
			Level: currentLevel,
			Place: place,
		}}, nil

	case spec.StringType:
		return StringField{
			Field: Field{
				Type:      "string",
				Name:      schema.Name,
				Parameter: parameter,
				Level:     currentLevel,
				Place:     place,
			},
			Format: schema.Format,
		}, nil

	default:
		return nil, fmt.Errorf("unknown data type: %v", schema.Type)
	}
}

func MakeIdentifier(s string) string {
	return strcase.ToCamel(strings.ReplaceAll(s, " ", "_"))
}

func MakeTitledIdentifier(s string) string {
	return strings.Title(MakeIdentifier(s))
}

func GetNameFromRef(s string) string {
	return s[len("#/components/schemas/"):]
}
