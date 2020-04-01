package fields

import (
	"fmt"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/translator"
)

var makeValueFieldTemplate = template.Must(template.New("setValue").Funcs(Constructors).Parse(`
	{{- if .IsBoolean }}
		{{ BooleanConstructor .Field .Place}}
	{{- end }}

	{{- if .IsString }}
		{{ StringConstructor .Field .Place}}
	{{- end }}
	
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
	if place == "InPath" {
		if f.IsStruct() || f.IsArray() || f.IsFile() {
			return "" , fmt.Errorf("unsupported type in path")
		}
	}

	return executeFieldConstructorTemplate(makeValueFieldTemplate, f, place)
}
