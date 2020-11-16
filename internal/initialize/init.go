package initialize

import (
	"flag"
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
	})
}
