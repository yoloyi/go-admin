package main

import (
	"github.com/sirupsen/logrus"
	"monitor/internal/initialize"
)

func init() {
	initialize.InitApp()
}

func main() {
	logrus.Warn(123)

}
