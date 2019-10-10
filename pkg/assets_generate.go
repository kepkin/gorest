// +build ignore

package main

import (
	"log"

	"github.com/kepkin/gorest/pkg"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(pkg.Assets, vfsgen.Options{
		PackageName:  "pkg",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
