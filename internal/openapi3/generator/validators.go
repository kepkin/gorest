package generator

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/openapi3/translator"
)

var validateTemplate = template.Must(template.New("validate").Parse(`
func (t {{ .Name }}) Validate() []FieldError {
	return nil
}
`))

func MakeValidateFunc(wr io.Writer, def translator.TypeDef) error {
	return validateTemplate.Execute(wr, def)
}
