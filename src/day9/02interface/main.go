package main

import (
	"fmt"
	"reflect"
)

type S struct {
	A int
	B string
}


func (s *S) Test(){
	fmt.Println("This is a test")
}

func main(){
	s := S{23,"SKIDO"}
	v := reflect.ValueOf(&s)
	t := v.Type()
	v.Elem().Field(0).SetInt(100)
	fmt.Println("method num:",v.NumMethod())
	for i := 0;i<v.Elem().NumField();i++{
		f := v.Elem().Field(i)
		fmt.Printf("%d: %s %s = %v\n",i,t.Elem().Field(i).Name,f.Type(),f.Interface())
	}
}