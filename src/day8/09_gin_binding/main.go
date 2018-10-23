package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



type UserInfo struct {

	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Age int `form:"age" json:"age"`
	Sex string `form:"sex" json:"sex"`
}

func handleUserInfo(c *gin.Context){
	var userinfo  UserInfo

	err := c.ShouldBind(&userinfo)
	if err !=nil {
		return
	}
	c.JSON(http.StatusOK,userinfo)
}

func handleUserInfoJson(c *gin.Context){
	var userinfo  UserInfo

	err := c.ShouldBindJSON(&userinfo)
	if err !=nil {
		return
	}
	c.JSON(http.StatusOK,userinfo)
}

func handleUserInfoQuery(c *gin.Context){
	var userinfo  UserInfo

	err := c.ShouldBind(&userinfo)
	if err !=nil {
		return
	}
	c.JSON(http.StatusOK,userinfo)
}

func main(){
	r := gin.Default()
	r.POST("/user/info",handleUserInfo)
	r.POST("/user/infojson",handleUserInfoJson)
	r.GET("/user/info",handleUserInfoQuery)
	r.Run()
}
