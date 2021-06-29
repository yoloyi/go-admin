package configs

type ILogConfig interface {
	GetLogPath() string
	GetLinkName() string
	GetShowLine() bool
	GetLevel() string
}

type LogConfig struct {
	LogPath  string `yaml:"log_path" json:"log_path"`
	LinkName string `yaml:"link_name" json:"link_name"`
	ShowLine bool   `yaml:"show_line" json:"show_line"`
	Level    string ` yaml:"level" json:"level"`
}

func (l LogConfig) GetLogPath() string {
	return l.LogPath
}

func (l LogConfig) GetLinkName() string {
	return l.LinkName
}

func (l LogConfig) GetShowLine() bool {
	return l.ShowLine
}

func (l LogConfig) GetLevel() string {
	return l.Level
}
