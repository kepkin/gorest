package translator

import (
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
	FileField
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
	Fields2 []FieldPair

	Level int
	Place string
}

func (d TypeDef) HasNoStringFields() bool {
	for _, f := range d.Fields2 {
		if f.Schema.Type != openapi3.StringType {
			return true
		}
	}
	return false
}

func (d TypeDef) ContextErrorRequired() bool {
	for _, f := range d.Fields2 {
		if f.FieldImpl.ContextErrorRequired() {
			return true
		}
	}
	return false
}

type FieldPair struct {
	Schema    openapi3.SchemaType
	FieldImpl FieldI
}

type ComponentFieldImpl struct {
	Field
}

func (c *ComponentFieldImpl) BuildCode(place string) (string, error) {
	return "Test me ComponentFiledImpl", nil
}

//func (c *ComponentFieldImpl) ImportsRequired() []string {
//	return []string{}
//}

type ObjectFieldConstructor func(f Field, parentName string) FieldI

type Translator struct {
	objectFieldConstructorMap map[openapi3.Type]map[openapi3.Format]ObjectFieldConstructor
	RefResolver func(string) string
}

func (tr *Translator) RegisterObjectFieldConstructor(t openapi3.Type, f openapi3.Format, constructor ObjectFieldConstructor) {
	if tr.objectFieldConstructorMap == nil {
		tr.objectFieldConstructorMap = make(map[openapi3.Type]map[openapi3.Format]ObjectFieldConstructor)
	}

	if _, ok := tr.objectFieldConstructorMap[t]; !ok {
		tr.objectFieldConstructorMap[t] = make(map[openapi3.Format]ObjectFieldConstructor)
	}
	tr.objectFieldConstructorMap[t][f] = constructor
}


func (tr *Translator) MakeObjectField(parentTypeName string, name string, schema openapi3.SchemaType, parameter string, required bool) (FieldPair, error) {
	if schema.Ref != "" {
		return FieldPair{
			schema,
			&Field{
				Name:      name,
				Parameter: parameter,
				Schema:    schema,
				Required:  required,
				GoType:    tr.RefResolver(schema.Ref),
			},
		}, nil

	}

	if constructor, ok := tr.objectFieldConstructorMap[schema.Type][schema.Format]; ok {
		return FieldPair{
			schema,
			constructor(Field{
				Name:      name,
				Parameter: parameter,
				Schema:    schema,
				Required:  required,
			}, parentTypeName),
		}, nil
	}

	//TODO: uncomment
	//return FieldPair{}, fmt.Errorf("field for this format is not defined: '%v:%v'  Please implement http://gorest", schema.Type, schema.Format)
	return FieldPair{}, nil
}
