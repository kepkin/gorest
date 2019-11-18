package generator

import (
	"fmt"
	"io"

	"github.com/kepkin/gorest/internal/spec/openapi3"
)

func (g Generator) makeInterfaceCheckers(wr io.Writer, _ openapi3.Spec) error {
	if len(g.customFields) > 0 {
		if _, err := fmt.Fprintln(wr, "\n// Custom types"); err != nil {
			return err
		}
	}

	//for _, def := range g.customFields {
	//	if _, err := fmt.Fprintf(wr, "var _ %s = (*%s)(nil)", def.InterfaceName, def.TypeName); err != nil {
	//		return err
	//	}
	//	if _, err := fmt.Fprintf(os.Stderr, "please implement own type `%s`\n", def.TypeName); err != nil {
	//		return err
	//	}
	//}
	return nil
}
