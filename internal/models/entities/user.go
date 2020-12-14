package entities

// Users 后台用户表
type User struct {
	Username string `gorm:"unique;index:udx_username;column:username;type:varchar(100);not null"` // 用户名
	Password string `gorm:"column:password;type:varchar(100);not null"`                           // 密码
	Status   int8   `gorm:"column:status;type:tinyint(1);not null"`                               // 状态
	BaseEntity
}

const (
	UserStatusNormal = iota  // 正常用户
	UserForbidden    = iota  // 禁用
)
