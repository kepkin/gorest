package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var _ PaymentGatewayAPI = (*PaymentGatewayAPIImpl)(nil)

type PaymentGatewayAPIImpl struct {
	ProvidedSumTotal int64
}

func (p PaymentGatewayAPIImpl) ProvidePayment(in ProvidePaymentRequest, c *gin.Context) {
	p.ProvidedSumTotal += in.Body.JSON.Sum.IntPart()
	c.AbortWithStatus(http.StatusOK)
}

func (p *PaymentGatewayAPIImpl) Example(in ExampleRequest, c *gin.Context) {
	panic("implement me")
}

func (p *PaymentGatewayAPIImpl) ProcessMakeRequestErrors(c *gin.Context, errors []FieldError) {
	panic("implement me")
}

func (p *PaymentGatewayAPIImpl) ProcessValidateErrors(c *gin.Context, errors []FieldError) {
	panic("implement me")
}

func (p *PaymentGatewayAPIImpl) GetPayment(in GetPaymentRequest, c *gin.Context) {
	panic("implement me")
}
