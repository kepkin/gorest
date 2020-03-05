package test

import (
	"log"
	"testing"

	"github.com/kepkin/gorest"
)

//go:generate go run generate_api.go

func TestForDebug(t *testing.T) {
	//t.Skip()
	err := gorest.Generate("swagger.yaml", gorest.Options{
		PackageName: "api",
		TargetFile:  "api/api_gorest.go",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func TestQueryParams(t *testing.T) {
	//t.Skip()
	err := gorest.Generate("swaggers/query_params.yaml", gorest.Options{
		PackageName: "apis/",
		TargetFile:  "apis/query_params/gorest.go",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func TestJSONBodyParams(t *testing.T) {
	//t.Skip()
	err := gorest.Generate("swaggers/json_body_params.yaml", gorest.Options{
		PackageName: "apis/",
		TargetFile:  "apis/json_body_params/gorest.go",
	})
	if err != nil {
		log.Fatal(err)
	}
}
