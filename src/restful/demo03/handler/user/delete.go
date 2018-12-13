package user

import (
	"github.com/gin-gonic/gin"
	"restful/demo03/handler"
	"restful/demo03/model"
	"restful/demo03/pkg/errno"
	"strconv"
)

func Delete(c *gin.Context) {

	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, nil)

}
