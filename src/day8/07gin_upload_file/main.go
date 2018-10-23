package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
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

func handleUpLoad(c *gin.Context){
	file,err := c.FormFile("testfile")
	if err != nil {
		fmt.Printf("upload failed \n")
		return
	}

	filename := fmt.Sprintf("/Users/northstar/Downloads/%s",file.Filename)
	err = c.SaveUploadedFile(file,filename)
	if err != nil {
		fmt.Printf("Save file failed:%v\n",err)
		return
	}
	c.JSON(http.StatusOK,"file upload success")
}

func handleMultiUpLoad(c *gin.Context){
	form,err := c.MultipartForm()

	if err != nil {
		fmt.Printf("upload failed \n")
		return
	}

	files := form.File["testfile"]
	for _,file := range files {
		filename := fmt.Sprintf("/Users/northstar/Downloads/%s",file.Filename)
		err = c.SaveUploadedFile(file,filename)
		if err != nil {
		fmt.Printf("Save file failed:%v\n",err)
		return
	   }
	}

	c.JSON(http.StatusOK,"file upload success")
}

func main(){
	r := gin.Default()
	r.POST("/file/upload",handleUpLoad)
	r.POST("/files/upload",handleMultiUpLoad)
	r.Run()
}