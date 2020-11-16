package main

import (
	"monitor/internal/initialize"
	"monitor/internal/models"
)

func init() {
	initialize.InitApp()
}

func main() {
	models.Test()

}
