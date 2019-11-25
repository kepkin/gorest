package constructors

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/generator/translator"
	"github.com/kepkin/gorest/internal/spec/openapi3"
)

func TestMakeIntConstructor(t *testing.T) {
	t.Run("No int field", func(t *testing.T) {
		_, err := makeIntFieldConstructor(translator.Field{
			Type: translator.IntegerField + 1,
		}, "InQuery")
		assert.Error(t, err)
	})

	t.Run("No bit size", func(t *testing.T) {
		s, err := makeIntFieldConstructor(translator.Field{
			Name:      "ID",
			Parameter: "id",
			Type:      translator.IntegerField,
		}, "InPath")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `result.ID, err = strconv.ParseInt(idStr, 10, 0)
if err != nil {
	errors = append(errors, NewFieldError(InPath, "id", "can't parse as integer", err))
}`, s)
	})

	t.Run("32 bit field", func(t *testing.T) {
		s, err := makeIntFieldConstructor(translator.Field{
			Name:      "ID",
			Parameter: "id",
			Type:      translator.IntegerField,
			Schema:    openapi3.SchemaType{NumberSchema: openapi3.NumberSchema{BitSize: 32}},
		}, "InHeader")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `result.ID, err = strconv.ParseInt(idStr, 10, 32)
if err != nil {
	errors = append(errors, NewFieldError(InHeader, "id", "can't parse as 32 bit integer", err))
}`, s)
	})

	t.Run("64 bit field", func(t *testing.T) {
		s, err := makeIntFieldConstructor(translator.Field{
			Name:      "ID",
			Parameter: "id",
			Type:      translator.IntegerField,
			Schema:    openapi3.SchemaType{NumberSchema: openapi3.NumberSchema{BitSize: 64}},
		}, "InQuery")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `result.ID, err = strconv.ParseInt(idStr, 10, 64)
if err != nil {
	errors = append(errors, NewFieldError(InQuery, "id", "can't parse as 64 bit integer", err))
}`, s)
	})
}

func TestMakeFloatConstructor(t *testing.T) {
	t.Run("No float field", func(t *testing.T) {
		_, err := makeFloatFieldConstructor(translator.Field{
			Type: translator.FloatField + 1,
		}, "InCookie")
		assert.Error(t, err)
	})

	t.Run("No bit size", func(t *testing.T) {
		s, err := makeFloatFieldConstructor(translator.Field{
			Name:      "Sum",
			Parameter: "sum",
			Type:      translator.FloatField,
		}, "InPath")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `result.Sum, err = strconv.ParseFloat(sumStr, 10, 0)
if err != nil {
	errors = append(errors, NewFieldError(InPath, "sum", "can't parse as float", err))
}`, s)
	})

	t.Run("32 bit field", func(t *testing.T) {
		s, err := makeFloatFieldConstructor(translator.Field{
			Name:      "Sum",
			Parameter: "sum",
			Type:      translator.FloatField,
			Schema:    openapi3.SchemaType{NumberSchema: openapi3.NumberSchema{BitSize: 32}},
		}, "InPath")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `result.Sum, err = strconv.ParseFloat(sumStr, 10, 32)
if err != nil {
	errors = append(errors, NewFieldError(InPath, "sum", "can't parse as 32 bit float", err))
}`, s)
	})

	t.Run("64 bit field", func(t *testing.T) {
		s, err := makeFloatFieldConstructor(translator.Field{
			Name:      "Sum",
			Parameter: "sum",
			Type:      translator.FloatField,
			Schema:    openapi3.SchemaType{NumberSchema: openapi3.NumberSchema{BitSize: 64}},
		}, "InQuery")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `result.Sum, err = strconv.ParseFloat(sumStr, 10, 64)
if err != nil {
	errors = append(errors, NewFieldError(InQuery, "sum", "can't parse as 64 bit float", err))
}`, s)
	})
}

func TestMakeCustomFieldConstructor(t *testing.T) {
	t.Run("No custom field", func(t *testing.T) {
		_, err := makeCustomFieldConstructor(translator.Field{
			Type: translator.CustomField + 1,
		}, "InCookie")
		assert.Error(t, err)
	})

	t.Run("Custom field example", func(t *testing.T) {
		s, err := makeCustomFieldConstructor(translator.Field{
			Name:      "Sum",
			Parameter: "sum",
			GoType:    "Decimal",
			Type:      translator.CustomField,
		}, "InQuery")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `result.Sum = Decimal{}
if err = result.Sum.SetFromString(sumStr); err != nil {
	errors = append(errors, NewFieldError(InQuery, "sum", fmt.Sprintf("can't create from string '%s'", sumStr), err))
}`, s)
	})
}
