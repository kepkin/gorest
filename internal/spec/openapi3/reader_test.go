package openapi3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const specWithEnum = `
paths:
  /api/v1/resource:
    post:
      operationId: CreateResource
      parameters:
        - name: color
          in: query
          schema:
            type: string
            enum: [white, "blue", red]
        
        - name: sort
          in: query
          schema:
            type: string
            enum:
              - "asc"
              - desc
        
        - name: id
          in: path
          schema:
            type: integer
            enum: [10, 20, 30]
`

func TestReadEnum(t *testing.T) {
	sp, err := ReadSpec([]byte(specWithEnum))
	if !assert.NoError(t, err) {
		return
	}

	method := sp.Paths["/api/v1/resource"].Post
	if !assert.NotNil(t, method) {
		return
	}

	enums := make([][]string, 0, len(method.Parameters))
	for _, p := range method.Parameters {
		enums = append(enums, p.Schema.Enum)
	}

	assert.Equal(t, [][]string{
		{"white", "blue", "red"},
		{"asc", "desc"},
		{"10", "20", "30"},
	}, enums)
}
