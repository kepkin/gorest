package generator

import (
	"fmt"
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/openapi3/spec"
	"github.com/kepkin/gorest/internal/openapi3/translator"
)

func MakeComponents(wr io.Writer, sp spec.Spec) error {

	for name, schema := range sp.Components.Schemas {
		schema.Name = name
		defs, _ := translator.ProcessRootSchema(*schema)
		for _, d := range defs {
			switch dd := d.(type) {
			case translator.TypeDef:
				if err := MakeStruct(wr, dd); err != nil {
					return err
				}

				if err := MakeComponentConstructor(wr, dd); err != nil {
					return err
				}

			case translator.InterfaceCheckerDef:
				if _, err := fmt.Fprintf(wr, "var _ %s = (*%s)(nil)", dd.InterfaceName, dd.TypeName); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

var componentConstructorTemplate = template.Must(template.New("componentConstructor").Parse(`
func Make{{ .Name }}(c *gin.Context) (result {{ .Name }}, errors []FieldError) {
	if err := json.NewDecoder(c.Request.Body).Decode(&result); err != nil {
		errors = append(errors, FieldError{
			Field:   "{{ .Name }}",
			Message: "can't parse JSON",
			Reason:  err.Error(),
		})
	}
	return 
}
`))

func MakeComponentConstructor(wr io.Writer, def translator.TypeDef) error {
	return componentConstructorTemplate.Execute(wr, def)
}
