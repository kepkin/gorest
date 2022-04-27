package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kepkin/gorest/test/api"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
		fmt.Sprintf(`{"id":"%s","name":"anthony","email":"test@test.ru","age":31}` + "\n", responseData.ID))

	// Get user avatar

	request = httptest.NewRequest(http.MethodGet, "/v1/files/"+photoName, nil)
	response = httptest.NewRecorder()

	r.ServeHTTP(response, request)
	if !assert.Equal(t, http.StatusOK, response.Code, response.Body.String()) {
		return
	}

	assert.Equal(t, response.Body.Bytes(), photoBytes)
}
