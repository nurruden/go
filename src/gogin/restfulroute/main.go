package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := gin.Default()
    /////user/jjkk,/user/jjkk/,/user/12都可以，/user/不可以
	//router.GET("/user/:name",func(c *gin.Context) {
	//	name := c.Param("name")
	//	c.String(http.StatusOK,"Hello %s",name)
	//})

	///user/22/sw/w可以
	router.GET("/user/:name/*action",func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK,message)
	})


	router.Run(":8000")
}

