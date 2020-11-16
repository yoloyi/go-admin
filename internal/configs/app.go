package configs

type IAppConfig interface {
	GetDebug() bool
	GetEnv() string
	GetLogPath() string
	GetName() string
}

type AppConfig struct {
	Debug   bool   `yaml:"debug"`
	Env     string `yaml:"env"`
	LogPath string `yaml:"log_path"`
	Name    string `yaml:"name"`
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
