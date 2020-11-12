package configs

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	once       sync.Once
	config     = new(Config)
	configPath = flag.String("config", "./config.yml", "配置文件路径")
)

type Config struct {
	DbConfig  *DbConfig  `yaml:"db"`
	AppConfig *AppConfig `yaml:"app"`
}

// GetDbConfig 获取配置文件
func GetDbConfig() IDbConfig {
	return config.DbConfig
}

// GetAppConfig 获取配置文件
func GetAppConfig() IAppConfig {
	return config.AppConfig
}

// InitConfig 初始化配置
func InitConfig() {
	flag.Parse()
	once.Do(func() {
		loadConfigYml()
	})
}

func loadConfigYml() {
	// 判断配置文件是否存在
	if _, err := os.Stat(*configPath); err != nil {
		log.Fatalln(err)
	}
	fileStream, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatalln(err)
	}
	if err := yaml.Unmarshal(fileStream, config); err != nil {
		log.Fatalln(err)
	}
}
