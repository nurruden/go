package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := gin.Default()
	////curl -X POST http://127.0.0.1:8080/form_post -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nick=rsj217" | python -m json.tool
	//router.POST("/form_post",func(c *gin.Context) {
	//	message := c.PostForm("message")
	//	nick := c.DefaultPostForm("nick","anonymous")
	//
	//	c.JSON(http.StatusOK,gin.H{
	//		"status": gin.H{
	//			"status_code": http.StatusOK,
	//			"status": "ok",
	//		},
	//		"message": message,
	//		"nick": nick,
	//	})
	//})
	//curl -X PUT http://127.0.0.1:8080/post -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nick=rsj217&id=ss&page=1" | python -m json.tool
	router.PUT("post",func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page","0")
		name := c.PostForm("message")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s\n",id,page,name,message)
		c.JSON(http.StatusOK,gin.H{
			"status_code": http.StatusOK,
		})
	})
	router.Run()
}