package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kepkin/gorest"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/kepkin/gorest/test/api"
)

//go:generate go run generate_api.go

func TestForDebug(t *testing.T) {
	t.Skip()
	err := gorest.Generate("swagger.yaml", gorest.Options{
		PackageName: "api",
		TargetFile:  "api/api_gorest.go",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func TestProvidePayment(t *testing.T) {
	r := gin.New()
	api.RegisterRoutes(r, api.NewPaymentGatewayAPI())

	request := httptest.NewRequest(http.MethodPost, "/v1/payment", bytes.NewReader([]byte(`
		{
			"payment_id": "c63dd8e4-ea77-11e9-b19a-5a001d190301",
			"merchant_id": "cdb39a14-ea77-11e9-a6a4-5a001d190301",
			"sum": "1000.50",
			"type": "deposit"
		}
	`)))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	if !assert.Equal(t, http.StatusOK, response.Code, response.Body.String()) {
		return
	}
	assert.Equal(t, response.Body.String(), `{"provided_total":1000}`)
}

func TestCreateUser(t *testing.T) {
	r := gin.New()
	api.RegisterRoutes(r, api.NewPaymentGatewayAPI())

	// Create user

	var form bytes.Buffer
	wr := multipart.NewWriter(&form)

	photoName := "main-photo.jpeg"
	photoBytes := []byte{1, 2, 3}
	{
		assert.Nil(t, wr.WriteField("name", "anthony"))
		assert.Nil(t, wr.WriteField("email", "test@test.ru"))
		assert.Nil(t, wr.WriteField("age", "31"))

		f, err := wr.CreateFormFile("avatar", photoName)
		assert.Nil(t, err)
		_, err = f.Write(photoBytes)
		assert.Nil(t, err)
	}
	assert.Nil(t, wr.Close())

	request := httptest.NewRequest(http.MethodPost, "/v1/user", &form)
	request.Header.Set("Content-Type", wr.FormDataContentType())

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)
	if !assert.Equal(t, http.StatusOK, response.Code, response.Body.String()) {
		return
	}

	responseData := struct{ ID string }{}
	assert.Nil(t, json.Unmarshal(response.Body.Bytes(), &responseData))
	assert.NotEqual(t, responseData.ID, "")

	// Get user data

	request = httptest.NewRequest(http.MethodGet, "/v1/user/"+responseData.ID, nil)
	response = httptest.NewRecorder()

	r.ServeHTTP(response, request)
	if !assert.Equal(t, http.StatusOK, response.Code, response.Body.String()) {
		return
	}

	assert.Equal(t, response.Body.String(),
		fmt.Sprintf(`{"id":"%s","name":"anthony","email":"test@test.ru","age":31}`, responseData.ID))

	// Get user avatar

	request = httptest.NewRequest(http.MethodGet, "/v1/files/"+photoName, nil)
	response = httptest.NewRecorder()

	r.ServeHTTP(response, request)
	if !assert.Equal(t, http.StatusOK, response.Code, response.Body.String()) {
		return
	}

	assert.Equal(t, response.Body.Bytes(), photoBytes)
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
