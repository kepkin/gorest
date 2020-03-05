package generator

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/kepkin/gorest/internal/translator"
)

var validateTemplate = template.Must(template.New("validate").Funcs(template.FuncMap{
	"MakeFieldValidators": makeFieldValidators,
}).Parse(`
func (t {{ .Name }}) Validate() (errors []FieldError) {
	{{- range $, $field := .Fields }}
		{{- if ne $field.GoType "ContentType" }}
			// {{ $field.Name }} field validators{{ if or $field.IsStruct $field.IsComponent }}
				errors = t.{{ $field.Name }}.Validate()
				if errors != nil {
					return
				}
			{{- else -}}
				{{ MakeFieldValidators $field }}
			{{- end -}}
		{{- end -}}
	{{ end }}
	return
}
`))

func (g Generator) makeValidateFunc(wr io.Writer, def translator.TypeDef) error {
	return validateTemplate.Execute(wr, def)
}

type validatorGenerator func(translator.Field) (string, error)

func makeFieldValidators(f translator.Field) (string, error) {
	result := strings.Builder{}

	for _, makeValidator := range [...]validatorGenerator{
		makeEnumValidator,
	} {
		v, err := makeValidator(f)
		if err != nil {
			return "", err
		}
		if v != "" {
			result.WriteString(v)
		}
	}
	return result.String(), nil
}

var enumValidator = template.Must(template.New("enumValidator").Parse(`
var {{ .Parameter }}InEnum bool
for _, elem := range [...]{{ .GoType }}{
	{{- range $_, $val := .Schema.Enum }}
		{{- if $.IsString }}
			"{{ $val }}",
		{{- else }}
			{{ $val }},
		{{- end }}
	{{- end }}
} {
	if elem == t.{{ $.Name }} {
		{{ .Parameter }}InEnum = true
		break
	}
}
if !{{ .Parameter }}InEnum {
	errors = append(errors, NewFieldError(UndefinedPlace, "{{ .Parameter }}", "allowed values: {{ .Schema.Enum }}", nil))
}
`))

func makeEnumValidator(f translator.Field) (string, error) {
	//if len(f.Schema.Enum) == 0 {
		return "", nil
	//}

	switch f.Type {
	case translator.IntegerField, translator.FloatField, translator.StringField:
		break
	default:
		return "", fmt.Errorf("enum doen't support for type: %s", f.Schema.Type)
	}

	b := &bytes.Buffer{}
	if err := enumValidator.Execute(b, f); err != nil {
		return "", err
	}
	return b.String(), nil
}
