package initialize

import (
	"flag"
	"monitor/internal/models"
	"sync"
)

var (
	once sync.Once
)

func init() {
	flag.Parse()
}

func InitApp() {
	once.Do(func() {
		initConfigYml()
		initLogger()
		models.SetUp()
	})
}
