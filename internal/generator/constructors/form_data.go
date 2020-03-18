package constructors

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/constructors/fields"
	"github.com/kepkin/gorest/internal/generator/translator"
)

// MakeFormDataConstructor receive a form-data struct definition and generate corresponding constructor
func MakeFormDataConstructor(wr io.Writer, def translator.TypeDef) error {
	return formDataConstructorTemplate.Execute(wr, def)
}

var formDataConstructorTemplate = template.Must(template.New("formDataConstructor").Funcs(fields.BaseConstructor).Parse(`
func Make{{ .Name }}(c *gin.Context) (result {{ .Name }}, errors []FieldError) {
	{{- if .HasNoStringFields }}
	var err error
	{{ end }}
	
	{{- with .Fields }}
	{{ if $.HasNoFileFields }}
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
	{{ end }}
	{{- end }}

	{{ range $, $field := .Fields }}
	{{- with $field }}
		{{- if not .IsFile}}
			{{- if .CheckDefault}}
				{{ .StrVarName }}, ok := getFormValue("{{ .Parameter }}")
				if !ok {
				   {{ .StrVarName }} = "{{ .Schema.Default }}"
				}
			{{- else }}
				{{ .StrVarName }}, _ := getFormValue("{{ .Parameter }}")
			{{- end }}
		{{- end }}
		
		{{- BaseValueFieldConstructor . "InFormData" }}

	{{- end -}}
	{{ end -}}
	return
}
`))
