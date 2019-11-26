package generator

import (
	"strings"
	"testing"

	"github.com/kepkin/gorest/internal/spec/openapi3"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/barber"
)

func TestMakeInterface(t *testing.T) {
	const swaggerExample = `---
info:
  title: Test API

paths:
  /api/v1/resource:
    get:  
      operationId: GetRecourse

    post:
      operationId: CreateResource

  /api/v2/user:
    options:
       operationId: CheckUser
    
    patch:
       operationId: UpdateUser

    delete:
       operationId: DeleteUser

  /api/v2/admin/anthony:
    put:
       operationId: CreateAdmin
`

	sp, err := openapi3.ReadSpec([]byte(swaggerExample))
	if !assert.NoError(t, err) {
		return
	}

	b := new(strings.Builder)
	if !assert.NoError(t, NewGenerator("api").makeInterface(b, sp)) {
		return
	}
	result := strings.NewReader("package api\n" + b.String())

	prettyResult := new(strings.Builder)
	if !assert.NoError(t, barber.PrettifySource(result, prettyResult)) {
		return
	}

	assert.Equal(t, `package api

type TestAPI interface {
	// GET /api/v1/resource
	GetRecourse(in GetRecourseRequest, c *gin.Context)
	// POST /api/v1/resource
	CreateResource(in CreateResourceRequest, c *gin.Context)

	// PUT /api/v2/admin/anthony
	CreateAdmin(in CreateAdminRequest, c *gin.Context)

	// DELETE /api/v2/user
	DeleteUser(in DeleteUserRequest, c *gin.Context)
	// OPTIONS /api/v2/user
	CheckUser(in CheckUserRequest, c *gin.Context)
	// PATCH /api/v2/user
	UpdateUser(in UpdateUserRequest, c *gin.Context)

	// Service methods
	ProcessMakeRequestErrors(c *gin.Context, errors []FieldError)
	ProcessValidateErrors(c *gin.Context, errors []FieldError)
}

type TestAPIServer struct {
	Srv TestAPI
}
`, prettyResult.String())
}
