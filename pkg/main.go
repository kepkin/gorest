package pkg

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/kepkin/gorest/internal/barber"
	"github.com/kepkin/gorest/internal/openapi3/generator"
	"github.com/kepkin/gorest/internal/openapi3/spec"
)

func GenerateFromFile(swaggerPath string, packageName string, wr io.Writer) error {
	content, err := ioutil.ReadFile(swaggerPath)
	if err != nil {
		return err
	}

	sp, err := spec.Read(content)
	if err != nil {
		return err
	}

	return generateFromSpec(wr, packageName, sp)
}

func generateFromSpec(wr io.Writer, packageName string, sp spec.Spec) error {
	var content bytes.Buffer

	if _, err := fmt.Fprintf(&content, generator.Predefined, packageName); err != nil {
		return err
	}

	for _, gen := range []func(io.Writer, spec.Spec) error{
		generator.MakeInterface,
		generator.MakeHandlers,
		generator.MakeRequests,
		generator.MakeComponents,
		generator.MakeRouter,
	} {
		if err := gen(&content, sp); err != nil {
			return err
		}
	}

	return barber.PrettifySource(&content, wr)
}
