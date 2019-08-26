package pkg

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v2"
	_ "gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"strings"
	"text/template"
)

type InfoType struct {
	Title   string
	Version string
}

type SchemaType struct {
	Type   string
	Properties map[string]SchemaType
	Format string
	Items  *SchemaType
	Ref    string `yaml:"$ref"`
}

type ContentType struct {
	Schema SchemaType
}

type RequestBodyType struct {
	Description string
	Required    bool
	Content     map[string]ContentType
}

type ParametersType struct {
	Name     string
	In       string
	Required bool
	Schema   SchemaType
}

type PathSpec struct {
	Summary     string
	Description string
	OperationID string `yaml:"operationId"`
	Parameters  []ParametersType
	RequestBody RequestBodyType `yaml:"requestBody"`
}

type PathType struct {
	Post    *PathSpec
	Get     *PathSpec
	Options *PathSpec
	Put     *PathSpec
	Delete  *PathSpec
}

type PathMap map[string]PathType

type ComponentsType struct {
	Schemas map[string]SchemaType
}

type Spec struct {
	Openapi string
	Info    InfoType
	Paths   PathMap
	Components ComponentsType
}

func ReadSpec(in []byte) (res Spec, err error) {
	err = yaml.Unmarshal(in, &res)
	return
}

func BuildTpls() (res *template.Template, err error) {
	a := sprig.GenericFuncMap()
	res = template.New("").Funcs(a)
	res = res.Funcs(template.FuncMap{
		"MakeIdentifier": MakeIdentifier,
		"ConvertType":    ConvertType,
		"GetNameFromRef": GetNameFromRef,
		"ToConstructorType": ToConstructorType,
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
		return res, fmt.Errorf("gorest: No template files!")
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

func GenerateFromFile(swaggerPath string, wr io.Writer) error {
	content, err := ioutil.ReadFile(swaggerPath)
	if err != nil {
		return err
	}
	spec, err := ReadSpec(content)
	if err != nil {
		return err
	}

	return GenerateFromSpec(spec, wr)
}

func GenerateFromSpec(spec Spec, wr io.Writer) error {
	t, err := BuildTpls()
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(wr, "main.tmpl", spec)
	return err
}

func MakeIdentifier(s string) (string, error) {
	s = strings.ReplaceAll(s, " ", "_")
	s = strcase.ToCamel(s)
	return s, nil


}

func GetNameFromRef(s string) string {
	return s[len("#/components/schemas/"):]
}

func ConvertType(spec SchemaType) string {
	if spec.Ref != "" {
		return GetNameFromRef(spec.Ref)
	}

	var type_ string

	switch spec.Type {
	case "integer":
		switch spec.Format {
		case "int64":
			type_ = "int64"
		default:
			type_ = "int"
		}

	case "array":
		type_ = "[]" + GetNameFromRef(spec.Items.Ref)

	default:
		type_ = spec.Type
	}

	return type_
}

type ConstructorType struct {
	InQuery      map[string]SchemaType
	InPath       map[string]SchemaType
	InHeader     map[string]SchemaType
	Body         map[string]SchemaType
	BodyRequired bool
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

//go:generate vfsgendev -source="github.com/kepkin/gorest/pkg".Assets
