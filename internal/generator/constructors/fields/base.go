package fields

import (
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var makeValueFieldTemplate = template.Must(template.New("setValue").Funcs(Constructors).Parse(`
	{{- if .IsCustom }}
		{{ CustomFieldConstructor .Field .Place }}
	{{- end }}
	
	{{- if .IsInteger }}
		{{ IntConstructor .Field .Place }}
	{{- end }}
	
	{{- if .IsFloat }}
		{{ FloatConstructor .Field .Place }}
	{{- end }}
	
	{{- if .IsDate }}
		{{ DateConstructor .Field .Place }}
	{{- end }}
	
	{{- if .IsDateTime }}
		{{ DateTimeConstructor .Field .Place }}
	{{- end }}
	
	{{- if .IsUnixTime }}
		{{ UnixTimeConstructor .Field .Place }}
	{{- end }}
	
	{{- if .IsFile }}
		{{ FileConstructor .Field .Place }}
	{{- end }}
`))

func MakeValueFieldConstructor(f translator.Field, place string) (string, error) {
	return executeFieldConstructorTemplate(makeValueFieldTemplate, f, place)
}
