package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"time"
)

var (
	homeTemplate *template.Template
)

var addr = flag.String("addr","localhost:8080","http serveice address")
var upgrader = websocket.Upgrader{}


func echo(w http.ResponseWriter, r *http.Request){
	c, err := upgrader.Upgrade(w,r,nil)
	if err !=nil {
		log.Print("upgrade:",err)
		return
	}

	defer c.Close()

	for {
		c.WriteMessage(websocket.TextMessage,[]byte("this is server push"))
		time.Sleep(time.Second)
	}
}

func home(w http.ResponseWriter,r *http.Request) {
	homeTemplate.Execute(w,"ws://"+r.Host+"/echo")
}

func initTemplate()(err error){
	homeTemplate,err = template.ParseFiles("./view/index.html")
	if err != nil {
		log.Fatalf("parse indx.html failed,err:%v",err)

	}
	return
}

func main(){
	flag.Parse()
	log.SetFlags(0)
	initTemplate()
	http.HandleFunc("/echo",echo)
	http.HandleFunc("/",home)
	log.Fatal(http.ListenAndServe(*addr,nil))
}