package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := gin.Default()
    //http://127.0.0.1:8080/welcome?firstname=Allan&lastname=Yang
	router.GET("/welcome",func(c *gin.Context){
		firstname := c.DefaultQuery("firstname","Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK,"Hello %s %s",firstname,lastname)
	})
	router.Run()
}