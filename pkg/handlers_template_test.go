package pkg

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
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

const requestTypeDeclarationExpected = `type CreateResourceReq struct {
            Id string
            FromDate int64
            XAccessToken string
                Body Resource
    }`

func TestRequestTypeDeclaration(t *testing.T) {
	var sut pathSpec
	assert.NoError(t, yaml.Unmarshal([]byte(pathSpecYaml), &sut))

	tpl, err := BuildTemplates()
	assert.NoError(t, err)

	assert.NoError(t, processPath(&sut))

	result := strings.Builder{}
	err = tpl.ExecuteTemplate(&result, "reqStructDecl", sut)
	assert.NoError(t , err)

	assert.Equal(t, requestTypeDeclarationExpected, result.String())
}

const requestTypeConstructorExpected = `func MakeCreateResourceReq(c *gin.Context) (res CreateResourceReq, errorList []error) {
            var err error
            _ = err// URI binding
                for _, param := range c.Params {
                    switch param.Key {

                    default:
                        errorList = append(errorList, fmt.Errorf("Unexpected uri parameter: ` + "`%s`" + `", param.Key))
                    }
                }
                        res.FromDate, err = strconv.ParseInt(c.Query("from_date"), 10, 64)
                        if err != nil {
                            errorList = append(errorList, errors.Wrap(err, "Expecting int64 in from_date"))
                        }
                res.XAccessToken = c.Request.Header.Get("X-access-token")
                if c.Request == nil || c.Request.Body == nil {
                    errorList = append(errorList, fmt.Errorf("Body is absent"))
                }
                switch ct := c.Request.Header.Get("Content-Type"); ct {
                case "application/json":
                    dec := json.NewDecoder(c.Request.Body)
                    if err := dec.Decode(&res.Body); err != nil {
                        errorList = append(errorList, errors.Wrap(err, "Can't Parse JSON"))
                    }
                default:
                    errorList = append(errorList, fmt.Errorf("Unsupported content-type: %v", ct))
                }
            return
        }`

func TestRequestTypeConstructor(t *testing.T) {
	var sut pathSpec
	assert.NoError(t, yaml.Unmarshal([]byte(pathSpecYaml), &sut))

	tpl, err := BuildTemplates()
	assert.NoError(t, err)

	assert.NoError(t, processPath(&sut))

	result := strings.Builder{}
	err = tpl.ExecuteTemplate(&result, "constructorDecl", sut)
	assert.NoError(t , err)

	assert.Equal(t, requestTypeConstructorExpected, result.String())
}

const requestTypeConstructorWithoutReqBodyExpected = `func MakeCreateResourceReq(c *gin.Context) (res CreateResourceReq, errorList []error) {
            var err error
            _ = err// URI binding
                for _, param := range c.Params {
                    switch param.Key {

                    default:
                        errorList = append(errorList, fmt.Errorf("Unexpected uri parameter: ` + "`%s`" + `", param.Key))
                    }
                }
                if c.Request == nil || c.Request.Body != http.NoBody {
                    errorList = append(errorList, fmt.Errorf("Unexpected body"))
                }
            return
        }`

func TestRequestTypeConstructorWithoutRequestBody(t *testing.T) {
	var sut pathSpec
	assert.NoError(t, yaml.Unmarshal([]byte(pathSpecWithoutRequestBodyYaml), &sut))

	tpl, err := BuildTemplates()
	assert.NoError(t, err)

	assert.NoError(t, processPath(&sut))

	result := strings.Builder{}
	err = tpl.ExecuteTemplate(&result, "constructorDecl", sut)
	assert.NoError(t , err)

	assert.Equal(t, requestTypeConstructorWithoutReqBodyExpected, result.String())
}
