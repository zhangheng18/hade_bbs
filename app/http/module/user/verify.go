package user

import (
	provider "bbs/app/provider/user"
	"github.com/gohade/hade/framework/gin"
	"net/http"
)

// Verify 代表验证注册信息
// @Summary 验证注册信息
// @Description 使用token验证用户注册信息
// @Accept  json
// @Produce  json
// @Tags user
// @Param token query string true "注册token"
// @Success 200 {string} Message "注册成功，请进入登录页面"
// @Router /api/user/register/verify [get]
func (api *UserApi) Verify(c *gin.Context) {
	// 验证参数
	userService := c.MustMake(provider.UserKey).(provider.Service)
	token := c.Query("token")
	if token == "" {
		c.ISetStatus(http.StatusNotFound).IText("参数错误")
		return
	}

	verified, err := userService.VerifyRegister(c, token)
	if err != nil {
		c.ISetStatus(http.StatusInternalServerError).IText(err.Error())
		return
	}

	if !verified {
		c.ISetStatus(http.StatusInternalServerError).IText("验证错误")
		return
	}

	// 输出
	c.IRedirect("/login").IText("注册成功，请进入登录页面")
}
