package qa

import (
	"bbs/app/http/middleware/auth"
	provider "bbs/app/provider/qa"
	"github.com/gohade/hade/framework/gin"
	"net/http"
)

// AnswerDelete 代表删除回答
// @Summary 创建回答
// @Description 创建回答
// @Accept  json
// @Produce  json
// @Tags qa
// @Param id query int true "删除id"
// @Success 200 string Msg "操作成功"
// @Router /api/answer/delete [post]
func (api *QAApi) AnswerDelete(c *gin.Context) {
	qaService := c.MustMake(provider.QaKey).(provider.Service)
	id, exist := c.DefaultQueryInt64("id", 0)
	if !exist {
		c.ISetStatus(http.StatusBadRequest).IText("参数错误")
		return
	}
	user := auth.GetAuthUser(c)

	answer, err := qaService.GetAnswer(c, id)
	if err != nil {
		c.ISetStatus(http.StatusInternalServerError).IText(err.Error())
		return
	}
	if answer.AuthorID != user.ID {
		c.ISetStatus(http.StatusInternalServerError).IText("没有权限做此操作")
		return
	}

	if err := qaService.DeleteAnswer(c, id); err != nil {
		c.ISetStatus(http.StatusInternalServerError).IText(err.Error())
		return
	}
	c.ISetOkStatus().IText("操作成功")
}
