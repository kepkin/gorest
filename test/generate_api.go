// +build ignore

package main

import (
	"log"

	"github.com/kepkin/gorest"
)

func main() {
	err := gorest.Generate("swagger.yaml", gorest.Options{
		PackageName: "api",
		TargetFile:  "api/api_gorest.go",
	})
	if err != nil {
		log.Fatal(err)
	}
}
