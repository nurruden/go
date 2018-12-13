package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"restful/demo03/handler"
	"restful/demo03/model"
	"restful/demo03/pkg/errno"
	"restful/demo03/util"
	"strconv"
)

func Update(c *gin.Context) {
	log.Info("Update function called", lager.Data{"X-Request-Id": util.GetReqID(c)})
	userId, _ := strconv.Atoi(c.Param("id"))

	var u model.UserModel
	if err := c.Bind(&u); err != nil {

		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// We update the record based on the user id.
	u.Id = uint64(userId)

	// Validate the data.
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err := u.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
