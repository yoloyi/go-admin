package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	log "github.com/sirupsen/logrus"
	"go-admin/internal/models/entities"
	"go-admin/internal/models/repositories"
	"go-admin/internal/routers/requests"
	"go-admin/internal/routers/requests/user"
	"go-admin/internal/utils/response"
	"go-admin/internal/utils/response/e"
)

var ServiceWireSet = wire.NewSet(NewUser)

type Service struct {
	userRepository repositories.User
}

func NewUser(userRepository repositories.User) Service {
	return Service{
		userRepository: userRepository,
	}
}

// UserChangePassword 用户修改密码
func (this Service) UserChangePassword(ctx *gin.Context) {
	var changePasswordRequest user.ChangePasswordRequest
	ctx.BindJSON(&changePasswordRequest)

	var r = response.NewResponse(ctx)
	if !requests.Validate(ctx, changePasswordRequest) {
		return
	}
	userInfo := ctx.MustGet("user").(*entities.User)
	if !userInfo.CheckPasswordValid(changePasswordRequest.OldPassword) {
		r.HttpOkError(e.UserChangePasswordNotMatch, nil)
		return
	}

	// 生成新的密码
	if err := userInfo.GenerateNewPassword(changePasswordRequest.Password); err != nil {
		log.Error(err)
		r.HttpOkError(e.FaultErrorCode, nil)
		return
	}

	err := this.userRepository.DB.Select("password").Updates(userInfo).Error
	if err != nil {
		r.HttpOkError(e.FaultErrorCode, nil)
		return
	}
	r.HttpOkSuccess(nil)
	return
}
