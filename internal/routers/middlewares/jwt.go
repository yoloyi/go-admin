package middlewares

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-admin/internal/configs"
	"go-admin/internal/models/repositories"
	authUtil "go-admin/internal/utils/auth"
	"go-admin/internal/utils/response"
	"go-admin/internal/utils/response/e"
	"net/http"
	"strings"
	"time"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		r := response.NewResponse(c)
		if token == "" {
			r.Error(http.StatusUnauthorized, e.NotLoginCode, nil)
			c.Abort()
			return
		}

		// 去除 Bearer
		token = strings.TrimPrefix(token, "Bearer ")
		// 解析 Jwt
		var jwtUtil = authUtil.NewJwt(configs.GetAppConfig().GetKey())
		jc, err := jwtUtil.ParseJwtToken(token)

		if err != nil {
			log.Info("解析登录信息失败", err)
			r.Error(http.StatusUnauthorized, e.NotLoginCode, nil)
			c.Abort()
			return
		}

		// 验证JWT 是否真实有效
		now := time.Now().Unix()
		if !jc.VerifyExpiresAt(now, true) {
			r.Error(http.StatusUnauthorized, e.TokenExpiredCode, nil)
			c.Abort()
			return
		}
		var userRepository = repositories.NewUserRepository()
		user, err := userRepository.First(jc.UserId)
		if err != nil {
			log.Info("用户找不到", jc, err)
			r.Error(http.StatusUnauthorized, e.TokenExpiredCode, nil)
			c.Abort()
			return
		}

		if !user.CheckUserNormal() {
			r.Error(http.StatusUnauthorized, e.UsesAbnormal, nil)
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
