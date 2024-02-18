package user

import (
	provider "bbs/app/provider/user"
	"github.com/gohade/hade/framework/gin"
	"net/http"
)

type loginParam struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,gte=6"`
}

// Login 登录
// @Summary 用户登录
// @Description 用户登录接口
// @Accept  json
// @Produce  json
// @Tags user
// @Param loginParam body loginParam  true "login with param"
// @Success 200 {string} Token "token"
// @Router /api/user/login [post]
func (api *UserApi) Login(c *gin.Context) {
	userService := c.MustMake(provider.UserKey).(provider.Service)

	//验证参数
	params := &loginParam{}
	if err := c.ShouldBind(params); err != nil {
		c.ISetStatus(http.StatusBadRequest).IText("参数错误，必须包含 username 和 password（大于6位）")
		return
	}

	//登录
	model := &provider.User{
		UserName: params.UserName,
		Password: params.Password,
	}
	userWithToken, err := userService.Login(c, model)
	if err != nil {
		c.ISetStatus(http.StatusInternalServerError).IText(err.Error())
		return
	}
	if userWithToken == nil {
		c.ISetStatus(http.StatusUnauthorized).IText("用户不存在")
		return
	}

	c.ISetOkStatus().IText(userWithToken.Token)
	return
}
