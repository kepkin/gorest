package constructors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/barber"
	"github.com/kepkin/gorest/internal/generator/translator"
	"github.com/kepkin/gorest/internal/spec/openapi3"
)

func TestMakeHeaderParamsConstructor(t *testing.T) {
	defaultXConsumerID := "1"

	def := translator.TypeDef{
		Name: "IncomeRequestHeaders",
		Fields: []translator.Field{
			{Name: "XAccessToken", GoType: "string", Parameter: "X-Access-Token", Type: translator.StringField},
			{Name: "XConsumerID", GoType: "int64", Parameter: "X-Consumer-ID", Type: translator.IntegerField, Schema: openapi3.SchemaType{Default: &defaultXConsumerID}},
		},
	}

	b := &strings.Builder{}
	if !assert.NoError(t, MakeHeaderParamsConstructor(b, def)) {
		return
	}
	result := strings.NewReader("package api\n" + b.String())

	prettyResult := &strings.Builder{}
	if !assert.NoError(t, barber.PrettifySource(result, prettyResult)) {
		return
	}

	assert.Equal(t, `package api

func MakeIncomeRequestHeaders(c *gin.Context) (result IncomeRequestHeaders, errors []FieldError) {
	var err error

	xAccessTokenStr := c.Request.Header.Get("X-Access-Token")
	result.XAccessToken = xAccessTokenStr

	xConsumerIDStr := c.Request.Header.Get("X-Consumer-ID")
	if xConsumerIDStr != "" {
		xConsumerIDStr = "1"
	}
	result.XConsumerID, err = strconv.ParseInt(xConsumerIDStr, 10, 0)
	if err != nil {
		errors = append(errors, NewFieldError(InHeader, "X-Consumer-ID", "can't parse as integer", err))
	}
	return
}
`, prettyResult.String())
}
