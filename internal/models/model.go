package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"go-admin/internal/configs"
	"net/url"
	"time"
)

var (
	db *gorm.DB
)

// Init db
func SetUp() {
	dsn := getDsn()
	var err error

	// connect db
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   configs.GetDbConfig().GetDbPrefix(),
			SingularTable: false,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
	if configs.GetAppConfig().GetDebug() {
		db = db.Debug()
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Set db pool config
	sqlDb.SetConnMaxLifetime(time.Hour)
	sqlDb.SetMaxIdleConns(configs.GetDbConfig().GetDbMaxIdleConns())
	sqlDb.SetMaxOpenConns(configs.GetDbConfig().GetDbMaxOpenConns())
}

// GetDB Get Gorm Db
func GetDB() *gorm.DB {
	return db
}

func getDsn() string {
	var dsn string

	switch configs.GetDbConfig().GetDbEngine() {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=%s",
			configs.GetDbConfig().GetDbUser(),
			configs.GetDbConfig().GetDbPassword(),
			configs.GetDbConfig().GetDbHost(),
			configs.GetDbConfig().GetDbPort(),
			configs.GetDbConfig().GetDbDatabase(),
			configs.GetDbConfig().GetDbCharset(),
			url.QueryEscape(configs.GetDbConfig().GetLoc()),
		)
	default:
		panic("不支持其他数据库")
	}
	return dsn
}
