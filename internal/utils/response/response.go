package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	NoErrorCode = 0
)

type ApiResponse struct {
	Code    uint
	Data    interface{}
	Message string
}

func (requestError *ApiResponse) Error() string {
	return fmt.Sprintf("Code：%d, Error message: %s", requestError.Code, requestError.Message)
}

func (requestError *ApiResponse) ToH() gin.H {
	return gin.H{
		"code":    requestError.Code,
		"data":    requestError.Data,
		"message": requestError.Message,
	}
}

func NewRequestError(code uint, data interface{}, message string) *ApiResponse {
	return &ApiResponse{
		Code:    code,
		Data:    data,
		Message: message,
	}
}

func Ok(data interface{}, msg string) *ApiResponse {
	return NewRequestError(NoErrorCode, data, msg)
}

func Error(code uint, data interface{}, msg string) *ApiResponse {
	return NewRequestError(code, data, msg)
}

func OkH(data interface{}, msg string) gin.H {
	return Ok(data, msg).ToH()
}

func ErrorH(code uint, data interface{}, msg string) gin.H {
	return NewRequestError(code, data, msg).ToH()
}
