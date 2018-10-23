package main

import (
	"fmt"
	"html/template"
	"net/http"
)
var (
	t *template.Template
)

type User struct {
	Name string
	Age int
	Address Address
}

type Address struct {
	City string
	Province string
	Code string
}
func initTemplate()(err error){
	t,err = template.ParseFiles("./index.html")

	if err != nil {
		fmt.Printf("load template failed,err:%v\n",err)
		return
	}
	return

}

func handleUserInfo(w http.ResponseWriter,r *http.Request){
	var user User = User{
		Name:"user01",
		Age:100,
		Address:Address{
			City:"tianjin",
			Province:"tianjin",
			Code:"2000",
		},

	}
	t.Execute(w,user)
}


func main(){
	err := initTemplate()
	if err != nil {
		fmt.Printf("init failed, err:%v\n",err)
		return
	}
	http.HandleFunc("/user/login",handleUserInfo)
	err = http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Printf("listen failed, err:%v\n",err)
		return
	}
}

