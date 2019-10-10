package pkg

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/scanner"
	"go/token"
	"io"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/iancoleman/strcase"
	"golang.org/x/tools/go/ast/astutil"
	"gopkg.in/yaml.v2"
)

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

type InfoType struct {
	Title   string
	Version string
}

type ServerType struct {
	Url         string
	Description string
}

type PropertyName = string

type SchemaType struct {
	Type                 string
	Format               string
	Properties           map[PropertyName]SchemaType
	AdditionalProperties *SchemaType `yaml:"additionalProperties"`
	Items                *SchemaType
	Ref                  string `yaml:"$ref"`

	GoType        string
	HasCustomType bool
}

func (s *SchemaType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type SchemaTypeYAML SchemaType // To avoid a recursive unmarshalling
	err := unmarshal((*SchemaTypeYAML)(s))
	if err != nil {
		return err
	}

	if err = DetermineType(s); err != nil {
		return err
	}
	return nil
}

type MIMEType = string

type ContentType struct {
	Schema SchemaType
}

type RequestBodyType struct {
	Description string
	Required    bool
	Content     map[MIMEType]ContentType
}

type ParameterType struct {
	Name     string
	In       string
	Required bool
	Schema   SchemaType
}

type PathSpec struct {
	Summary     string
	Description string
	OperationID string `yaml:"operationId"`
	Parameters  []ParameterType
	RequestBody RequestBodyType `yaml:"requestBody"`
	// TODO(a.telyshev): Responses
}

type Path = string

type PathType struct {
	Post    *PathSpec
	Patch   *PathSpec
	Get     *PathSpec
	Options *PathSpec
	Put     *PathSpec
	Delete  *PathSpec
}

type PathMap map[Path]PathType

type SchemaName = string

// ComponentsType store components described at https://swagger.io/docs/specification/components/
type ComponentsType struct {
	Schemas map[SchemaName]SchemaType
	// TODO(a.telyshev): Parameters
	// TODO(a.telyshev): securitySchemes
	// TODO(a.telyshev): requestBodies
	// TODO(a.telyshev): responses
	// TODO(a.telyshev): headers
	// TODO(a.telyshev): examples
	// TODO(a.telyshev): links
	// TODO(a.telyshev): callbacks
}

type Spec struct {
	OpenAPI    string `yaml:"openapi"`
	Info       InfoType
	Servers    []ServerType
	Paths      PathMap
	Components ComponentsType
}

func ReadSpec(in []byte) (res Spec, err error) {
	err = yaml.Unmarshal(in, &res)
	return
}

func BuildTemplates() (res *template.Template, err error) {
	a := sprig.GenericFuncMap()
	res = template.New("").Funcs(a)
	res = res.Funcs(template.FuncMap{
		"MakeIdentifier":    MakeIdentifier,
		"GetNameFromRef":    GetNameFromRef,
		"ToConstructorType": ToConstructorType,
		"ConvertUrl":        ConvertUrl,
	})

	dir, err := Assets.Open("/")
	if err != nil {
		return
	}
	files, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	if len(files) == 0 {
		return res, fmt.Errorf("gorest: no template files")
	}

	for _, f := range files {
		fData, err := Assets.Open(f.Name())
		if err != nil {
			return res, err
		}
		data, err := ioutil.ReadAll(fData)
		if err != nil {
			return res, err
		}

		res, err = res.New(f.Name()).Parse(string(data))
		if err != nil {
			return res, err
		}
	}

	return
}

func GenerateFromFile(swaggerPath string, packageName string, wr io.Writer) error {
	content, err := ioutil.ReadFile(swaggerPath)
	if err != nil {
		return err
	}
	spec, err := ReadSpec(content)
	if err != nil {
		return err
	}

	return GenerateFromSpec(spec, packageName, wr)
}

