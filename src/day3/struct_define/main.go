package main

import (
	"fmt"
)

type User struct{
	Name string
	Sex string
	Age int
	AvatarURL string
}

func NewUser(name,sex string,age int, url string) User{
	var user User
	user.Age = 100
	user.Sex = "male"
	user.Name = "JIM"
	user.AvatarURL = "www.baidu.com"
	return user
}

func main(){
	var user User
	user.Age = 100
	user.Sex = "male"
	user.Name = "JIM"
	user.AvatarURL = "www.baidu.com"
	fmt.Printf("name=%s,sex=%s\n",user.Name,user.Sex)

	user02 := User{
		Name:"user01",
		Age:18,
	}
	fmt.Printf("name=%s,sex=%s\n",user02.Name,user02.Sex)

	user04 := NewUser("user04","femal",18,"2221")
	fmt.Printf("user04:%#v\n",user04)
}