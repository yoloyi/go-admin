package response

import (
	"github.com/gin-gonic/gin"
	"go-admin/internal/utils/response/e"
)

type Response struct {
	c *gin.Context
}

func NewResponse(c *gin.Context) *Response {
	return &Response{c: c}
}

func (r *Response) Success(httpCode int, data interface{}) {
	var apiResponse = OkResponse(data)
	r.c.JSON(httpCode, apiResponse.ToH())
	return
}

func (r *Response) Error(httpCode, error int, data interface{}) {
	var apiResponse = ErrorResponse(error, data, e.CodeToMessage(error))
	r.c.JSON(httpCode, apiResponse.ToH())
	return
}

type ApiResponse struct {
	Error   int
	Data    interface{}
	Message string
}

func (requestError *ApiResponse) ToH() gin.H {
	return gin.H{
		"error":   requestError.Error,
		"data":    requestError.Data,
		"message": requestError.Message,
	}
}

func NewApiResponse(error int, data interface{}, message string) *ApiResponse {
	return &ApiResponse{
		Error:   error,
		Data:    data,
		Message: message,
	}
}

func OkResponse(data interface{}) *ApiResponse {
	return NewApiResponse(e.SuccessErrorCode, data, e.CodeToMessage(e.SuccessErrorCode))
}

func ErrorResponse(error int, data interface{}, msg string) *ApiResponse {
	return NewApiResponse(error, data, msg)
}
