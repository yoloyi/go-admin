package middlewares

import (
	"github.com/gin-gonic/gin"
)

// 处理跨域请求,支持options访问
func Trim() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 处理请求
		c.Next()
	}
}
