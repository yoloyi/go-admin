package entities

// Users 后台用户表
type Users struct {
	UserName string `gorm:"unique;index:udx_user_name;column:user_name;type:varchar(100);not null"` // 用户名
	Password string `gorm:"column:password;type:varchar(100);not null"`                             // 密码
	Status   int8   `gorm:"column:status;type:tinyint(1);not null"`                                 // 状态
	BaseModel
}
