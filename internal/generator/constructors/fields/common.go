package fields

import (
	"strings"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var Constructors = template.FuncMap{
	"BooleanFieldConstructor": MakeBooleanFieldConstructor,
	"CustomFieldConstructor": MakeCustomFieldConstructor,
	"DateConstructor":        MakeDateFieldConstructor,
	"DateTimeConstructor":    MakeDateTimeFieldConstructor,
	"FileConstructor":        MakeFileFieldConstructor,
	"FloatConstructor":       MakeFloatFieldConstructor,
	"IntConstructor":         MakeIntFieldConstructor,
	"UnixTimeConstructor":    MakeUnixTimeFieldConstructor,
	"StringConstructor":      MakeStringFieldConstructor,
}

var BaseConstructor = template.FuncMap{
	"BaseValueFieldConstructor": MakeValueFieldConstructor,
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
