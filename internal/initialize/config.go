package initialize

import (
	"flag"
	"go-admin/internal/configs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var (
	configPath = flag.String("c", "./config.yml", "配置文件路径")
)

// 初始化配置文件
func initConfigYml() {
	config := new(configs.Config)
	// 判断配置文件是否存在
	if _, err := os.Stat(*configPath); err != nil {
		log.Fatal(err)
	}
	fileStream, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(fileStream, config); err != nil {
		log.Fatal(err)
	}
	configs.SetConfig(config)
}
