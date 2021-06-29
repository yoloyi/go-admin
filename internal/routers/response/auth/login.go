package auth

type LoginResponse struct {
	Token string `json:"token" label:"秘钥 Token"`
}
