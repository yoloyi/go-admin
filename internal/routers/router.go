package routers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-admin/internal/configs"
	"go-admin/internal/routers/handlers"
	middleware "go-admin/internal/routers/middlewares"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func initRouter() *gin.Engine {

	// 当 env 不为 release 则统一为 debug 环境
	if configs.GetAppConfig().GetEnv() != gin.ReleaseMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	var router = gin.Default()
	log.Printf("【App】use middleware cors\n")
	router.Use(middleware.Cors())
	log.Printf("Init middleware cors success\n")
	registerRouter(router)
	return router

}

// registerRouter register router handler
func registerRouter(router *gin.Engine) {
	routerGroup := router.Group("")

	handlers.RegisterAuthRouter(routerGroup)
	handlers.RegisterUserRouter(routerGroup)
}

func RunHttp() {
	router := initRouter()

	srv := &http.Server{
		Addr:    getAddr(),
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func getAddr() string {
	return fmt.Sprintf(":%d", configs.GetAppConfig().GetPort())
}
