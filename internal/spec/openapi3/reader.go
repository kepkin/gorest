package openapi3

import "gopkg.in/yaml.v2"

type Type string

const (
	ArrayType   Type = "array"
	BooleanType      = "boolean"
	IntegerType      = "integer"
	NumberType       = "number"
	ObjectType       = "object"
	StringType       = "string"
)

type Format string

const (
	Integer32bit Format = "int32"
	Integer64bit        = "int64"

	NumberFloat  = "float"
	NumberDouble = "double"

	Date     = "date"
	DateTime = "date-time"
	UnixTime = "unix-time"
)

func ReadSpec(in []byte) (res Spec, err error) {
	err = yaml.Unmarshal(in, &res)
	return
}

type Spec struct {
	OpenAPI    string `yaml:"openapi"`
	Info       InfoType
	Servers    []ServerType
	Paths      PathMap
	Components ComponentsType
}

type InfoType struct {
	Title   string
	Version string
}

type ServerType struct {
	Url         string
	Description string
}

// Paths

type PathMap map[Path]PathType

type Path = string

type PathType struct {
	Get     *PathSpec
	Post    *PathSpec
	Patch   *PathSpec
	Delete  *PathSpec
	Put     *PathSpec
	Options *PathSpec
}

func (p PathType) Methods() []*PathSpec {
	return []*PathSpec{p.Get, p.Post, p.Patch, p.Delete, p.Put, p.Options}
}

type PathSpec struct {
	Summary     string
	Description string
	OperationID string `yaml:"operationId"`
	Parameters  []ParameterType
	RequestBody *RequestBodyType `yaml:"requestBody"`
	// TODO(a.telyshev): Responses
}

type ParameterType struct {
	Name     string
	In       string
	Required bool
	Schema   *SchemaType
}

type RequestBodyType struct {
	Description string
	Required    bool
	Content     map[MimeType]ContentType
}

type MimeType = string

type ContentType struct {
	Schema *SchemaType
}

// Components

type SchemaName = string

type ComponentsType struct {
	Schemas map[SchemaName]*SchemaType
	// TODO(a.telyshev): Parameters
	// TODO(a.telyshev): securitySchemes
	// TODO(a.telyshev): requestBodies
	// TODO(a.telyshev): responses
	// TODO(a.telyshev): headers
	// TODO(a.telyshev): examples
	// TODO(a.telyshev): links
	// TODO(a.telyshev): callbacks
}

// Schema

type SchemaType struct {
	Title       string
	Description string
	// TODO(a.telyshev) Enum
	// TODO(a.telyshev) Default
	Format     Format
	Type       Type
	IsNullable bool `yaml:"nullable"`

	Ref string `yaml:"$ref"`

	ArraySchema  `yaml:",inline"`
	StringSchema `yaml:",inline"`
	NumberSchema `yaml:",inline"`
	ObjectSchema `yaml:",inline"`

	// Service fields
	Name string
}

type ArraySchema struct {
	Items       *SchemaType
	MaxItems    int
	MinItems    int
	UniqueItems bool
}

type StringSchema struct {
	Pattern   string
	MaxLength int
	MinLength int
}

type NumberSchema struct {
	Maximum          float64
	ExclusiveMaximum bool
	Minimum          float64
	ExclusiveMinimum bool
	MultipleOf       float64

	// Service fields
	BitSize int
}

type ObjectSchema struct {
	Required             []PropertyName
	Properties           PropertiesType
	AdditionalProperties *SchemaType `yaml:"additionalProperties"`
}

type PropertiesType map[PropertyName]*SchemaType

type PropertyName = string

func NewObjectSchema(name string) SchemaType {
	return SchemaType{
		Type:         ObjectType,
		ObjectSchema: ObjectSchema{Properties: make(PropertiesType)},

		Name: name,
	}
}
