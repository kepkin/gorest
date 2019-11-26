package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var _ PaymentGatewayAPI = (*PaymentGatewayAPIImpl)(nil)

type PaymentGatewayAPIImpl struct {
	ProvidedSumTotal int64
}

func (p PaymentGatewayAPIImpl) ProvidePayment(in ProvidePaymentRequest, c *gin.Context) {
	p.ProvidedSumTotal += in.Body.JSON.Sum.IntPart()
	c.JSON(http.StatusOK, struct {
		ProvidedTotal int64 `json:"provided_total"`
	}{
		ProvidedTotal: p.ProvidedSumTotal,
	})
}

func (p *PaymentGatewayAPIImpl) Example(in ExampleRequest, c *gin.Context) {
	panic("implement me")
}

func (p *PaymentGatewayAPIImpl) ProcessMakeRequestErrors(c *gin.Context, errors []FieldError) {
	c.JSON(http.StatusBadRequest, fmt.Sprintf("parse request error: %v", errors))
}

func (p *PaymentGatewayAPIImpl) ProcessValidateErrors(c *gin.Context, errors []FieldError) {
	c.JSON(http.StatusBadRequest, fmt.Sprintf("validate request error: %v", errors))
}

func (p *PaymentGatewayAPIImpl) GetPayment(in GetPaymentRequest, c *gin.Context) {
	panic("implement me")
}
