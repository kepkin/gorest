package generator

import (
	"io"

	"github.com/kepkin/gorest/internal/generator/translator"
	"github.com/kepkin/gorest/internal/spec/openapi3"
)

func (g *Generator) makeComponents(wr io.Writer, sp openapi3.Spec) error {

	for name, schema := range sp.Components.Schemas {
		schema.Name = name
		defs, _ := translator.ProcessRootSchema(*schema)
		for _, d := range defs {
			if err := g.makeStruct(wr, d, true); err != nil {
				return err
			}
		}
	}

	return nil
}
