package constructors

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

// MakeCookieParamsConstructor receive a cookie params struct definition and generate corresponding constructor
func MakeCookieParamsConstructor(wr io.Writer, def translator.TypeDef) error {
	return cookieParamsConstructorTemplate.Execute(wr, def)
}

var cookieParamsConstructorTemplate = template.Must(template.New("cookieParamsConstructor").Funcs(template.FuncMap{
	"IntConstructor":   makeIntConstructor,
	"FloatConstructor": makeFloatConstructor,
}).Parse(`
func Make{{ .Name }}(c *gin.Context) (result {{ .Name }}, errors []FieldError) {
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
		
		{{- if .IsString }}
			result.{{ .Name }}, _ = getCookie("{{ .Parameter }}")
		{{- end }}

		{{- if .IsInteger }}
			{{ .StrVarName }}, _ = getCookie("{{ .Parameter }}")
			{{ IntConstructor . "InCookie" }}
		{{- end }}

		{{- if .IsFloat }}
			{{ .StrVarName }}, _ = getCookie("{{ .Parameter }}")
			{{ FloatConstructor . "InCookie" }}
		{{- end }}

	{{- end }}
	{{ end -}}
	return
}
`))
