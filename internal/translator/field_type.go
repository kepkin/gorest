package translator

import (
	"github.com/iancoleman/strcase"
	"github.com/kepkin/gorest/internal/spec/openapi3"
)

type FieldI interface {
	Name2() string
	ParameterName() string
	StrVarName() string
	SecondsVarName() string
	IsRequired() bool
	IsCustom() bool
	CheckDefault() bool
	GoTypeString() string
	SchemaType() openapi3.SchemaType

	Fields() []FieldI

	ContextErrorRequired() bool
	ImportsRequired() []string
	BuildGlobalCode() (string, error)
	BuildDefinition() (string, error)
}

// Field represents struct field
type Field struct {
	Name      string    // UserID
	GoType    string    // int64
	Parameter string    // user_id
	Required  bool
	Type      FieldType // IntegerField
	Schema    openapi3.SchemaType
}

func (f Field) Name2() string {
	return f.Name
}

func (f Field) ParameterName() string {
	return f.Parameter
}

func (f Field) ContextErrorRequired() bool {
	return false
}

func (f Field) ImportsRequired() []string {
	return []string{}
}

func (f Field) StrVarName() string {
	return strcase.ToLowerCamel(f.Parameter) + "Str"
}

func (f Field) SecondsVarName() string {
	return strcase.ToLowerCamel(f.Parameter) + "Sec"
}

func (f Field) IsStruct() bool {
	return f.Type == StructField
}

func (f Field) IsComponent() bool {
	return f.Type == ComponentField
}

func (f Field) IsCustom() bool {
	return f.Type == CustomField
}

func (f Field) IsArray() bool {
	return f.Type == ArrayField
}

func (f Field) IsBoolean() bool {
	return f.Type == BooleanField
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

func (f Field) IsFile() bool {
	return f.Type == FileField
}

func (f Field) IsRequired() bool {
	return f.Required
}

func (f Field) CheckDefault() bool {
	//return f.Schema.Default != nil
	return false
}

func (f Field) GoTypeString() string {
	return f.GoType
}

func (c Field) BuildGlobalCode() (string, error) {
	return "", nil
}

func (c Field) BuildDefinition() (string, error) {
	return "type " + c.Name + " " + c.GoType + "\n", nil
}

func (c Field) Fields() []FieldI {
	return []FieldI{}
}

func (c Field) SchemaType() openapi3.SchemaType {
	return c.Schema
}
