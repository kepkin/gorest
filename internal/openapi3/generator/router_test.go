package generator

import (
	"testing"

	"github.com/kepkin/gorest/internal/openapi3/spec"
	"github.com/stretchr/testify/assert"
)

func TestMakeRouter(t *testing.T) {
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

	_, err := spec.Read([]byte(swaggerExample))
	if !assert.NoError(t, err) {
		return
	}
	//
	//	b := new(strings.Builder)
	//	if !assert.NoError(t, MakeRouter(b, sp)) {
	//		return
	//	}
	//
	//	result, err := barber.PrettifySource("package api\n" + b.String())
	//	if !assert.NoError(t, err) {
	//		return
	//	}
	//
	//	assert.Equal(t, `package api
	//
	//// Router
	//func RegisterRoutes(r *gin.Engine, api TestAPI) {
	//	e := &TestAPIServer{api}
	//
	//	r.Handle("GET", "/api/v1/resource", e._TestAPI_GetRecourse_Handler)
	//	r.Handle("POST", "/api/v1/resource", e._TestAPI_CreateResource_Handler)
	//
	//	r.Handle("PUT", "/api/v2/admin/anthony", e._TestAPI_CreateAdmin_Handler)
	//
	//	r.Handle("DELETE", "/api/v2/user", e._TestAPI_DeleteUser_Handler)
	//	r.Handle("OPTIONS", "/api/v2/user", e._TestAPI_CheckUser_Handler)
	//	r.Handle("PATCH", "/api/v2/user", e._TestAPI_UpdateUser_Handler)
	//}
	//`, result)
}
