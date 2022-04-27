package translator

import (
	"fmt"
	"github.com/kepkin/gorest/internal/spec/openapi3"
	"strings"
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

type FieldPair struct {
	Schema    openapi3.SchemaType
	FieldImpl Field
}

type FieldConstructor func(f BaseField, parentName string) Field

type Translator struct {
	objectFieldConstructorMap map[openapi3.Type]map[openapi3.Format]FieldConstructor
	RefResolver func(string) string
	MakeTitledIdentifier func(string) string
	MakeIdentifier func(string) string
	uniqFields []FieldConstructor2
}

type FieldConstructor2 interface {
	RegisterAllFormats(translator Translator)
	BuildGlobalCode() (string, error)
	ImportsRequired() []string
}

func (tr *Translator) RegisterField(f FieldConstructor2) {
	if tr.objectFieldConstructorMap == nil {
		tr.objectFieldConstructorMap = make(map[openapi3.Type]map[openapi3.Format]FieldConstructor)
	}

	f.RegisterAllFormats(*tr)
	tr.uniqFields = append(tr.uniqFields, f)
}

func (tr *Translator) RegisterObjectFieldConstructor(t openapi3.Type, f openapi3.Format, constructor FieldConstructor) {
	if tr.objectFieldConstructorMap == nil {
		tr.objectFieldConstructorMap = make(map[openapi3.Type]map[openapi3.Format]FieldConstructor)
	}

	if _, ok := tr.objectFieldConstructorMap[t]; !ok {
		tr.objectFieldConstructorMap[t] = make(map[openapi3.Format]FieldConstructor)
	}
	tr.objectFieldConstructorMap[t][f] = constructor
}

func (tr *Translator) ImportsRequired() (res []string) {
	for _, f := range tr.uniqFields {
		res = append(res, f.ImportsRequired()...)
	}
	return
}

func (tr *Translator) BuildGlobalCode() (string, error) {
	res := strings.Builder{}
	for _, f := range tr.uniqFields {
		fstr, err := f.BuildGlobalCode()
		if err != nil {
			return "", err
		}
		_, err = res.WriteString(fstr)
		if err != nil {
			return "", err
		}
	}

	return res.String(), nil
}

func (tr *Translator) MakeObjectField(parentTypeName string, name string, schema openapi3.SchemaType, parameter string, required bool) (FieldPair, error) {
	if schema.Ref != "" {
		//TODO: maybe not needed anymore? If still needed, write clarifying comment
		return FieldPair{
			schema,
			&BaseField{
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
			constructor(BaseField{
				Name:      name,
				Parameter: parameter,
				Schema:    schema,
				Required:  required,
			}, parentTypeName),
		}, nil
	}

	//TODO: write dev guide for writing your own type
	return FieldPair{}, fmt.Errorf("field for this format is not defined: '%v: %v(%v)'. Please implement http://gorest", parameter, schema.Type, schema.Format)
}
