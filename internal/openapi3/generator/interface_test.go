package generator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/openapi3/barber"
	"github.com/kepkin/gorest/internal/openapi3/spec"
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

	sp, err := spec.Read([]byte(swaggerExample))
	if !assert.NoError(t, err) {
		return
	}

	b := new(strings.Builder)
	if !assert.NoError(t, MakeInterface(b, sp)) {
		return
	}

	result, err := barber.PrettifySource("package api\n" + b.String())
	if !assert.NoError(t, err) {
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
}

type TestAPIServer struct {
	Srv TestAPI
}
`, result)
}
