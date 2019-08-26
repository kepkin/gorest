package pkg

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"strings"
	"testing"
)

const stringAliasSpec = `type: string`
const stringAliasExpected = `type SutType string
`

func TestStringAlias(t *testing.T) {
	var sut SchemaType

	err := yaml.Unmarshal([]byte(stringAliasSpec), &sut)
	assert.Nil(t, err, "unexpected: %v", err)
	tpl, err := BuildTpls()
	assert.Nil(t, err, "unexpected: %v", err)

	result := strings.Builder{}
	err = tpl.ExecuteTemplate(&result, "componentSchema", map[string]interface{}{"name": "SutType", "desc": sut})
	assert.Nil(t, err, "unexpected: %v", err)

	assert.Equal(t, stringAliasExpected, result.String())
}

const arraySpec =
`type: array
items:
  $ref: "#/components/schemas/Author"
`

const arrayExpected = `type SutType []Author
`

func TestArray(t *testing.T) {
	var sut SchemaType

	err := yaml.Unmarshal([]byte(arraySpec), &sut)
	assert.Nil(t, err, "unexpected: %v", err)
	tpl, err := BuildTpls()
	assert.Nil(t, err, "unexpected: %v", err)

	result := strings.Builder{}
	err = tpl.ExecuteTemplate(&result, "componentSchema", map[string]interface{}{"name": "SutType", "desc": sut})
	assert.Nil(t, err, "unexpected: %v", err)

	assert.Equal(t, arrayExpected, result.String())
}

const objectSpec =
`type: object
properties:
  title:
    type: string
  authors_list:
    type: array
    items:
      $ref: "#/components/schemas/Author"
`

const objectExpected =
`type SutType struct {
    AuthorsList []Author `+"`json:\"authors_list\"`" + `
    Title string `+"`json:\"title\"`" + `
}
`

func TestObject(t *testing.T) {
	var sut SchemaType

	err := yaml.Unmarshal([]byte(objectSpec), &sut)
	assert.Nil(t, err, "unexpected: %v", err)
	tpl, err := BuildTpls()
	assert.Nil(t, err, "unexpected: %v", err)

	result := strings.Builder{}
	err = tpl.ExecuteTemplate(&result, "componentSchema", map[string]interface{}{"name": "SutType", "desc": sut})
	assert.Nil(t, err, "unexpected: %v", err)

	assert.Equal(t, objectExpected, result.String())
}