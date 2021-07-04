package user

type ChangePasswordRequest struct {
	OldPassword    string `json:"old_password" validate:"required" label:"旧密码"`
	Password       string `json:"password" validate:"required,min=6,max=24" label:"密码"`
	RepeatPassword string `json:"repeat_password" validate:"required,eqfield=Password,min=6,max=24" label:"重复密码"`
}
