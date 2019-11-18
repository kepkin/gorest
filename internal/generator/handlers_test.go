package generator

import (
	"testing"

	"github.com/kepkin/gorest/internal/spec/openapi3"

	"github.com/stretchr/testify/assert"
)

const pathSpecYaml = `
paths:
  /api/v1/resource:
    post:
      operationId: CreateResource
      parameters:
        - name: id
          in: query
          required: true
          schema:
            type: string
    
        - name: from_date
          in: query
          required: true
          schema:
            type: integer
            format: int64

        - name: creds
          in: query
          schema:
            type: object
            properties:
              name:
                type: string
    
        - in: header
          name: X-access-token
          required: true
          schema:
            type: string
`

func TestMakeHandler(t *testing.T) {
	_, err := openapi3.ReadSpec([]byte(pathSpecYaml))
	assert.NoError(t, err)

	//assert.NoError(t,
	//	makeHandlers(*spec.Paths["/api/v1/resource"].Post))
}
