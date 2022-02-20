package translator

import (
	"github.com/iancoleman/strcase"
	"github.com/kepkin/gorest/internal/spec/openapi3"
)

type Field interface {
	Name2() string
	ParameterName() string
	StrVarName() string
	SecondsVarName() string
	IsRequired() bool
	IsCustom() bool
	CheckDefault() bool
	GoTypeString() string
	SchemaType() openapi3.SchemaType

	Fields() []Field

	BuildDefinition() (string, error)
}

// BaseField represents struct field
type BaseField struct {
	Name      string    // UserID
	GoType    string    // int64
	Parameter string    // user_id
	Required  bool
	Type      FieldType // IntegerField
	Schema    openapi3.SchemaType
}

func (f BaseField) Name2() string {
	return f.Name
}

func (f BaseField) ParameterName() string {
	return f.Parameter
}

func (f BaseField) ContextErrorRequired() bool {
	return false
}

func (f BaseField) ImportsRequired() []string {
	return []string{}
}

func (f BaseField) StrVarName() string {
	return strcase.ToLowerCamel(f.Parameter) + "Str"
}

func (f BaseField) SecondsVarName() string {
	return strcase.ToLowerCamel(f.Parameter) + "Sec"
}

func (f BaseField) IsStruct() bool {
	return f.Type == StructField
}

func (f BaseField) IsComponent() bool {
	return f.Type == ComponentField
}

func (f BaseField) IsCustom() bool {
	return f.Type == CustomField
}

func (f BaseField) IsArray() bool {
	return f.Type == ArrayField
}

func (f BaseField) IsBoolean() bool {
	return f.Type == BooleanField
}

func (f BaseField) IsString() bool {
	return f.Type == StringField
}

func (f BaseField) IsInteger() bool {
	return f.Type == IntegerField
}

func (f BaseField) IsFloat() bool {
	return f.Type == FloatField
}

func (f BaseField) IsDate() bool {
	return f.Type == DateField
}

func (f BaseField) IsDateTime() bool {
	return f.Type == DateTimeField
}

func (f BaseField) IsUnixTime() bool {
	return f.Type == UnixTimeField
}

func (f BaseField) IsFile() bool {
	return f.Type == FileField
}

func (f BaseField) IsRequired() bool {
	return f.Required
}

func (f BaseField) CheckDefault() bool {
	//return f.Schema.Default != nil
	return false
}

func (f BaseField) GoTypeString() string {
	return f.GoType
}

func (f BaseField) BuildDefinition() (string, error) {
	return "type " + f.Name + " " + f.GoType + "\n", nil
}

func (f BaseField) Fields() []Field {
	return []Field{}
}

func (f BaseField) SchemaType() openapi3.SchemaType {
	return f.Schema
}
