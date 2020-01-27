package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var _ PaymentGatewayAPI = (*PaymentGatewayAPIImpl)(nil)

type PaymentGatewayAPIImpl struct {
	ProvidedSumTotal int64
	Files            map[string][]byte
	Users            map[ID]User
}

func NewPaymentGatewayAPI() PaymentGatewayAPI {
	return &PaymentGatewayAPIImpl{
		Files: make(map[string][]byte),
		Users: make(map[ID]User),
	}
}

// Errors processing

func (p *PaymentGatewayAPIImpl) ProcessMakeRequestErrors(c *gin.Context, errors []FieldError) {
	c.JSON(http.StatusBadRequest, fmt.Sprintf("parse request error: %+v", errors))
}

func (p *PaymentGatewayAPIImpl) ProcessValidateErrors(c *gin.Context, errors []FieldError) {
	c.JSON(http.StatusBadRequest, fmt.Sprintf("validate request error: %+v", errors))
}

// Logic

func (p PaymentGatewayAPIImpl) ProvidePayment(in ProvidePaymentRequest, c *gin.Context) {
	p.ProvidedSumTotal += in.Body.JSON.Sum.IntPart()
	c.JSON(http.StatusOK, struct {
		ProvidedTotal int64 `json:"provided_total"`
	}{
		ProvidedTotal: p.ProvidedSumTotal,
	})
}

func (p *PaymentGatewayAPIImpl) GetFile(in GetFileRequest, c *gin.Context) {
	file, ok := p.Files[in.Path.Filename]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Data(http.StatusOK, "application/octet-stream", file)
}

func (p *PaymentGatewayAPIImpl) CreateUser(in CreateUserRequest, c *gin.Context) {
	id := ID(uuid.New().String())
	p.Users[id] = User{
		ID:        id,
		Name:      in.Body.Form.Name,
		Email:     in.Body.Form.Email,
		Age:       int(in.Body.Form.Age),
		AvatarURL: "",
	}

	avatarHdr := in.Body.Form.Avatar
	avatar, err := avatarHdr.Open()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	buf := bytes.NewBuffer([]byte{})
	if _, err := io.Copy(buf, avatar); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	p.Files[avatarHdr.Filename] = buf.Bytes()

	c.JSON(http.StatusOK, struct{ ID ID }{id})
}

func (p PaymentGatewayAPIImpl) GetUser(in GetUserRequest, c *gin.Context) {
	user, ok := p.Users[ID(in.Path.UserID)]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}

// Not implemented

func (p *PaymentGatewayAPIImpl) Example(in ExampleRequest, c *gin.Context) {
	panic("implement me")
}

func (p *PaymentGatewayAPIImpl) GetPayment(in GetPaymentRequest, c *gin.Context) {
	panic("implement me")
}
