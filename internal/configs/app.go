package configs

type IAppConfig interface {
	GetDebug() bool
	GetEnv() string
	GetName() string
	GetPort() int16
	GetKey() string
}

type AppConfig struct {
	Debug   bool   `yaml:"debug"`
	Env     string `yaml:"env"`
	Name    string `yaml:"name"`
	Port    int16  `yaml:"port"`
	Key     string `yaml:"key"`
}

func (a AppConfig) GetDebug() bool {
	return a.Debug
}

func (a AppConfig) GetEnv() string {
	return a.Env
}

func (a AppConfig) GetName() string {
	return a.Name
}

func (a AppConfig) GetPort() int16 {
	return a.Port
}

func (a AppConfig) GetKey() string {
	return a.Key
}
