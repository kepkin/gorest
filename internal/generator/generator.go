package generator

import (
	"bytes"
	"fmt"
	"io"

	"github.com/kepkin/gorest/internal/barber"
	"github.com/kepkin/gorest/internal/translator"
	"github.com/kepkin/gorest/internal/spec/openapi3"
)

type Generator struct {
	packageName  string
	customFields customFieldsSet
	translator   translator.Translator
}

type typeName = string
type customFieldsSet map[typeName]translator.FieldPair

func NewGenerator(packageName string) *Generator {
	return &Generator{
		packageName:  packageName,
		customFields: make(customFieldsSet),
	}
}

func (g *Generator) Generate(wr io.Writer, sp openapi3.Spec) error {
	var content bytes.Buffer

	g.translator = MakeTranslator()

	if _, err := fmt.Fprintf(&content, PredefinedHeader, g.packageName); err != nil {
		return err
	}

	g.makeImports(&content, []string{
		"encoding/json",
		//"encoding/xml",
		"fmt",
		//"mime/multipart",
		"net/http",
		"strconv",
		"strings",
		"time",

		"github.com/gin-gonic/gin",
	})

	//noinspection GoPrintFunctions
	if _, err := fmt.Fprintf(&content, Predefined); err != nil {
		return err
	}

	for _, gen := range []func(io.Writer, openapi3.Spec) error{
		g.makeGlobals,
		g.makeInterface,
		g.makeHandlers,
		g.makeRouter,
		g.makeComponents,
		g.makeInterfaceCheckers,
	} {
		if err := gen(&content, sp); err != nil {
			return err
		}
	}

	return barber.PrettifySource(&content, wr)
}
