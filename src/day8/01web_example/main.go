package main

import (
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
)

func greet(w http.ResponseWriter,r *http.Request){
	world := r.FormValue("world")
	fmt.Printf("%v",world)
	fmt.Fprintf(w,"greet hello world! world:%v", world)
}

func userInfo(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"Hello world! %s", time.Now())
}

func userLogin(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET" {
		data,err := ioutil.ReadFile("./login.html")
		if err != nil {
			http.Redirect(w,r,"/404.html",http.StatusNotFound)
			return
		}
		w.Write(data)
	}else if r.Method == "POST"{
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin" && password =="admin"{
			fmt.Fprintf(w,"login succ")
		} else{
			fmt.Fprintf(w,"login failed")
		}
	}
}

func main(){
	http.HandleFunc("/index",greet)
	http.HandleFunc("/user/info",userInfo)
	http.HandleFunc("/user/login",userLogin)
	http.ListenAndServe(":8080",nil)
}
