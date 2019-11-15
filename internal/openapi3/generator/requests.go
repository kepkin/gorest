package generator

import (
	"fmt"
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/openapi3/spec"
	"github.com/kepkin/gorest/internal/openapi3/translator"
)

func MakeRequests(wr io.Writer, sp spec.Spec) error {
	for _, path := range sp.Paths {
		for _, method := range path.Methods() {
			if method != nil {
				if err := MakeRequest(wr, method); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func MakeRequest(wr io.Writer, path *spec.PathSpec) error {
	request := spec.NewObjectSchema(path.OperationID+"Request", 0, "")

	queryParams := spec.NewObjectSchema("", 1, "Query")
	pathParams := spec.NewObjectSchema("", 1, "Path")
	cookieParams := spec.NewObjectSchema("", 1, "Cookie")
	headerParams := spec.NewObjectSchema("", 1, "Header")

	for _, param := range path.Parameters {
		switch param.In {
		case "query":
			queryParams.Properties[param.Name] = param.Schema
		case "path":
			pathParams.Properties[param.Name] = param.Schema
		case "cookie":
			cookieParams.Properties[param.Name] = param.Schema
		case "header":
			headerParams.Properties[param.Name] = param.Schema
		default:
			return fmt.Errorf("unknown param place: %s", param.In)
		}
	}

	if len(queryParams.Properties) > 0 {
		request.Properties["Query"] = &queryParams
	}
	if len(pathParams.Properties) > 0 {
		request.Properties["Path"] = &pathParams
	}
	if len(cookieParams.Properties) > 0 {
		request.Properties["Cookie"] = &cookieParams
	}
	if len(headerParams.Properties) > 0 {
		request.Properties["Header"] = &headerParams
	}

	if path.RequestBody != nil {
		for mimeType, content := range path.RequestBody.Content {
			switch mimeType {
			case "application/json":
				request.Properties["JSONBody"] = content.Schema
			case "application/xml":
				request.Properties["XMLBody"] = content.Schema
			default:
				return fmt.Errorf("unsupported content type: %s", mimeType)
			}
		}
	}

	// Determine definitions
	defs, err := translator.ProcessRootSchema(request)
	if err != nil {
		return err
	}

	// Make structs
	for _, d := range defs {
		if typeDef, ok := d.(translator.TypeDef); ok {
			err = MakeStruct(wr, typeDef)
			if err != nil {
				return err
			}

			err = MakeValidateFunc(wr, typeDef)
			if err != nil {
				return err
			}
		}
	}

	// Make constructors
	for _, d := range defs {
		switch dd := d.(type) {
		case translator.TypeDef:
			if err := MakeConstructor(wr, dd); err != nil {
				return err
			}

		case translator.InterfaceCheckerDef:
			if _, err := fmt.Fprintf(wr, "var _ %s = (*%s)(nil)", dd.InterfaceName, dd.TypeName); err != nil {
				return err
			}
		}
	}

	// Make validators
	// TODO(a.telyshev)

	return nil
}

var structTemplate = template.Must(template.New("struct").Parse(`
type {{ .Name }} struct {
	{{- range $, $field := .Fields }}
	{{ $field.Name }} {{ $field.Type -}}
	{{ end }}
}
`))

func MakeStruct(wr io.Writer, def translator.TypeDef) error {
	return structTemplate.Execute(wr, def)
}
