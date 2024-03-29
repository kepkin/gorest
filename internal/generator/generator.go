package generator

import (
	"bytes"
	"fmt"
	"io"

	"github.com/kepkin/gorest/internal/barber"
	"github.com/kepkin/gorest/internal/spec/openapi3"
	"github.com/kepkin/gorest/internal/translator"
)

type Generator struct {
	packageName  string
	customFields customFieldsSet
	translator   translator.Translator
}

type typeName = string
type customFieldsSet map[typeName]translator.FieldPair

func NewGenerator(packageName string, translator translator.Translator) *Generator {
	return &Generator{
		packageName:  packageName,
		customFields: make(customFieldsSet),
		translator: translator,
	}
}

func (g *Generator) Generate(wr io.Writer, sp openapi3.Spec) error {
	var content bytes.Buffer

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
	} {
		if err := gen(&content, sp); err != nil {
			return err
		}
	}

	return barber.PrettifySource(&content, wr)
}
