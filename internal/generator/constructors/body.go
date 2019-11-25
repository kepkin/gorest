package constructors

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

// MakeBodyConstructor receive a body struct definition and generate corresponding constructor
func MakeBodyConstructor(wr io.Writer, def translator.TypeDef) error {
	return bodyConstructorTemplate.Execute(wr, def)
}

var bodyConstructorTemplate = template.Must(template.New("bodyConstructor").Parse(`
func Make{{ .Name }}(c *gin.Context) (result {{ .Name }}, errors []FieldError) {
	switch c.Request.Header.Get("Content-Type") {

	{{- range $, $field := .Fields }}
	{{- with $field }}
		
		{{- if eq .Name "JSON" }}
		case "application/json":
			result.Type = AppJSON
			if err := json.NewDecoder(c.Request.Body).Decode(result.JSON); err != nil {
				errors = append(errors, NewFieldError(InBody, "requestBody", "can't decode body from JSON", err))
			}
		{{ end }}

		{{- if eq .Name "XML" }}
		case "application/xml":
			result.Type = AppXML
			if err := xml.NewDecoder(c.Request.Body).Decode(result.XML); err != nil {
				errors = append(errors, NewFieldError(InBody, "requestBody", "can't decode body from XML", err))
			}
		{{ end }}{{ end }}{{ end -}}
	}
	return
}
`))
