package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/test/api"
)

//go:generate go run generate_api.go

func TestProvidePayment(t *testing.T) {
	r := gin.New()
	api.RegisterRoutes(r, &api.PaymentGatewayAPIImpl{})

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
