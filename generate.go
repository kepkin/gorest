//Tool for generating REST handlers from OpenApi (swagger) specification.
package gorest

import (
	"github.com/kepkin/gorest/pkg"
	"os"
)

// Options for gorest code generation.
type Options struct {
	// PackageName is the name of the package in the generated code.
	// If left empty, it defaults to "api".
	PackageName string

	// Name for generated file
	// If left empty, it defaults to stdout.
	TargetFile string
}

// Generates go file from swagger specification
func Generate(swaggerFile string, options Options) error {
	var err error
	f := os.Stdout
	pkgName := "api"

	if options.TargetFile != "" {
		f, err = os.Create(options.TargetFile)
		if err != nil {
			return err
		}
	}

	if options.PackageName == "" {
		pkgName = options.PackageName
	}

	err = pkg.GenerateFromFile(swaggerFile, pkgName, f)
	return err
}
