package main

import (
	"go-admin/internal/initialize"
	"go-admin/internal/routers"
)

func init() {
	initialize.InitApp()
}

func main() {
	routers.RunHttp()

}