func GenerateFromSpec(spec Spec, packageName string, wr io.Writer) error {
	t, err := BuildTemplates()
	if err != nil {
		return err
	}

	specialTypes := map[string]SchemaType{}
	for schemaName, schema := range spec.Components.Schemas {
		for propName, propSchema := range schema.Properties {
			if propSchema.Type == "object" {
				normalizedPropName := MakeIdentifier(propName)
				newTypeName := schemaName + normalizedPropName + "Type"
				specialTypes[newTypeName] = propSchema
				schema.Properties[propName] = SchemaType{Ref: "#/components/schemas/" + newTypeName}
			}
		}
	}

	for k, v := range specialTypes {
		spec.Components.Schemas[k] = v
	}
	tmpResult := strings.Builder{}
	err = t.ExecuteTemplate(&tmpResult, "main.tmpl", map[string]interface{}{"package": packageName, "spec": spec})
	if err != nil {
		return err
	}
	err = finalizeGoSource(tmpResult.String(), wr)
	return err
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

func DetermineType(spec *SchemaType) error {
	if spec.Ref != "" {
		spec.GoType = GetNameFromRef(spec.Ref)
		return nil
	}

	var type_ string

	switch spec.Type {
	case arrayType:
		type_ = "[]" + GetNameFromRef(spec.Items.Ref)

	case booleanType:
		type_ = "bool"

	case integerType:
		switch spec.Format {
		case "":
			type_ = "int" // Integer numbers

		case integer32bit:
			type_ = "int32" // Signed 32-bit integers (commonly used integer type)

		case integer64bit:
			type_ = "int64" // Signed 64-bit integers (long type)

		default:
			type_ = MakeTitledIdentifier(spec.Format)
			spec.HasCustomType = true
			fmt.Printf("please implement own integer type `%s`\n", type_)
		}

	case numberType:
		switch spec.Format {
		case "":
			type_ = "float" // Any numbers

		case numberFloat:
			type_ = "float32" // Floating-point numbers

		case numberDouble:
			type_ = "float64" // Floating-point numbers with double precision

		default:
			type_ = MakeTitledIdentifier(spec.Format)
			spec.HasCustomType = true
			fmt.Printf("please implement own number type `%s`\n", type_)
		}

	case objectType:
		break

	case stringType:
		// TODO(a.telyshev): Support format
		type_ = "string"

	default:
		return fmt.Errorf("unknown data type: %v", spec.Type)
	}

	spec.GoType = type_
	return nil
}

type ConstructorType struct {
	InQuery      map[string]SchemaType
	InPath       map[string]SchemaType
	InHeader     map[string]SchemaType
	Body         map[string]SchemaType
	BodyRequired bool
}

func ConvertUrl(url string) string {
	return strings.ReplaceAll(strings.ReplaceAll(url, "{", ":"), "}", "")
}

func ToConstructorType(spec PathSpec) (res ConstructorType, err error) {
	res.InPath = make(map[string]SchemaType)
	res.InQuery = make(map[string]SchemaType)
	res.InHeader = make(map[string]SchemaType)
	res.Body = make(map[string]SchemaType)

	for _, v := range spec.Parameters {
		switch v.In {
		case "query":
			res.InQuery[v.Name] = v.Schema
		case "path":
			res.InPath[v.Name] = v.Schema
		case "header":
			res.InHeader[v.Name] = v.Schema
		}
	}

	res.BodyRequired = spec.RequestBody.Required

	for k, v := range spec.RequestBody.Content {
		res.Body[k] = v.Schema
	}

	return
}

// finalizeGoSource removes unneeded imports from the given Go source file and
// runs go fmt on it.
func finalizeGoSource(content string, wr io.Writer) error {
	// Make sure file parses and print content if it does not.
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		var buf bytes.Buffer
		scanner.PrintError(&buf, err)
		return fmt.Errorf("%s\n========\nContent:\n%s", buf.String(), content)
	}

	// Clean unused imports
	imps := astutil.Imports(fset, file)
	for _, group := range imps {
		for _, imp := range group {
			path := strings.Trim(imp.Path.Value, `"`)
			if !astutil.UsesImport(file, path) {
				if imp.Name != nil {
					astutil.DeleteNamedImport(fset, file, imp.Name.Name, path)
				} else {
					astutil.DeleteImport(fset, file, path)
				}
			}
		}
	}
	ast.SortImports(fset, file)
	if err := format.Node(wr, fset, file); err != nil {
		return err
	}

	return nil
}

//go:generate go run -tags=dev assets_generate.go -source="github.com/kepkin/gorest/pkg".Assets
