package pkg

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/kepkin/gorest/internal/openapi3/barber"
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

	return generateFromSpec(sp, packageName, wr)
}

func generateFromSpec(sp spec.Spec, packageName string, wr io.Writer) error {
	var content bytes.Buffer

	if _, err := fmt.Fprintf(&content, generator.Predefined, packageName); err != nil {
		return err
	}

	if err := generator.MakeInterface(&content, sp); err != nil {
		return err
	}

	if err := generator.MakeHandlers(&content, sp); err != nil {
		return err
	}

	if err := generator.MakeRouter(&content, sp); err != nil {
		return err
	}

	if err := barber.PrettifySource(&content, wr); err != nil {
		return err
	}
	return nil
}
