package generator

const Predefined = `// Code generated by gorest; DO NOT EDIT.

package %s

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const handlerNameKey = "handler"

type ParamPlace int

const (
	UndefinedPlace ParamPlace = iota
	InBody
	InCookie
	InFormData
	InHeader
	InPath
	InQuery
)

type ContentType int

const (
	UndefinedContentType ContentType = iota
	AppJSON
	AppXML
	AppFormUrlencoded
	MultipartFormData
	TextPlain
)

type FieldError struct {
	In      ParamPlace
	Field   string
	Message string
	Reason  error
}

func NewFieldError(in ParamPlace, f string, msg string, err error) FieldError {
	return FieldError{
		In:      in,
		Field:   f,
		Message: msg,
		Reason:  err,
	}
}
`
