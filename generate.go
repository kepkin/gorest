// Tool for generating REST handlers from OpenAPI (Swagger) specification.
package gorest

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kepkin/gorest/pkg"
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
func Generate(swaggerFile string, options Options) (err error) {
	pkgName := "api"
	if options.PackageName != "" {
		pkgName = options.PackageName
	}

	// Generate to stdout
	if options.TargetFile == "" {
		return pkg.GenerateFromFile(swaggerFile, pkgName, os.Stdout)
	}

	// Generate to tmp file. If no errors presents then will move tmp file to `options.TargetFile`
	genFile, err := ioutil.TempFile("", "*-gorest.go")
	if err != nil {
		return fmt.Errorf("error while tmp file creating: %v", err)
	}

	err = pkg.GenerateFromFile(swaggerFile, pkgName, genFile)
	if err != nil {
		err = fmt.Errorf("api generating error: %v", err)

		if removeErr := removeFile(genFile); removeErr != nil {
			return fmt.Errorf("%s, error while tmp file removing: %v", err, removeErr)
		}
		return err
	}
	return os.Rename(genFile.Name(), options.TargetFile)
}

func removeFile(f *os.File) (err error) {
	err = f.Close()
	if err != nil {
		return
	}

	err = os.Remove(f.Name())
	if err != nil {
		return
	}
	return
}
