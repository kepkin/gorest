//nolint:dupl
package fields

import (
	"testing"

	"github.com/kepkin/gorest/internal/generator/translator"
	"github.com/stretchr/testify/assert"
)

func TestMakeBooleanConstructor(t *testing.T) {
	t.Run("No Boolean field", func(t *testing.T) {
		_, err := MakeBooleanFieldConstructor(translator.Field{
			Type: translator.BooleanField + 1,
		}, "InCookie")
		assert.Error(t, err)
	})

	t.Run("No bit size", func(t *testing.T) {
		s, err := MakeBooleanFieldConstructor(translator.Field{
			Name:      "Flag",
			Parameter: "flag",
			Type:      translator.BooleanField,
		}, "InPath")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `flagStr = strings.strings.ToLower(flagStr)
switch flagStr {
    case "1", "true", "t":
        result.Flag = true
    default:
        result.Flag = false
}`, s)
	})
}
