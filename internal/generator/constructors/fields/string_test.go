package fields

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/generator/translator"
	"github.com/kepkin/gorest/internal/spec/openapi3"
)

func TestMakeStringFieldConstructor(t *testing.T) {
	t.Run("No string field", func(t *testing.T) {
		_, err := MakeIntFieldConstructor(translator.Field{
			Type: translator.StringField + 1,
		}, "InQuery")
		assert.Error(t, err)
	})

	t.Run("Correct string", func(t *testing.T) {
		s, err := MakeStringFieldConstructor(translator.Field{
			Name:      "Name",
			Parameter: "Name",
			Type:      translator.StringField,
			Schema:    openapi3.SchemaType{NumberSchema: openapi3.NumberSchema{BitSize: 32}},
		}, "InQuery")
		fmt.Println(s)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `result.Name = nameStr
`, s)
	})
}
