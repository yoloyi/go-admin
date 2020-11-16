package configs

type IDbConfig interface {
	GetDbEngine() string
	getDbHost() string
	getDbPort() int
	getDbUser() string
	getDbPassword() string
	getDbDatabase() string
	getDbPrefix() string
	getDbCharset() string
	getDbMaxIdleConns() int
	getDbMaxOpenConns() int
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
}

func (d *DbConfig) GetDbEngine() string {
	return d.Engine
}

func (d *DbConfig) getDbHost() string {
	return d.Host
}

func (d *DbConfig) getDbPort() int {
	return d.Port
}

func (d *DbConfig) getDbUser() string {
	return d.User
}

func (d *DbConfig) getDbPassword() string {
	return d.Password
}

func (d *DbConfig) getDbDatabase() string {
	return d.Database
}

func (d *DbConfig) getDbPrefix() string {
	return d.Prefix
}

func (d *DbConfig) getDbCharset() string {
	return d.Charset
}

func (d *DbConfig) getDbMaxIdleConns() int {
	return d.MaxIdleConns
}

func (d *DbConfig) getDbMaxOpenConns() int {
	return d.MaxOpenConns
}
