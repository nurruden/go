package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct{
	Message string `json:"message"`
	Code int `json:"code"`
}

type UserInfo struct {
	Result
	UserName string `json:"username"`
	Password string `json:"password"`
}

func handleUserInfo(c *gin.Context){
	username := c.Query("username")
	password := c.DefaultQuery("password","sssddf")
	var result UserInfo = UserInfo{
		UserName:username,
		Password:password,
	}
	result.Code = 0
	result.Message = "Success"
	c.JSON(http.StatusOK,result)
}

func handleUserParams(c *gin.Context){
	username := c.Param("username")
	password := c.Param("password")
	var result UserInfo = UserInfo{
		UserName:username,
		Password:password,
	}
	result.Code = 0
	result.Message = "Success"
	c.JSON(http.StatusOK,result)
}

func main(){
	r := gin.Default()
	r.GET("/user/info",handleUserInfo)
    r.GET("/user/info/:username/:password",handleUserParams)
	r.Run()
}