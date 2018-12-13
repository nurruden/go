package user

import (
	"github.com/gin-gonic/gin"
	"restful/demo03/handler"
	"restful/demo03/pkg/errno"
	"restful/demo03/service"
)

func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})

}
