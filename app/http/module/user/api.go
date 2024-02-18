package user

import (
	"bbs/app/http/middleware/auth"
	"bbs/app/provider/user"
	"github.com/gohade/hade/framework/gin"
)

type UserApi struct{}

func RegisterRoutes(r *gin.Engine) error {
	api := &UserApi{}

	if !r.IsBind(user.UserKey) {
		r.Bind(&user.UserProvider{})
	}

	// 登录
	r.POST("/api/user/login", api.Login)
	// 登出
	r.GET("/api/user/logout", auth.AuthMiddleware(), api.Logout)
	// 注册
	r.POST("/api/user/register", api.Register)
	// 注册验证
	r.GET("/api/user/register/verify", api.Verify)
	return nil
}
