package fields

import (
	"testing"

	"github.com/kepkin/gorest/internal/generator/translator"
	"github.com/stretchr/testify/assert"
)

func TestMakeFileFieldConstructor(t *testing.T) {
	t.Run("No file field", func(t *testing.T) {
		_, err := MakeFileFieldConstructor(translator.Field{
			Type: translator.FileField + 1,
		}, "InFormData")
		assert.Error(t, err)
	})

	t.Run("File field example", func(t *testing.T) {
		s, err := MakeFileFieldConstructor(translator.Field{
			Name:      "Document",
			Parameter: "doc",
			GoType:    "*multipart.FileHeader",
			Type:      translator.FileField,
		}, "InFormData")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `result.Document, err = c.FormFile("doc")
if err != nil {
	errors = append(errors, NewFieldError(InFormData, "doc", "can't extract file from form-data", err))
}`, s)
	})
}
