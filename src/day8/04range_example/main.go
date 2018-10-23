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
	var users []*User
	for i := 0; i< 15; i++ {
		var user User = User{
			Name:"user01",
			Age:10,
			Address:Address{
				City:"langfang",
				Province:"hebei",
				Code:"100086",
			},
		}
		users=append(users,&user)
	}
	t.Execute(w,users)
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


