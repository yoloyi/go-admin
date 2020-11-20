package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primary_key;column:id;type:int(10) unsigned;not null"`           // 主键
	CreatedAt time.Time      `gorm:"column:created_time;type:timestamp"`                             // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_time;type:timestamp"`                             // 修改时间
	DeletedAt gorm.DeletedAt `gorm:"index:idx_deleted_time;column:deleted_time;type:timestamp;null"` // 删除时间
}
