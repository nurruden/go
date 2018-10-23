package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	"net/http"
)

func StatCost(c *gin.Context){
	start := time.Now()
	fmt.Printf("start stat cost\n")
	c.Next()
	lattancy := time.Since(start)
	fmt.Printf("Process request cost:%d ms\n",lattancy/1000/1000)
}

func handleUserInfo(c *gin.Context){
	fmt.Printf("Reauest start process\n")
	time.Sleep(3*time.Second)
	c.JSON(http.StatusOK,"32354")
}

func main(){
	r := gin.Default()
	r.Use(StatCost)

	r.GET("/user/info",handleUserInfo)
	r.Run()
}