package qa

import (
	"bbs/app/http/middleware/auth"
	provider "bbs/app/provider/qa"
	"github.com/gohade/hade/framework/gin"
	"net/http"
)

type questionEditParam struct {
	ID      int64  `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Context string `json:"context" binding:"required"`
}

// QuestionEdit 编辑问题
// @Summary 编辑问题
// @Description 编辑问题
// @Accept  json
// @Produce  json
// @Tags qa
// @Param questionEditParam body questionEditParam true "编辑问题参数"
// @Success 200 string Msg "操作成功"
// @Router /api/question/edit [post]
func (api *QAApi) QuestionEdit(c *gin.Context) {
	qaService := c.MustMake(provider.QaKey).(provider.Service)

	param := &questionEditParam{}
	if err := c.ShouldBind(param); err != nil {
		c.ISetStatus(http.StatusBadRequest).IText(err.Error())
		return
	}

	questionOld, err := qaService.GetQuestion(c, param.ID)
	if err != nil || questionOld == nil {
		c.ISetStatus(http.StatusForbidden).IText("操作的问题不存在")
		return
	}

	user := auth.GetAuthUser(c)
	if user == nil || user.ID != questionOld.AuthorID {
		c.ISetStatus(http.StatusForbidden).IText("无权限操作")
		return
	}

	question := &provider.Question{
		ID:      param.ID,
		Title:   param.Title,
		Context: param.Context,
	}
	if err := qaService.UpdateQuestion(c, question); err != nil {
		c.ISetStatus(http.StatusInternalServerError).IText(err.Error())
		return
	}

	c.ISetOkStatus().IText("操作成功")
}
