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

var formDataConstructorTemplate = template.Must(template.New("formDataConstructor").Funcs(fields.Constructors).Parse(`
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
		
		{{- if .IsString }}
			result.{{ .Name }}, _ = getFormValue("{{ .Parameter }}")
		{{- end }}

		{{- if .IsCustom }}
			{{ .StrVarName }}, _ := getFormValue("{{ .Parameter }}")
			{{ CustomFieldConstructor . "InFormData" }}
		{{- end }}

		{{- if .IsInteger }}
			{{ .StrVarName }}, _ := getFormValue("{{ .Parameter }}")
			{{ IntConstructor . "InFormData" }}
		{{- end }}

		{{- if .IsFloat }}
			{{ .StrVarName }}, _ := getFormValue("{{ .Parameter }}")
			{{ FloatConstructor . "InFormData" }}
		{{- end }}

		{{- if or .IsDate .IsDateTime }}
			{{ .StrVarName }}, _ := getFormValue("{{ .Parameter }}")
			{{ DateTimeConstructor . "InFormData" }}
		{{- end }}

        {{- if .IsDateTime }}
			{{ .StrVarName }}, _ := getFormValue("{{ .Parameter }}")
			{{ DateTimeConstructor . "InFormData" }}
		{{- end }}

		{{- if .IsUnixTime }}
			{{ .StrVarName }}, _ := getFormValue("{{ .Parameter }}")
			{{ UnixTimeConstructor . "InFormData" }}
		{{- end }}

		{{- if .IsFile }}
			{{ FileConstructor . "InFormData" }}
		{{- end }}

	{{- end }}
	{{ end -}}
	return
}
`))
