package generator

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/openapi3/spec"
	"github.com/kepkin/gorest/internal/openapi3/translator"
)

var handlerTemplate = template.Must(template.New("handler").Parse(`
func (server {{ .InterfaceName }}Server) _{{ .InterfaceName }}_{{ .Path.OperationID }}_Handler(c *gin.Context) {
	c.Set(handlerNameKey, "{{ .Path.OperationID }}")

	req, errors := Make{{ .Path.OperationID }}Request(c)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "Parse request error",
			Errors:  errors,
		})
		return
	}

	errors = req.Validate()
	if len(errors) > 0 {
		c.JSON(http.StatusUnprocessableEntity, Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Validate request error",
			Errors:  errors,
		})
	}

	server.Srv.{{ .Path.OperationID }}(req, c)
}
`))

func MakeHandlers(wr io.Writer, sp spec.Spec) error {
	interfaceName := translator.MakeIdentifier(sp.Info.Title)

	for _, path := range sp.Paths {
		for _, method := range path.Methods() {
			if method != nil {
				if err := makeHandler(wr, interfaceName, method); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func makeHandler(wr io.Writer, interfaceName string, path *spec.PathSpec) error {
	return handlerTemplate.Execute(wr, struct {
		InterfaceName string
		Path          *spec.PathSpec
	}{
		InterfaceName: interfaceName,
		Path:          path,
	})
}
