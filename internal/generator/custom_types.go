package generator

import (
	"fmt"
	"io"
	"os"

	"github.com/kepkin/gorest/internal/spec/openapi3"
)

// TODO(a.telyshev): Test me
func (g Generator) makeInterfaceCheckers(wr io.Writer, _ openapi3.Spec) error {
	if len(g.customFields) > 0 {
		if _, err := fmt.Fprintln(wr, `
// Custom types

type FromStringSetter interface {
	SetFromString(string) error
}`); err != nil {
			return err
		}
	}

	for _, field := range g.customFields {
		if _, err := fmt.Fprintf(wr, `
var _ json.Marshaler = (*%s)(nil)
var _ json.Unmarshaler = (*%s)(nil)
var _ FromStringSetter = (*%s)(nil)
`, field.Schema.Type, field.Schema.Type, field.Schema.Type); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(os.Stderr, "please implement own type `%s`\n", field.Schema.Type); err != nil {
			return err
		}
	}
	return nil
}
