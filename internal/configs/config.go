package configs

var (
	config = new(Config)
)

type Config struct {
	DbConfig  DbConfig  `yaml:"db" json:"db"`
	AppConfig AppConfig `yaml:"app" json:"app"`
	LogConfig LogConfig `yaml:"log" json:"log"`
}

// GetDbConfig 获取配置
func GetDbConfig() IDbConfig {
	return config.DbConfig
}

// GetAppConfig 获取配置
func GetAppConfig() IAppConfig {
	return config.AppConfig
}

// GetLogConfig 获取日志配置
func GetLogConfig() ILogConfig {
	return config.LogConfig
}

func SetConfig(c *Config) {
	config = c
}
