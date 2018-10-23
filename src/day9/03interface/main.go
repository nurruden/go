package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age int
	Gender string
}

func (u *User) Print(){
	fmt.Printf("Name:%s,Age:%d,Gender:%s\n",u.Name,u.Age,u.Gender)
}

func (u *User) SetName(name string){
	u.Name = name
}

func TestValue(a interface{}){
	v:=reflect.ValueOf(a)
	m := v.MethodByName("Print")

	var args []reflect.Value
	m.Call(args)

	v=reflect.ValueOf(a)
	m=v.MethodByName("SetName")

	args = args[0:0]
	args = append(args,reflect.ValueOf("HELLO WORLD"))
	m.Call(args)
}

func main(){
	var user User
	user.Name = "xx"
	user.Age = 100
	user.Gender = "male"
	TestValue(&user)
	fmt.Printf("USER:%#v\n",user)
}