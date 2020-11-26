package main

import (
	"monitor/internal/initialize"
	"monitor/internal/routers"
)

func init() {
	initialize.InitApp()
}

func main() {
	routers.RunHttp()
}
