package constructors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/barber"
	"github.com/kepkin/gorest/internal/generator/translator"
	"github.com/kepkin/gorest/internal/spec/openapi3"
)

func TestMakeFormDataConstructor(t *testing.T) {
	defaultAge := "22"

	def := translator.TypeDef{
		Name: "UserProfileRequestBodyForm",
		Fields: []translator.Field{
			{Name: "Name", GoType: "string", Parameter: "name", Type: translator.StringField},
			{Name: "Age", GoType: "int64", Parameter: "age", Type: translator.IntegerField, Schema: openapi3.SchemaType{Default: &defaultAge}},
			{Name: "Photo", GoType: "*multipart.FileHeader", Parameter: "photo", Type: translator.FileField},
		},
	}

	b := &strings.Builder{}
	if !assert.NoError(t, MakeFormDataConstructor(b, def)) {
		return
	}
	result := strings.NewReader("package api\n" + b.String())

	prettyResult := &strings.Builder{}
	if !assert.NoError(t, barber.PrettifySource(result, prettyResult)) {
		return
	}

	assert.Equal(t, `package api

func MakeUserProfileRequestBodyForm(c *gin.Context) (result UserProfileRequestBodyForm, errors []FieldError) {
	var err error

	form, err := c.MultipartForm()
	if err != nil {
		errors = append(errors, NewFieldError(InFormData, "", "can't parse multipart form", err))
		return
	}

	getFormValue := func(param string) (string, bool) {
		values, ok := form.Value[param]
		if !ok {
			return "", false
		}
		if len(values) == 0 {
			return "", false
		}
		return values[0], true
	}

	nameStr, _ := getFormValue("name")
	result.Name = nameStr

	ageStr, ok := getFormValue("age")
	if !ok {
		ageStr = "22"
	}
	result.Age, err = strconv.ParseInt(ageStr, 10, 0)
	if err != nil {
		errors = append(errors, NewFieldError(InFormData, "age", "can't parse as integer", err))
	}

	result.Photo, err = c.FormFile("photo")
	if err != nil {
		errors = append(errors, NewFieldError(InFormData, "photo", "can't extract file from form-data", err))
	}
	return
}
`, prettyResult.String())
}

func TestMakeFormDataConstructorForOneFileField(t *testing.T) {
	def := translator.TypeDef{
		Name: "UploadDocumentRequestBodyForm",
		Fields: []translator.Field{
			{Name: "Document", GoType: "*multipart.FileHeader", Parameter: "document", Type: translator.FileField},
		},
	}

	b := &strings.Builder{}
	if !assert.NoError(t, MakeFormDataConstructor(b, def)) {
		return
	}
	result := strings.NewReader("package api\n" + b.String())

	prettyResult := &strings.Builder{}
	if !assert.NoError(t, barber.PrettifySource(result, prettyResult)) {
		return
	}

	assert.Equal(t, `package api

func MakeUploadDocumentRequestBodyForm(c *gin.Context) (result UploadDocumentRequestBodyForm, errors []FieldError) {
	var err error

	result.Document, err = c.FormFile("document")
	if err != nil {
		errors = append(errors, NewFieldError(InFormData, "document", "can't extract file from form-data", err))
	}
	return
}
`, prettyResult.String())
}
