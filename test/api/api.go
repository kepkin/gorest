package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type Decimal = decimal.Decimal

var _ PaymentGatewayAPI = (*PaymentGatewayAPIImpl)(nil)

type PaymentGatewayAPIImpl struct {
	ProvidedSumTotal int64
}

func (p *PaymentGatewayAPIImpl) GetPayment(in GetPaymentRequest, c *gin.Context) {
	panic("implement me")
}

func (p PaymentGatewayAPIImpl) ProvidePayment(in ProvidePaymentRequest, c *gin.Context) {
	p.ProvidedSumTotal += in.JSONBody.Sum.IntPart()
	c.AbortWithStatus(http.StatusOK)
}

func (p *PaymentGatewayAPIImpl) Example(in ExampleRequest, c *gin.Context) {
	panic("implement me")
}
