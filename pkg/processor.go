package pkg

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

type specMeta struct {
	CustomTypes  []string
	SpecialTypes []schemaType
}

func processSpec(spec *spec) (meta specMeta, err error) {
	customTypes := make(map[string]struct{})
	specialTypes := make(map[string]*schemaType)

	schemaHandler := func(path []string, schema *schemaType) (err error) {
		if err = determineTypes(path, schema); err != nil {
			return
		}
		collectTypes(schema, customTypes, specialTypes)
		return
	}

	for _, path := range spec.Paths {
		if err = processPath(path.Delete); err != nil {
			return
		}
		if err = processPath(path.Get); err != nil {
			return
		}
		if err = processPath(path.Options); err != nil {
			return
		}
		if err = processPath(path.Patch); err != nil {
			return
		}
		if err = processPath(path.Post); err != nil {
			return
		}
		if err = processPath(path.Put); err != nil {
			return
		}
	}

	for name, schema := range spec.Components.Schemas {
		if err = walkObjectProperties([]string{name}, schema.Properties, schemaHandler); err != nil {
			return
		}
	}

	meta.CustomTypes = make([]string, 0, len(customTypes))
	for t := range customTypes {
		meta.CustomTypes = append(meta.CustomTypes, t)
	}

	meta.SpecialTypes = make([]schemaType, 0, len(specialTypes))
	for _, schema := range specialTypes {
		meta.SpecialTypes = append(meta.SpecialTypes, *schema)
	}
	return
}

func processPath(p *pathSpec) error {
	if p == nil {
		return nil
	}

	for _, param := range p.Parameters {
		if err := determineTypes(nil, param.Schema); err != nil {
			return err
		}
	}

	if p.RequestBody != nil {
		for _, content := range p.RequestBody.Content {
			// TODO(a.telyshev): Not only `ref` support
			if err := determineTypes(nil, content.Schema); err != nil {
				return err
			}
		}
	}
	return nil
}

func determineTypes(path []string, schema *schemaType) error {
	if err := determineType(schema); err != nil {
		return err
	}
	if schema.Type == objectType && len(path) > 1 {
		typeBuilder := new(strings.Builder)
		for _, name := range path {
			typeBuilder.WriteString(MakeTitledIdentifier(name))
		}
		t := typeBuilder.String()
		schema.GoType = t
		schema.HasSpecialType = true
	}
	return nil
}

func collectTypes(schema *schemaType, customTypes map[string]struct{}, specialTypes map[string]*schemaType) {
	if schema.HasCustomType {
		customTypes[schema.GoType] = struct{}{}
	}
	if schema.HasSpecialType {
		specialTypes[schema.GoType] = schema
	}
}

func determineType(spec *schemaType) error {
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

type handlerFunc func(path []string, schema *schemaType) error

func walkObjectProperties(path []string, props propertiesType, handler handlerFunc) error {
	if len(props) == 0 {
		return nil
	}
	for propName, propSchema := range props {
		p := append(path[:0:0], path...) // Just slice copying
		p = append(p, propName)

		if err := handler(p, propSchema); err != nil {
			return err
		}
		if propSchema.Type == objectType {
			if err := walkObjectProperties(p, propSchema.Properties, handler); err != nil {
				return err
			}
		}
	}
	return nil
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

type ConstructorType struct {
	InQuery      map[string]*schemaType
	InPath       map[string]*schemaType
	InHeader     map[string]*schemaType
	Body         map[string]*schemaType
	BodyRequired bool
}

func ConvertUrl(url string) string {
	return strings.ReplaceAll(strings.ReplaceAll(url, "{", ":"), "}", "")
}

func ToConstructorType(spec pathSpec) (res ConstructorType, err error) {
	res.InPath = make(map[string]*schemaType)
	res.InQuery = make(map[string]*schemaType)
	res.InHeader = make(map[string]*schemaType)

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

	if spec.RequestBody != nil {
		res.Body = make(map[string]*schemaType)
		res.BodyRequired = spec.RequestBody.Required

		for k, v := range spec.RequestBody.Content {
			res.Body[k] = v.Schema
		}
	}
	return
}
