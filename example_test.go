package gorest_test

import (
	"github.com/kepkin/gorest"
	"log"
)

func ExampleGenerate() {
	err := gorest.Generate("../assets/swagger.yaml", gorest.Options{
		PackageName: "api",
		TargetFile:  "gingen.go",
	})
	if err != nil {
		log.Fatal(err)
	}
}
