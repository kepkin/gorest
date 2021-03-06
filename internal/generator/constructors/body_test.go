package constructors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/barber"
	"github.com/kepkin/gorest/internal/generator/translator"
)

func TestMakeBodyConstructor(t *testing.T) {
	t.Run("Only JSON body", func(t *testing.T) {
		def := translator.TypeDef{
			Name: "IncomeRequestBody",
			Fields: []translator.Field{
				{Name: "JSON", GoType: "IncomeRequestBodyJSON"},
				{Name: "Type", GoType: "ContentType"},
			},
		}

		assertGeneratedCode(t, def, `package api

func MakeIncomeRequestBody(c *gin.Context) (result IncomeRequestBody, errors []FieldError) {
	contentType := c.Request.Header.Get("Content-Type")

	if contentType == "" {
		errors = append(errors, NewFieldError(InBody, "-", "no content type specified", nil))
		return
	}
	contentType = strings.Split(contentType, ";")[0]

	switch contentType {
	case "application/json":
		result.Type = AppJSON
		if err := json.NewDecoder(c.Request.Body).Decode(&result.JSON); err != nil {
			errors = append(errors, NewFieldError(InBody, "requestBody", "can't decode body from JSON", err))
		}

	default:
		errors = append(errors, NewFieldError(InBody, "-", "unknown content type", nil))
	}
	return
}
`)
	})

	t.Run("Multivariant body", func(t *testing.T) {
		def := translator.TypeDef{
			Name: "IncomeRequestBody",
			Fields: []translator.Field{
				{Name: "JSON", GoType: "IncomeRequestBodyJSON"},
				{Name: "XML", GoType: "IncomeRequestBodyXML"},
				{Name: "Type", GoType: "ContentType"},
			},
		}

		assertGeneratedCode(t, def, `package api

func MakeIncomeRequestBody(c *gin.Context) (result IncomeRequestBody, errors []FieldError) {
	contentType := c.Request.Header.Get("Content-Type")

	if contentType == "" {
		errors = append(errors, NewFieldError(InBody, "-", "no content type specified", nil))
		return
	}
	contentType = strings.Split(contentType, ";")[0]

	switch contentType {
	case "application/json":
		result.Type = AppJSON
		if err := json.NewDecoder(c.Request.Body).Decode(&result.JSON); err != nil {
			errors = append(errors, NewFieldError(InBody, "requestBody", "can't decode body from JSON", err))
		}

	case "application/xml":
		result.Type = AppXML
		if err := xml.NewDecoder(c.Request.Body).Decode(&result.XML); err != nil {
			errors = append(errors, NewFieldError(InBody, "requestBody", "can't decode body from XML", err))
		}

	default:
		errors = append(errors, NewFieldError(InBody, "-", "unknown content type", nil))
	}
	return
}
`)
	})
}

func assertGeneratedCode(t *testing.T, def translator.TypeDef, expected string) {
	b := &strings.Builder{}
	if !assert.NoError(t, MakeBodyConstructor(b, def)) {
		return
	}

	result := strings.NewReader("package api\n" + b.String())

	prettyResult := &strings.Builder{}
	if !assert.NoError(t, barber.PrettifySource(result, prettyResult)) {
		return
	}

	assert.Equal(t, expected, prettyResult.String())
}
