package main

import (
	"log"

	"github.com/kepkin/gorest"
)

func main() {
	err := gorest.Generate("test/swagger.yaml", gorest.Options{
		PackageName: "api",
		TargetFile:  "/tmp/api_gorest.go",
	})
	if err != nil {
		log.Fatal(err)
	}
}
