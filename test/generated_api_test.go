package test

import (
	"testing"

	"github.com/kepkin/gorest/test/api"
)

//go:generate go run generate_api.go

var _ api.TestApi = (*api.TestApiImpl)(nil)

func TestCreatePoliciesInBatch(t *testing.T) {

}

func TestListPolicies(t *testing.T) {

}
