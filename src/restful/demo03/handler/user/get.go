package user

import (
	"github.com/gin-gonic/gin"
	"restful/demo03/handler"
	"restful/demo03/model"
	"restful/demo03/pkg/errno"
)

func Get(c *gin.Context) {
	username := c.Param("username")

	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, user)
}
