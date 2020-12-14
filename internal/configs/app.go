package configs

type IAppConfig interface {
	GetDebug() bool
	GetEnv() string
	GetLogPath() string
	GetName() string
	GetPort() int16
	GetKey() string
}

type AppConfig struct {
	Debug   bool   `yaml:"debug"`
	Env     string `yaml:"env"`
	LogPath string `yaml:"log_path"`
	Name    string `yaml:"name"`
	Port    int16  `yaml:"port"`
	Key     string `yaml:"key"`
}

func (a *AppConfig) GetDebug() bool {
	return a.Debug
}

func (a *AppConfig) GetEnv() string {
	return a.Env
}

func (a *AppConfig) GetLogPath() string {
	return a.LogPath
}
func (a *AppConfig) GetName() string {
	return a.Name
}

func (a *AppConfig) GetPort() int16 {
	return a.Port
}

func (a *AppConfig) GetKey() string {
	return a.Key
}
