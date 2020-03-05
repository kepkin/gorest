package generator

import (
	"io"
	"sort"

	"github.com/kepkin/gorest/internal/spec/openapi3"
	"github.com/kepkin/gorest/internal/translator"
)

import (
	"fmt"
)

func (g *Generator) makeImports(wr io.Writer, imports []string) error {

	//TODO: triple duplication: here and in MakeTranslator
	for _, fieldImpl := range []translator.FieldI{
		&translator.BooleanFieldImpl{ImplIdentifier: "a"},
		&translator.StringFieldImpl{},
		&translator.IntegerFieldImpl{},
		&translator.ObjectFieldImpl{},
		&translator.DecimalFieldImpl{},
	} {
		imports = append(imports, fieldImpl.ImportsRequired()...)
	}

	sort.Strings(imports)

	wr.Write([]byte("import (\n"))
	for _, importPath := range imports {
		wr.Write([]byte(fmt.Sprintf("\t \"%v\"\n", importPath)))
	}
	wr.Write([]byte(")\n\n"))

	return nil
}

func (g *Generator) makeGlobals(wr io.Writer, sp openapi3.Spec) error {

	//TODO: triple duplication: here and in MakeTranslator
	for _, v := range []translator.FieldI {
		&translator.BooleanFieldImpl{ImplIdentifier: "a"},
		&translator.StringFieldImpl{},
		&translator.IntegerFieldImpl{},
		&translator.ObjectFieldImpl{},
		&translator.DecimalFieldImpl{},
	} {
		str, err := v.BuildGlobalCode()
		if err != nil {
			return err
		}
		wr.Write([]byte(str))
	}

	return ginGlobalTemplate.Execute(wr, nil)
}
