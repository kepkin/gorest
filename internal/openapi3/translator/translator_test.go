package translator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"

	"github.com/kepkin/gorest/internal/openapi3/spec"
)

const schemaExample = `
---
type: object
properties:
    payment_id:
      type: string
    
    merchant_id:
      type: string
    
    sum:
      type: number
      format: decimal
    
    meta:
      type: object
`

func TestProcessRootSchema(t *testing.T) {
	var schema spec.SchemaType

	assert.NoError(t, yaml.Unmarshal([]byte(schemaExample), &schema))
	schema.Name = "Payment"

	defs, err := processRootSchema(schema)
	assert.NoError(t, err)

	assert.ElementsMatch(t, []Definition{
		TypeDef{
			Name: "Payment",
			Fields: []Field{
				{Name: "payment_id", Type: "string"},
				{Name: "merchant_id", Type: "string"},
				{Name: "sum", Type: "Decimal"},
				{Name: "meta", Type: "PaymentMeta"},
			},
		},
		InterfaceCheckerDef{
			TypeName:      "Decimal",
			InterfaceName: "json.Marshaller",
		},
		TypeDef{
			Name:   "PaymentMeta",
		},
	}, defs)
}
