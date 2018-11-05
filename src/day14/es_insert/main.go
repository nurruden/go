package main

import (
	"fmt"
	"gopkg.in/olivere/elastic.v2"
	"xlog"
)

type Person struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	City string `json:"city"`
	Desc string `json:"desc"`
}

func main(){
	client,err:=elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))

	if err != nil {
		fmt.Printf("new client faile,err:%v\n",err)
		return
	}

	for i := 0; i < 100000; i++{
		a := &Person{
			Id:i,
			Name:fmt.Sprintf("HELLO %d",i),
			Age:28,
			Desc: `what a fuck`,
		}
		_, err := client.Index().Index("account2").Type("Person").BodyJson(a).Do()
		if err != nil{
			xlog.LogError("Insert fail,err: %v",err)
		}
	}
}
