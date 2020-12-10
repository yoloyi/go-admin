package auth

type LoginRequest struct {
	Username string `json:"username" validate:"required,ip" label:"账号"`
	Password string `json:"password" validate:"required" label:"密码"`
}
