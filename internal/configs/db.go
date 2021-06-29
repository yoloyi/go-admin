package configs

type IDbConfig interface {
	GetDbPort() int
	GetDbMaxIdleConns() int
	GetDbMaxOpenConns() int
	GetDbEngine() string
	GetDbHost() string
	GetDbUser() string
	GetDbPassword() string
	GetDbDatabase() string
	GetDbPrefix() string
	GetExtra() string
}

type DbConfig struct {
	Port         int    `yaml:"port"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	Engine       string `yaml:"engine"`
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Database     string `yaml:"database"`
	TablePrefix  string `yaml:"table_prefix"`
	Extra        string `yaml:"extra"`
}

func (d DbConfig) GetExtra() string {
	return d.Extra
}

func (d DbConfig) GetDbEngine() string {
	return d.Engine
}

func (d DbConfig) GetDbHost() string {
	return d.Host
}

func (d DbConfig) GetDbPort() int {
	return d.Port
}

func (d DbConfig) GetDbUser() string {
	return d.User
}

func (d DbConfig) GetDbPassword() string {
	return d.Password
}

func (d DbConfig) GetDbDatabase() string {
	return d.Database
}

func (d DbConfig) GetDbPrefix() string {
	return d.TablePrefix
}

func (d DbConfig) GetDbMaxIdleConns() int {
	return d.MaxIdleConns
}

func (d DbConfig) GetDbMaxOpenConns() int {
	return d.MaxOpenConns
}
