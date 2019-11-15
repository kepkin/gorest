package generator

import (
	"strings"
	"testing"

	"github.com/kepkin/gorest/internal/barber"
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
	if !assert.NoError(t, MakeConstructor(b, def)) {
		return
	}
	result := strings.NewReader("package api\n" + b.String())

	prettyResult := new(strings.Builder)
	if !assert.NoError(t, barber.PrettifySource(result, prettyResult)) {
		return
	}

	assert.Equal(t, `package api

func MakeQueryParams(ctx *gin.Context) (result QueryParams, errors []FieldError) {
	result.Filter = ExtractParameter(ctx, "filter", Query)
	return
}
`, prettyResult.String())
}
