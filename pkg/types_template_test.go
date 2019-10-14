package pkg

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"strings"
	"testing"
)

const stringAliasSpec = `type: string`
const stringAliasExpected = `
        type SutType string`

const arraySpec = `type: array
items:
  $ref: "#/components/schemas/Author"
`
const arrayExpected = `
        type SutType []Author`

const objectSpec = `type: object
properties:
  title:
    type: string
  authors_list:
    type: array
    items:
      $ref: "#/components/schemas/Author"
`
const objectExpected = `
            type SutType struct {
                AuthorsList []Author ` + "`json:\"authors_list\"`" + `
                Title string ` + "`json:\"title\"`" + `
            }`

const additionalPropsSpec = `type: object
additionalProperties:
  type: object
`
const additionalPropsExpected = `
            type SutType = json.RawMessage`

func TestTemplates(t *testing.T) {
	tests := []struct {
		Name     string
		Spec     string
		Expected string
	}{
		{"StringAlias", stringAliasSpec, stringAliasExpected},
		{"Array", arraySpec, arrayExpected},
		{"Object", objectSpec, objectExpected},
		{"AdditionalProps", additionalPropsSpec, additionalPropsExpected},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var sut schemaType

			assert.NoError(t, yaml.Unmarshal([]byte(tt.Spec), &sut))
			assert.NoError(t, determineTypes(nil, &sut))
			assert.NoError(t, walkObjectProperties(nil, sut.Properties, determineTypes))

			tpl, err := BuildTemplates()
			assert.NoError(t, err)

			result := strings.Builder{}
			err = tpl.ExecuteTemplate(&result, "componentSchema", map[string]interface{}{"schemaName": "SutType", "schema": sut})
			assert.NoError(t, err)

			assert.Equal(t, tt.Expected, result.String())
		})
	}
}
