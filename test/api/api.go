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

func (p PaymentGatewayAPIImpl) ProvidePayment(in ProvidePaymentReq, c *gin.Context) {
	p.ProvidedSumTotal += in.Body.Sum.IntPart()
	c.AbortWithStatus(http.StatusOK)
}
