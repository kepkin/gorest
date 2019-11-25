package translator

import (
	"testing"

	"github.com/kepkin/gorest/internal/spec/openapi3"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

const schemaExample = `
---
type: object
properties:
    payment_id:
      type: string
    
    merchant_id:
      type: integer
    
    sum:
      type: number
      format: decimal
    
    meta:
      type: object
      additionalProperties:
        type: object
`

func TestProcessRootSchema(t *testing.T) {
	var schema openapi3.SchemaType

	assert.NoError(t, yaml.Unmarshal([]byte(schemaExample), &schema))
	schema.Name = "Payment"

	defs, err := ProcessRootSchema(schema)
	if !assert.NoError(t, err) {
		return
	}
	for i := range defs {
		for j := range defs[i].Fields {
			defs[i].Fields[j].Schema = openapi3.SchemaType{} // TODO(a.telyshev)
		}
	}

	if !assert.Equal(t, len(defs), 1) {
		return
	}
	def := defs[0]

	assert.Equal(t, "Payment", def.Name)
	assert.Equal(t, "struct", def.GoType)
	assert.ElementsMatch(t, def.Fields, []Field{
		{Parameter: "payment_id", Name: "PaymentID", GoType: "string", Type: StringField},
		{Parameter: "merchant_id", Name: "MerchantID", GoType: "int64", Type: IntegerField},
		{Parameter: "sum", Name: "Sum", GoType: "Decimal", Type: CustomField},
		{Parameter: "meta", Name: "Meta", GoType: "json.RawMessage", Type: FreeFormObject},
	})
}
