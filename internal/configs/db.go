package configs

type IDbConfig interface {
	GetDbEngine() string
	GetDbHost() string
	GetDbPort() int
	GetDbUser() string
	GetDbPassword() string
	GetDbDatabase() string
	GetDbPrefix() string
	GetDbCharset() string
	GetDbMaxIdleConns() int
	GetDbMaxOpenConns() int
	GetLoc() string
}

type DbConfig struct {
	Engine       string `yaml: "engine"`
	Host         string `yaml: "host"`
	Port         int    `yaml: "port"`
	User         string `yaml: "user"`
	Password     string `yaml: "password"`
	Database     string `yaml: "database"`
	Prefix       string `yaml: "prefix"`
	Charset      string `yaml: "charset"`
	MaxIdleConns int    `yaml: "maxidleconns"`
	MaxOpenConns int    `yaml: "maxopenconns"`
	Loc          string `yaml: "loc"`
}

func (d *DbConfig) GetDbEngine() string {
	return d.Engine
}

func (d *DbConfig) GetDbHost() string {
	return d.Host
}

func (d *DbConfig) GetDbPort() int {
	return d.Port
}

func (d *DbConfig) GetDbUser() string {
	return d.User
}

func (d *DbConfig) GetDbPassword() string {
	return d.Password
}

func (d *DbConfig) GetDbDatabase() string {
	return d.Database
}

func (d *DbConfig) GetDbPrefix() string {
	return d.Prefix
}

func (d *DbConfig) GetDbCharset() string {
	return d.Charset
}

func (d *DbConfig) GetDbMaxIdleConns() int {
	return d.MaxIdleConns
}

func (d *DbConfig) GetDbMaxOpenConns() int {
	return d.MaxOpenConns
}

func (d *DbConfig) GetLoc() string {
	return d.Loc
}
