package main

import (
	"monitor/configs"
	"monitor/utils/logger"
)

func init() {
	initMonitor()
}

func main() {

}

func initMonitor() {
	configs.InitConfig()
	logger.InitLogger()
}
