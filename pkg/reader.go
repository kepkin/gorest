package pkg

import "gopkg.in/yaml.v2"

const (
	arrayType   = "array"
	booleanType = "boolean"
	integerType = "integer"
	numberType  = "number"
	objectType  = "object"
	stringType  = "string"
)

const (
	integer32bit = "int32"
	integer64bit = "int64"

	numberFloat  = "float"
	numberDouble = "double"
)

func readSpec(in []byte) (res spec, err error) {
	err = yaml.Unmarshal(in, &res)
	return
}

type spec struct {
	OpenAPI    string `yaml:"openapi"`
	Info       infoType
	Servers    []serverType
	Paths      pathMap
	Components componentsType
}

type infoType struct {
	Title   string
	Version string
}

type serverType struct {
	Url         string
	Description string
}

// Paths

type pathMap map[path]pathType

type path = string

type pathType struct {
	Post    *pathSpec
	Patch   *pathSpec
	Get     *pathSpec
	Options *pathSpec
	Put     *pathSpec
	Delete  *pathSpec
}

type pathSpec struct {
	Summary     string
	Description string
	OperationID string `yaml:"operationId"`
	Parameters  []parameterType
	RequestBody *requestBodyType `yaml:"requestBody"`
	// TODO(a.telyshev): Responses
}

type parameterType struct {
	Name     string
	In       string
	Required bool
	Schema   *schemaType
}

type requestBodyType struct {
	Description string
	Required    bool
	Content     map[mimeType]contentType
}

type mimeType = string

type contentType struct {
	Schema *schemaType
}

// Components

type schemaName = string

type componentsType struct {
	Schemas map[schemaName]*schemaType
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
type schemaType struct {
	Type                 string
	Format               string
	Properties           propertiesType
	AdditionalProperties *schemaType `yaml:"additionalProperties"`
	Items                *schemaType
	Ref                  string `yaml:"$ref"`

	Name           string
	GoType         string
	HasCustomType  bool
	HasSpecialType bool
}

type propertiesType map[propertyName]*schemaType

type propertyName = string
