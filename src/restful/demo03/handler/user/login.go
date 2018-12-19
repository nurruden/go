package user

import (
	"github.com/gin-gonic/gin"
	. "restful/demo03/handler"
	"restful/demo03/model"
	"restful/demo03/pkg/auth"
	"restful/demo03/pkg/errno"
	"restful/demo03/pkg/token"
)

func Login(c *gin.Context) {
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}
	SendResponse(c, nil, model.Token{Token: t})
}
