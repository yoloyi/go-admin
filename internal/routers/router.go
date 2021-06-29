package routers

import (
	"github.com/gin-gonic/gin"
	"go-admin/internal/configs"
	"go-admin/internal/routers/handlers"
	middleware "go-admin/internal/routers/middlewares"
	"log"
)

func InitRouter() *gin.Engine {
	// 当 env 不为 release 则统一为 debug 环境
	if configs.GetAppConfig().GetEnv() != gin.ReleaseMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	var router = gin.Default()
	log.Println("【App】use middleware cors")
	router.Use(middleware.Cors(), middleware.Trim())
	log.Println("Init middleware cors success")
	registerRouter(router)
	return router
}

// registerRouter register router handler
func registerRouter(router *gin.Engine) {
	routerGroup := router.Group("")

	handlers.RegisterAuthRouter(routerGroup)
	handlers.RegisterUserRouter(routerGroup)
}
