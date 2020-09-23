package configs

type Config struct {
	Db *DbConfig `yaml:"db"`
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
