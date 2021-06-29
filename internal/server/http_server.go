package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-admin/internal/configs"
	"go-admin/internal/routers"
)

type server interface {
	ListenAndServe() error
}

func RunHttpServer() {
	router := routers.InitRouter()
	var address = getAddr()
	s := initServer(address, router)
	log.Info("server run sucess on ", address)
	log.Error(s.ListenAndServe().Error())
}

func getAddr() string {
	return fmt.Sprintf(":%d", configs.GetAppConfig().GetPort())
}
