package models

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"monitor/internal/configs"
	"net/url"
	"time"
)

var db *gorm.DB

type BaseModel struct {
	ID        uint           `gorm:"primary_key;column:id;type:int(10) unsigned;not null"`        // 主键
	CreatedAt time.Time      `gorm:"column:created_time;type:datetime"`                             // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_time;type:datetime"`                             // 修改时间
	DeletedAt gorm.DeletedAt `gorm:"index:idx_deleted_time;column:deleted_time;type:datetime;null"` // 删除时间
}

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
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	// Set db pool config
	sqlDb.SetConnMaxLifetime(time.Hour)
	sqlDb.SetMaxIdleConns(configs.GetDbConfig().GetDbMaxIdleConns())
	sqlDb.SetMaxOpenConns(configs.GetDbConfig().GetDbMaxOpenConns())
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
