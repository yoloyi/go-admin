package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	log "github.com/sirupsen/logrus"
	"go-admin/internal/configs"
	"go-admin/internal/models/entities"
	"go-admin/internal/models/repositories"
	"go-admin/internal/routers/requests"
	auth2 "go-admin/internal/routers/requests/auth"
	"go-admin/internal/routers/response/auth"
	authUtil "go-admin/internal/utils/auth"
	"go-admin/internal/utils/response"
	"go-admin/internal/utils/response/e"
	"gorm.io/gorm"
	"time"
)

var ServiceWireSet = wire.NewSet(NewAuth, authUtil.NewAuth)

type Service struct {
	userRepository repositories.User
	authUtil       authUtil.Auth
}

func NewAuth(userRepository repositories.User, authUtil authUtil.Auth) Service {
	return Service{
		userRepository: userRepository,
		authUtil:       authUtil,
	}
}

func (this Service) LoginService(ctx *gin.Context) {
	var loginRequest auth2.LoginRequest
	ctx.BindJSON(&loginRequest)

	var r = response.NewResponse(ctx)
	if !requests.Validate(ctx, loginRequest) {
		return
	}

	user, err := this.userRepository.GetUserByUserName(loginRequest.Username)
	if err != nil {
		// 用户不存在
		if err == gorm.ErrRecordNotFound {
			r.HttpOkError(e.UserNameOrPasswordNotMatch, nil)
			return
		}
		log.Error("登录查询用户错误", err)
		r.HttpOkError(e.FaultErrorCode, nil)
		return
	}
	// 验证密码是否正确
	if !this.authUtil.PasswordVerify(loginRequest.Password, user.Password) {
		r.HttpOkError(e.UserNameOrPasswordNotMatch, nil)
		return
	}

	if user.Status != entities.UserStatusNormal {
		r.HttpOkError(e.UsesAbnormal, nil)
		return
	}

	var jwtUtil = authUtil.NewJwt(configs.GetAppConfig().GetKey())
	jwtToken, err := jwtUtil.GenerateJwtToken(user.ID, 12*time.Hour)
	if err != nil {
		log.Errorf("登录生成 Token 异常", err)
		r.HttpOkError(e.FaultErrorCode, nil)
		return
	}

	r.HttpOkSuccess(auth.LoginResponse{Token: jwtToken})
}
