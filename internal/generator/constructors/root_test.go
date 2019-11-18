package constructors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/barber"
	"github.com/kepkin/gorest/internal/generator/translator"
)

func TestMakeRequestConstructor(t *testing.T) {
	def := translator.TypeDef{
		Name: "IncomeRequest",
		Fields: []translator.Field{
			{Name: "Body", GoType: "IncomeRequestBody", Type: translator.StructField},
			{Name: "Cookie", GoType: "IncomeRequestCookie", Type: translator.StructField},
		},
	}

	b := &strings.Builder{}
	if !assert.NoError(t, MakeRequestConstructor(b, def)) {
		return
	}
	result := strings.NewReader("package api\n" + b.String())

	prettyResult := &strings.Builder{}
	if !assert.NoError(t, barber.PrettifySource(result, prettyResult)) {
		return
	}

	assert.Equal(t, `package api

func MakeIncomeRequest(c *gin.Context) (result IncomeRequest, errors []FieldError) {
	result.Body, errors = MakeIncomeRequestBody(c)
	if errors != nil {
		return
	}

	result.Cookie, errors = MakeIncomeRequestCookie(c)
	if errors != nil {
		return
	}
	return
}
`, prettyResult.String())
}
