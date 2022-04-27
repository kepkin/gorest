package test

import (
	"log"
	"testing"

	"github.com/kepkin/gorest"
)

//go:generate go run generate_api.go

func TestForDebug(t *testing.T) {
	//t.Skip()
	err := gorest.Generate("swagger.yaml", gorest.Options{
		PackageName: "api",
		TargetFile:  "api/api_gorest.go",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func TestQueryParams(t *testing.T) {
	//t.Skip()
	err := gorest.Generate("swaggers/query_params.yaml", gorest.Options{
		PackageName: "apis/",
		TargetFile:  "apis/query_params/gorest.go",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func TestJSONBodyParams(t *testing.T) {
	//t.Skip()
	err := gorest.Generate("swaggers/json_body_params.yaml", gorest.Options{
		PackageName: "apis/",
		TargetFile:  "apis/json_body_params/gorest.go",
	})
	if err != nil {
		log.Fatal(err)
	}
}


func TestCustomRouteWithMiddleWare(t *testing.T) {
	r := gin.New()
	api.RegisterRoutesCustom(func(operationID, httpMethod, relativePath string, handler gin.HandlerFunc) {
		handlers := make([]gin.HandlerFunc, 0)

		if operationID == "GetUser" {
			handlers = append(handlers, func(c *gin.Context){
				c.AbortWithStatus(401)
			})
		}

		handlers = append(handlers, handler)
		r.Handle(httpMethod, relativePath, handlers...)
	}, api.NewPaymentGatewayAPI())

	// Go to route with middle ware
	request := httptest.NewRequest(http.MethodGet, "/v1/user/123", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	if !assert.Equal(t, http.StatusUnauthorized, response.Code, response.Body.String()) {
		return
	}
}
