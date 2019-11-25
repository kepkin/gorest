package constructors

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var parseIntTemplate = template.Must(template.New("parseInt").Parse(
	`result.{{ .Name }}, err = strconv.ParseInt({{ .StrVarName }}, 10, {{ .Schema.BitSize }})
if err != nil {
	{{- if eq .Schema.BitSize 0 }}
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as integer", err))
	{{- else }}
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as {{ .Schema.BitSize }} bit integer", err))
	{{- end }}
}`))

func makeIntFieldConstructor(f translator.Field, place string) (string, error) {
	if f.Type != translator.IntegerField {
		return "", fmt.Errorf("%v isn't integer field", f)
	}
	wr := &strings.Builder{}
	err := parseIntTemplate.Execute(wr, struct {
		translator.Field
		Place string
	}{
		Field: f,
		Place: place,
	})
	return wr.String(), err
}

var parseFloatTemplate = template.Must(template.New("parseFloat").Parse(
	`result.{{ .Name }}, err = strconv.ParseFloat({{ .StrVarName }}, 10, {{ .Schema.BitSize }})
if err != nil {
	{{- if eq .Schema.BitSize 0 }}
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as float", err))
	{{- else }}
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", "can't parse as {{ .Schema.BitSize }} bit float", err))
	{{- end }}
}`))

func makeFloatFieldConstructor(f translator.Field, place string) (string, error) {
	if f.Type != translator.FloatField {
		return "", fmt.Errorf("%v isn't float field", f)
	}
	wr := &strings.Builder{}
	err := parseFloatTemplate.Execute(wr, struct {
		translator.Field
		Place string
	}{
		Field: f,
		Place: place,
	})
	return wr.String(), err
}

var setFromStringTemplate = template.Must(template.New("setFromString").Parse(
	`result.{{ .Name }} = {{ .GoType }}{}
if err = result.{{ .Name }}.SetFromString({{ .StrVarName }}); err != nil {
	errors = append(errors, NewFieldError({{ .Place }}, "{{ .Parameter }}", fmt.Sprintf("can't create from string '%s'", {{ .StrVarName }}), err))
}`))

func makeCustomFieldConstructor(f translator.Field, place string) (string, error) {
	if f.Type != translator.CustomField {
		return "", fmt.Errorf("%v isn't field of user defined type", f)
	}
	wr := &strings.Builder{}
	err := setFromStringTemplate.Execute(wr, struct {
		translator.Field
		Place string
	}{
		Field: f,
		Place: place,
	})
	return wr.String(), err
}
