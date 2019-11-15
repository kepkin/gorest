package gorest_test

import (
	"log"
	"testing"

	"github.com/kepkin/gorest"
)

func TestExampleGenerate(t *testing.T) {
	err := gorest.Generate("test/swagger.yaml", gorest.Options{
		PackageName: "api",
		TargetFile:  "__gingen.go",
	})
	if err != nil {
		log.Fatal(err)
	}
}
