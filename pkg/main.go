package pkg

import (
	"io"
	"io/ioutil"

	"github.com/kepkin/gorest/internal/generator"
	"github.com/kepkin/gorest/internal/spec/openapi3"
)

func GenerateFromFile(swaggerPath string, packageName string, wr io.Writer) error {
	content, err := ioutil.ReadFile(swaggerPath)
	if err != nil {
		return err
	}

	sp, err := openapi3.ReadSpec(content)
	if err != nil {
		return err
	}

	g := generator.NewGenerator(packageName)
	return g.Generate(wr, sp)
}
