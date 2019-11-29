package translator

import (
	"container/list"
	"fmt"
	"sort"
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
	DateField
	DateTimeField
	FloatField
	FreeFormObject
	IntegerField
	StringField
	StructField
	UnixTimeField
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

func (f Field) SecondsVarName() string {
	return strcase.ToLowerCamel(f.Parameter) + "Sec"
}

func (f Field) IsCustom() bool {
	return f.Type == CustomField
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

func (f Field) IsDate() bool {
	return f.Type == DateField
}

func (f Field) IsDateTime() bool {
	return f.Type == DateTimeField
}

func (f Field) IsUnixTime() bool {
	return f.Type == UnixTimeField
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

	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result, nil
}

func ProcessObjSchema(schema openapi3.SchemaType, queue *list.List) (def TypeDef, err error) {
	if schema.Type != openapi3.ObjectType {
		// TODO(a.telyshev): More complex processing
		// err = fmt.Errorf("schema must be `object`, got: `%s`", schema.Type)
		def.Name = MakeIdentifier(schema.Name)

		var goType string

		switch schema.Type {
		case openapi3.BooleanType:
			goType = "bool"
		case openapi3.IntegerType:
			goType = "int64"
		case openapi3.NumberType:
			goType = "float64"
		case openapi3.StringType:
			goType = "string"
		case openapi3.ArrayType:
			goType = "[]" + GetNameFromRef(schema.Items.Ref)
		default:
			err = fmt.Errorf("unsupported type: `%s` of schema %v", schema.Type, schema)
			return
		}

		def.GoType = goType
		return
	}

	def.Name = MakeIdentifier(schema.Name)
	def.GoType = "struct"

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
	sort.Slice(def.Fields, func(i, j int) bool {
		return def.Fields[i].Name < def.Fields[j].Name
	})
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
			type_ := MakeTitledIdentifier(string(schema.Format))
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
			type_ := MakeTitledIdentifier(string(schema.Format))
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
				Type:      FreeFormObject,
				Parameter: parameter,
				Name:      schema.Name,
				GoType:    "json.RawMessage",
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
		type_ := StringField
		goType := "string"

		switch schema.Format {
		case openapi3.Date:
			type_ = DateField
			goType = "time.Time"

		case openapi3.DateTime:
			type_ = DateTimeField
			goType = "time.Time"

		case openapi3.UnixTime:
			type_ = UnixTimeField
			goType = "time.Time"
		}

		return Field{
			Type:      type_,
			GoType:    goType,
			Name:      schema.Name,
			Parameter: parameter,
		}, nil

	default:
		return Field{}, fmt.Errorf("%s.%s: unknown data type: %v", parentName, parameter, schema.Type)
	}
}

func MakeIdentifier(s string) string {
	result := strcase.ToCamel(strings.ReplaceAll(s, " ", "_"))

	for _, suff := range [...]string{
		"Db",
		"Http",
		"Id",
		"Inn",
		"Json",
		"Sql",
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
