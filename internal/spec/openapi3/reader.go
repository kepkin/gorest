package openapi3

//import "gopkg.in/yaml.v2"
import "github.com/ghodss/yaml"

type Type string

const (
	ArrayType   Type = "array"
	BooleanType Type = "boolean"
	IntegerType Type = "integer"
	NumberType  Type = "number"
	ObjectType  Type = "object"
	StringType  Type = "string"
)

type Format string

const (
	None Format = ""

	Integer32bit Format = "int32"
	Integer64bit Format = "int64"

	NumberFloat  Format = "float"
	NumberDouble Format = "double"

	Binary   Format = "binary"
	Date     Format = "date"
	DateTime Format = "date-time"
	UnixTime Format = "unix-time"
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
	URL         string `yaml:"url"`
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

const (
	MultiPartFormDataMimeType MimeType = "multipart/form-data"
	ApplicationJsonMimeType MimeType = "application/json"
)

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
	//Enum        []string
	//Default     *string
	Format      Format
	Type        Type
	IsNullable  bool `yaml:"nullable"`

	Ref string `json:"$ref"`

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
	Minimum          float64
	MultipleOf       float64
	ExclusiveMaximum bool
	ExclusiveMinimum bool

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
