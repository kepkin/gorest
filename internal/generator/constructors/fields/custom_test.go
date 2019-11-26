package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/generator/translator"
)

func TestMakeCustomFieldConstructor(t *testing.T) {
	t.Run("No custom field", func(t *testing.T) {
		_, err := MakeCustomFieldConstructor(translator.Field{
			Type: translator.CustomField + 1,
		}, "InCookie")
		assert.Error(t, err)
	})

	t.Run("Custom field example", func(t *testing.T) {
		s, err := MakeCustomFieldConstructor(translator.Field{
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
