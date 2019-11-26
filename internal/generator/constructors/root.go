package constructors

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

/*
MakeRequestConstructor receive a root request struct definition and generate corresponding constructor, for example

	type IncomeRequest struct {
		Body   IncomeRequestBody
		Cookie IncomeRequestCookie
		Header IncomeRequestHeader
		Path   IncomeRequestPath
		Query  IncomeRequestQuery
	}

	func MakeIncomeRequest(c *gin.Context) (result IncomeRequest, errors []FieldError) {
		result.Body, errors = MakeIncomeRequestBody(c)
		if errors != nil {
			return
		}

		result.Cookie, errors = MakeIncomeRequestCookie(c)
		if errors != nil {
			return
		}

		result.Header, errors = MakeIncomeRequestHeader(c)
		if errors != nil {
			return
		}

		result.Path, errors = MakeIncomeRequestPath(c)
		if errors != nil {
			return
		}

		result.Query, errors = MakeIncomeRequestQuery(c)
		if errors != nil {
			return
		}
		return
	}
*/
func MakeRequestConstructor(wr io.Writer, def translator.TypeDef) error {
	return constructorTemplate.Execute(wr, def)
}

var constructorTemplate = template.Must(template.New("rootRequestConstructor").Parse(`
func Make{{ .Name }}(c *gin.Context) (result {{ .Name }}, errors []FieldError) {
	{{- range $, $field := .Fields }}
		result.{{ .Name }}, errors = Make{{ .GoType }}(c)
		if errors != nil {
			return
		}
	{{ end -}}
	return
}
`))
