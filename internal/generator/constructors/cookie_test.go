package constructors

import (
	"github.com/kepkin/gorest/internal/spec/openapi3"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/barber"
	"github.com/kepkin/gorest/internal/generator/translator"
)

func TestMakeCookieParamsConstructor(t *testing.T) {
	def := translator.TypeDef{
		Name: "IncomeRequestCookie",
		Fields: []translator.Field{
			{Name: "SessionID", GoType: "string", Parameter: "sessionID", Type: translator.StringField, Schema: openapi3.SchemaType{Default: "id session"}},
			{Name: "MaxAge", GoType: "int64", Parameter: "Max-Age", Type: translator.IntegerField, Schema: openapi3.SchemaType{Default: "22"}},
			{Name: "Domain", GoType: "string", Parameter: "Domain", Type: translator.StringField},
		},
	}

	b := &strings.Builder{}
	if !assert.NoError(t, MakeCookieParamsConstructor(b, def)) {
		return
	}
	result := strings.NewReader("package api\n" + b.String())

	prettyResult := &strings.Builder{}
	if !assert.NoError(t, barber.PrettifySource(result, prettyResult)) {
		return
	}

	assert.Equal(t, `package api

func MakeIncomeRequestCookie(c *gin.Context) (result IncomeRequestCookie, errors []FieldError) {
	var err error

	getCookie := func(param string) (string, bool) {
		cookie, err := c.Request.Cookie(param)
		if err == http.ErrNoCookie {
			return "", false
		}
		return cookie.Value, true
	}

	result.SessionID, ok = getCookie("sessionID")
	if !ok {
		result.SessionID = "id session"
	}

	maxAgeStr, ok := getCookie("Max-Age")
	if !ok {
		maxAgeStr = "22"
	}

	result.MaxAge, err = strconv.ParseInt(maxAgeStr, 10, 0)
	if err != nil {
		errors = append(errors, NewFieldError(InCookie, "Max-Age", "can't parse as integer", err))
	}

	result.Domain, _ = getCookie("Domain")
	return
}
`, prettyResult.String())
}
