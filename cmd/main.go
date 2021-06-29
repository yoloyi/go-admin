package main

import (
	"go-admin/internal/initialize"
	"go-admin/internal/server"
)

func main() {
	initialize.InitApp()
	
	server.RunHttpServer()
}
