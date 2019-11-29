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

var cookieParamsConstructorTemplate = template.Must(template.New("cookieParamsConstructor").Funcs(fields.Constructors).Parse(`
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
		
		{{- if .IsString }}
			result.{{ .Name }}, _ = getCookie("{{ .Parameter }}")
		{{- end }}

		{{- if .IsCustom }}
			{{ .StrVarName }}, _ := getCookie("{{ .Parameter }}")
			{{ CustomFieldConstructor . "InCookie" }}
		{{- end }}

		{{- if .IsInteger }}
			{{ .StrVarName }}, _ := getCookie("{{ .Parameter }}")
			{{ IntConstructor . "InCookie" }}
		{{- end }}

		{{- if .IsFloat }}
			{{ .StrVarName }}, _ := getCookie("{{ .Parameter }}")
			{{ FloatConstructor . "InCookie" }}
		{{- end }}

		{{- if or .IsDate .IsDateTime }}
			{{ .StrVarName }}, _ := c.GetQuery("{{ .Parameter }}")
			{{ DateTimeConstructor . "InCookie" }}
		{{- end }}

        {{- if .IsDateTime }}
			{{ .StrVarName }}, _ := c.GetQuery("{{ .Parameter }}")
			{{ DateTimeConstructor . "InCookie" }}
		{{- end }}

		{{- if .IsUnixTime }}
			{{ .StrVarName }}, _ := c.GetQuery("{{ .Parameter }}")
			{{ UnixTimeConstructor . "InCookie" }}
		{{- end }}

	{{- end }}
	{{ end -}}
	return
}
`))
