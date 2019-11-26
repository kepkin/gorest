package generator

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

// TODO(a.telyshev): Write me
var validateTemplate = template.Must(template.New("validate").Parse(`
func (t {{ .Name }}) Validate() []FieldError {
	return nil
}
`))

func makeValidateFunc(wr io.Writer, def translator.TypeDef) error {
	return validateTemplate.Execute(wr, def)
}
