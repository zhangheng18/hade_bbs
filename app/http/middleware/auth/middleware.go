package auth

import (
	"bbs/app/provider/user"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		envService := c.MustMake(contract.EnvKey).(contract.Env)
		userService := c.MustMake(user.UserKey).(user.Service)

		if envService.AppEnv() == contract.EnvDevelopment {
			userID, exist := c.DefaultQueryInt64("user_id", 0)
			if exist {
				authUser, _ := userService.GetUser(c, userID)
				if authUser != nil {
					c.Set("auth_user", authUser)
					c.Next()
					return
				}
			}
		}

		token, err := c.Cookie("hade_bbs")
		if err != nil || token == "" {
			c.ISetStatus(http.StatusUnauthorized).IText("请登录后操作")
			return
		}

		authUser, err := userService.VerifyLogin(c, token)
		if err != nil || authUser == nil {
			c.ISetStatus(http.StatusUnauthorized).IText("请登录后操作")
			return
		}
		c.Set("auth_user", authUser)

		c.Next()
	}
}

// GetAuthUser 获取已经验证的用户
func GetAuthUser(c *gin.Context) *user.User {
	t, exist := c.Get("auth_user")
	if !exist {
		return nil
	}
	return t.(*user.User)
}
