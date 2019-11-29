package fields

import (
	"strings"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var Constructors = template.FuncMap{
	"CustomFieldConstructor": MakeCustomFieldConstructor,
	"FloatConstructor":       MakeFloatFieldConstructor,
	"IntConstructor":         MakeIntFieldConstructor,
	"DateTimeConstructor":    MakeDateTimeFieldConstructor,
	"DateConstructor":        MakeDateFieldConstructor,
	"UnixTimeConstructor":    MakeUnixTimeFieldConstructor,
}

func executeFieldConstructorTemplate(t *template.Template, f translator.Field, place string) (string, error) {
	wr := &strings.Builder{}
	err := t.Execute(wr, struct {
		translator.Field
		Place string
	}{
		Field: f,
		Place: place,
	})
	return wr.String(), err
}
