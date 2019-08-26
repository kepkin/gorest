package pkg

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"strings"
	"testing"
)

const pathSpecYaml = `
summary: Adds resource
tags: [srv]
description: Adds new resource and updates old ones if they exists
operationId: CreateResource
parameters:
  - name: id
    in: path
    required: true
    schema:
      type: string

  - in: query
    name: from_date
    required: true
    schema:
      type: integer
      format: int64

  - in: header
    name: X-access-token

    required: true
    schema:
      type: string

requestBody:
  description: "Create test request"
  required: true
  content:
    application/json:
      schema:
        $ref: "#/components/schemas/Resource"

responses:
  '200':
     description: A successful response.
  '400':
     description: Error on invalid input
     content:
       application/json:
         schema:
           $ref: '#/components/schemas/Resource'
`


const pathSpecWithoutRequestBodyYaml = `
summary: Adds resource
tags: [srv]
description: Adds new resource and updates old ones if they exists
operationId: CreateResource
parameters:
  - name: id
    in: path
    required: true
    schema:
      type: string

responses:
  '200':
     description: A successful response.
  '400':
     description: Error on invalid input
     content:
       application/json:
         schema:
           $ref: '#/components/schemas/Resource'
`

const requestTypeDeclarationExpected =
`type CreateResourceReq struct {
    Id string
    FromDate int64
    XAccessToken string
    Body Resource
}`

func TestRequestTypeDeclaration(t *testing.T) {
	var sut PathSpec

	err := yaml.Unmarshal([]byte(pathSpecYaml), &sut)
	assert.Nil(t, err)
	tpl, err := BuildTpls()
	assert.Nil(t, err)

	result := strings.Builder{}
	err = tpl.ExecuteTemplate(&result, "reqStructDecl", sut)
	assert.Nil(t, err)

	assert.Equal(t, requestTypeDeclarationExpected, result.String())
}

const requestTypeConstructorExpected =
`func MakeCreateResourceReq(c *gin.Context) (res CreateResourceReq, err error) {
    var errors strings.Builder

	// URI binding
	for _, param := range c.Params {
		switch param.Key {
		case "id":
          res.Id = param.Value

		default:
			fmt.Fprintf(&errors, "Unexpected uri parameter: ` + "`%s`" + `", param.Key)
		}
	}
    res.FromDate = strconv.ParseInt(c.Query("from_date"), 10, 64)
    res.XAccessToken = c.Request.Header.Get("X-access-token")

    if c.Request == nil || c.Request.Body == nil {
        fmt.Fprintf(&errors, "Body is absent")
    }

    switch ct := c.Request.Header.Get("Content-Type"); ct {
    case "application/json":
        dec := json.NewDecoder(c.Request.Body)
    	if err := dec.Decode(&res.Body); err != nil {
    		fmt.Fprintf(&errors, "Cant parse JSON") //TODO: investigate for more detailed error
    	}
    default:
        fmt.Fprintf(&errors, "Unsupported content-type: %v", ct)
    }

	if errors.Len() > 0 {
		err = fmt.Errorf(errors.String())
	}
	return
}`

func TestRequestTypeConstructor(t *testing.T) {
	var sut PathSpec

	err := yaml.Unmarshal([]byte(pathSpecYaml), &sut)
	assert.Nil(t, err)
	tpl, err := BuildTpls()
	assert.Nil(t, err)

	result := strings.Builder{}
	err = tpl.ExecuteTemplate(&result, "constructorDecl", sut)
	assert.Nil(t, err, "Unexpected errors: %v", err)

	assert.Equal(t, requestTypeConstructorExpected, result.String())
}

const requestTypeConstructorWithoutReqBodyExpected =
	`func MakeCreateResourceReq(c *gin.Context) (res CreateResourceReq, err error) {
    var errors strings.Builder

	// URI binding
	for _, param := range c.Params {
		switch param.Key {
		case "id":
          res.Id = param.Value

		default:
			fmt.Fprintf(&errors, "Unexpected uri parameter: ` + "`%s`" + `", param.Key)
		}
	}

    if c.Request == nil || c.Request.Body != nil {
        fmt.Fprintf(&errors, "Unexpected body")
    }

	if errors.Len() > 0 {
		err = fmt.Errorf(errors.String())
	}
	return
}`

func TestRequestTypeConstructorWithoutRequestBody(t *testing.T) {
	var sut PathSpec

	err := yaml.Unmarshal([]byte(pathSpecWithoutRequestBodyYaml), &sut)
	assert.Nil(t, err)
	tpl, err := BuildTpls()
	assert.Nil(t, err)

	result := strings.Builder{}
	err = tpl.ExecuteTemplate(&result, "constructorDecl", sut)
	assert.Nil(t, err)

	assert.Equal(t, requestTypeConstructorWithoutReqBodyExpected, result.String())
}
