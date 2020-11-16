package configs

var (
	config = new(Config)
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

func SetConfig(c *Config) {
	config = c
}
