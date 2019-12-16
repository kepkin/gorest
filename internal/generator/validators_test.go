package generator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/barber"
	"github.com/kepkin/gorest/internal/generator/translator"
	"github.com/kepkin/gorest/internal/spec/openapi3"
)

func TestMakeValidateFunc(t *testing.T) {
	t.Run("Call Validate() for struct field", func(t *testing.T) {
		def := translator.TypeDef{
			Name:   "RequestQuery",
			GoType: "struct",
			Fields: []translator.Field{
				{Name: "User", Parameter: "user", GoType: "struct", Type: translator.StructField},
			},
		}

		b := &strings.Builder{}
		if !assert.NoError(t, NewGenerator("api").makeValidateFunc(b, def)) {
			return
		}
		result := strings.NewReader("package api\n" + b.String())

		prettyResult := &strings.Builder{}
		if !assert.NoError(t, barber.PrettifySource(result, prettyResult)) {
			return
		}

		assert.Equal(t, `package api

func (t RequestQuery) Validate() (errors []FieldError) {
	// User field validators
	errors = t.User.Validate()
	if errors != nil {
		return
	}
	return
}
`, prettyResult.String())
	})
}

func TestMakeEnumValidator(t *testing.T) {
	def := translator.TypeDef{
		Name:   "RequestQuery",
		GoType: "struct",
		Fields: []translator.Field{
			{
				Name: "Async", Parameter: "async", GoType: "string", Type: translator.StringField,
				Schema: openapi3.SchemaType{Enum: []string{"true", "false"}},
			},
			{
				Name: "Quantile", Parameter: "quantile", GoType: "float64", Type: translator.FloatField,
				Schema: openapi3.SchemaType{Enum: []string{"0.25", "0.5", "0.75", "1"}},
			},
			{
				Name: "Percentage", Parameter: "percentage", GoType: "int64", Type: translator.IntegerField,
				Schema: openapi3.SchemaType{Enum: []string{"0", "50", "100"}},
			},
		},
	}

	b := &strings.Builder{}
	if !assert.NoError(t, NewGenerator("api").makeValidateFunc(b, def)) {
		return
	}
	result := strings.NewReader("package api\n" + b.String())

	prettyResult := &strings.Builder{}
	if !assert.NoError(t, barber.PrettifySource(result, prettyResult)) {
		return
	}

	assert.Equal(t, `package api

func (t RequestQuery) Validate() (errors []FieldError) {
	// Async field validators
	var asyncInEnum bool
	for _, elem := range [...]string{
		"true",
		"false",
	} {
		if elem == t.Async {
			asyncInEnum = true
			break
		}
	}
	if !asyncInEnum {
		errors = append(errors, NewFieldError(UndefinedPlace, "async", "allowed values: [true false]", nil))
	}

	// Quantile field validators
	var quantileInEnum bool
	for _, elem := range [...]float64{
		0.25,
		0.5,
		0.75,
		1,
	} {
		if elem == t.Quantile {
			quantileInEnum = true
			break
		}
	}
	if !quantileInEnum {
		errors = append(errors, NewFieldError(UndefinedPlace, "quantile", "allowed values: [0.25 0.5 0.75 1]", nil))
	}

	// Percentage field validators
	var percentageInEnum bool
	for _, elem := range [...]int64{
		0,
		50,
		100,
	} {
		if elem == t.Percentage {
			percentageInEnum = true
			break
		}
	}
	if !percentageInEnum {
		errors = append(errors, NewFieldError(UndefinedPlace, "percentage", "allowed values: [0 50 100]", nil))
	}

	return
}
`, prettyResult.String())
}
