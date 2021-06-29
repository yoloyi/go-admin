package entities

import (
	"time"
)

type BaseEntity struct {
	ID        uint           `gorm:"primary_key;column:id;type:int(10) unsigned;not null"`           // 主键
	CreatedAt time.Time      `gorm:"column:created_time;type:timestamp"`                             // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_time;type:timestamp"`                             // 修改时间
}
