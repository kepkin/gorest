// +build ignore

package main

import (
	"log"

	"github.com/kepkin/gorest"
)

func main() {
	err := gorest.Generate("swagger_wallet.yaml", gorest.Options{
		PackageName: "api",
		TargetFile:  "api/api_gorest_w1.go",
	})
	if err != nil {
		log.Fatal(err)
	}
}
