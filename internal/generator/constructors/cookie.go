package constructors

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/constructors/fields"
	"github.com/kepkin/gorest/internal/generator/translator"
)

// MakeCookieParamsConstructor receive a cookie params struct definition and generate corresponding constructor
func MakeCookieParamsConstructor(wr io.Writer, def translator.TypeDef) error {
	return cookieParamsConstructorTemplate.Execute(wr, def)
}

var cookieParamsConstructorTemplate = template.Must(template.New("cookieParamsConstructor").Funcs(fields.BaseConstructor).Parse(`
func Make{{ .Name }}(c *gin.Context) (result {{ .Name }}, errors []FieldError) {
	{{- if .HasNoStringFields }}
	var err error
	{{ end }}

	{{- with .Fields }}
		getCookie := func(param string) (string, bool) {
			cookie, err := c.Request.Cookie(param)
			if err == http.ErrNoCookie {
				return "", false
			}
			return cookie.Value, true
		}
	{{- end }}

	{{ range $, $field := .Fields }}
	{{- with $field }}
		{{- if .CheckDefault}}
			{{ .StrVarName }}, ok := getCookie("{{ .Parameter }}")
			if !ok {
			   {{ .StrVarName }} = "{{ .Schema.Default }}"
			}
		{{- else }}
			{{ .StrVarName }}, _ := getCookie("{{ .Parameter }}")
		{{- end }}

		{{- BaseValueFieldConstructor . "InCookie" }}

	{{- end -}}
	{{ end -}}
	return
}
`))
