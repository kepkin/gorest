package generator

import (
	"io"
	"sort"

	"github.com/kepkin/gorest/internal/generator/translator"
	"github.com/kepkin/gorest/internal/spec/openapi3"
)

func (g *Generator) makeComponents(wr io.Writer, sp openapi3.Spec) error {

	sortedComponents := make([]string, 0, len(sp.Components.Schemas))
	for k := range sp.Components.Schemas {
		sortedComponents = append(sortedComponents, k)
	}

	sort.Strings(sortedComponents)

	for _, name := range sortedComponents {
		schema := sp.Components.Schemas[name]
		schema.Name = name
		defs, err := translator.ProcessRootSchema(*schema)
		if err != nil {
			return err
		}
		for _, d := range defs {
			if err := g.makeStruct(wr, d, true); err != nil {
				return err
			}

			if err := g.makeValidateFunc(wr, d); err != nil {
				return err
			}
		}
	}

	return nil
}
