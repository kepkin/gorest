package constructors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/barber"
	"github.com/kepkin/gorest/internal/generator/translator"
)

func TestMakeHeaderParamsConstructor(t *testing.T) {
	def := translator.TypeDef{
		Name: "IncomeRequestHeaders",
		Fields: []translator.Field{
			{Name: "XAccessToken", GoType: "string", Parameter: "X-Access-Token", Type: translator.StringField},
			{Name: "XConsumerID", GoType: "int64", Parameter: "X-Consumer-ID", Type: translator.IntegerField},
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

	result.XAccessToken = c.Request.Header.Get("X-Access-Token")

	xConsumerIDStr := c.Request.Header.Get("X-Consumer-ID")
	result.XConsumerID, err = strconv.ParseInt(xConsumerIDStr, 10, 0)
	if err != nil {
		errors = append(errors, NewFieldError(InHeader, "X-Consumer-ID", "can't parse as integer", err))
	}
	return
}
`, prettyResult.String())
}
