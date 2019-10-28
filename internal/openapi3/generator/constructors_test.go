package generator

import (
	"strings"
	"testing"

	"github.com/kepkin/gorest/internal/openapi3/barber"
	"github.com/kepkin/gorest/internal/openapi3/translator"
	"github.com/stretchr/testify/assert"
)

func TestMakeQueryParamsConstructor(t *testing.T) {
	def := translator.TypeDef{
		Name: "QueryParams",
		Fields: []interface{}{
			translator.StringField{
				Field: translator.Field{Name: "Filter", Parameter: "filter", Place: "Query"},
			},
		},
	}

	b := new(strings.Builder)
	if !assert.NoError(t, MakeConstructor(def, b)) {
		return
	}

	result, err := barber.PrettifySource("package api\n" + b.String())
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, `package api

func MakeQueryParams(c *gin.Context) (result QueryParams, errors []FieldError) {
	result.Filter = ExtractParameter(c, "filter", Query)
	return
}
`, result)
}
