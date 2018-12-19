package middleware

import (
	"github.com/gin-gonic/gin"
	"restful/demo03/handler"
	"restful/demo03/pkg/errno"
	"restful/demo03/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
