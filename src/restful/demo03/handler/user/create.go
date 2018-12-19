package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	. "restful/demo03/handler"
	"restful/demo03/model"
	"restful/demo03/pkg/errno"
	"restful/demo03/util"
)

func Create(c *gin.Context) {
	log.Info("User Create function is called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		log.Error("Bing failed,", err)
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		log.Error("Validation failed.", err)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Error("Password encrypted failed", err)
		return
	}
	log.Info("start to create")
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	SendResponse(c, nil, rsp)

}

//func (r *CreateRequest) checkParam() error {
//	if r.Username == "" {
//		return errno.New(errno.ErrValidation,nil).Add("username is empty")
//	}
//	if r.Password == "" {
//		return errno.New(errno.ErrValidation,nil).Add("password is empty.")
//	}
//	return nil
//}
