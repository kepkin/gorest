package constructors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/barber"
	"github.com/kepkin/gorest/internal/generator/translator"
)

func TestMakeQueryParamsConstructor(t *testing.T) {
	def := translator.TypeDef{
		Name: "IncomeRequestQuery",
		Fields: []translator.Field{
			{Name: "ID", GoType: "string", Parameter: "id", Type: translator.StringField},
			{Name: "Size", GoType: "int64", Parameter: "size", Type: translator.IntegerField},
			{Name: "Sum", GoType: "float64", Parameter: "sum", Type: translator.FloatField},
			{Name: "User", GoType: "User", Parameter: "user", Type: translator.CustomField},
		},
	}

	b := &strings.Builder{}
	if !assert.NoError(t, MakeQueryParamsConstructor(b, def)) {
		return
	}
	result := strings.NewReader("package api\n" + b.String())

	prettyResult := &strings.Builder{}
	if !assert.NoError(t, barber.PrettifySource(result, prettyResult)) {
		return
	}

	assert.Equal(t, `package api

func MakeIncomeRequestQuery(c *gin.Context) (result IncomeRequestQuery, errors []FieldError) {
	var err error

	result.ID, _ = c.GetQuery("id")

	sizeStr, _ := c.GetQuery("size")
	result.Size, err = strconv.ParseInt(sizeStr, 10, 0)
	if err != nil {
		errors = append(errors, NewFieldError(InQuery, "size", "can't parse as integer", err))
	}

	sumStr, _ := c.GetQuery("sum")
	result.Sum, err = strconv.ParseFloat(sumStr, 10, 0)
	if err != nil {
		errors = append(errors, NewFieldError(InQuery, "sum", "can't parse as float", err))
	}

	userStr, _ := c.GetQuery("user")
	result.User = User{}
	if err = result.User.SetFromString(userStr); err != nil {
		errors = append(errors, NewFieldError(InQuery, "user", fmt.Sprintf("can't create from string '%s'", userStr), err))
	}
	return
}
`, prettyResult.String())
}
