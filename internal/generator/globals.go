package generator

import (
	"io"
	"sort"

	"github.com/kepkin/gorest/internal/spec/openapi3"
)

import (
	"fmt"
)

func (g *Generator) makeImports(wr io.Writer, imports []string) error {

	imports = append(imports, g.translator.ImportsRequired()...)

	sort.Strings(imports)

	wr.Write([]byte("import (\n"))
	for _, importPath := range imports {
		wr.Write([]byte(fmt.Sprintf("\t \"%v\"\n", importPath)))
	}
	wr.Write([]byte(")\n\n"))

	return nil
}

func (g *Generator) makeGlobals(wr io.Writer, sp openapi3.Spec) error {

	str, err := g.translator.BuildGlobalCode()
	if err != nil {
		return err
	}
	wr.Write([]byte(str))

	return ginGlobalTemplate.Execute(wr, nil)
}
