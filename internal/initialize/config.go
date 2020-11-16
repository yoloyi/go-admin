package initialize

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"monitor/internal/configs"
	"os"
)

var (
	configPath = flag.String("config", "./config.yml", "配置文件路径")
)

// 初始化配置文件
func initConfigYml() {
	config := new(configs.Config)
	// 判断配置文件是否存在
	if _, err := os.Stat(*configPath); err != nil {
		panic(err)
	}
	fileStream, err := ioutil.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(fileStream, config); err != nil {
		panic(err)
	}
	configs.SetConfig(config)
}
