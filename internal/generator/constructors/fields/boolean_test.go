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

	t.Run("Correct boolean", func(t *testing.T) {
		s, err := MakeBooleanFieldConstructor(translator.Field{
			Name:      "Flag",
			Parameter: "flag",
			Type:      translator.BooleanField,
		}, "InPath")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `switch strings.ToLower(flagStr) {
    case "1", "true", "t":
        result.Flag = true
    case "0", "false", "f":
        result.Flag = false
    default:
        errors = append(errors, NewFieldError(InPath, "flag", "can't parse as boolean", nil))
}`, s)
	})
}
