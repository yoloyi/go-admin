package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type basic struct{}

// Healthz 健康检查
func (h basic) Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func RegisterBasicRouter(router *gin.RouterGroup) {
	var h basic
	r := router.Group("")
	{
		r.Any("healthz", h.Healthz)
	}
}
