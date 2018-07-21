package main

import (
	"fmt"
	"encoding/json"
)

type User struct{
	Name string `json:"name"`
	Sex string `json:"sex"`
	Age int `json:"age"`
	AvatarURL string `json:"avatarurl"`
}

func main(){
	var user User
	user.Age = 100
	user.Sex = "male"
	user.Name = "JIM"
	user.AvatarURL = "www.baidu.com"
	data,err := json.Marshal(user)
	if err != nil {
		fmt.Printf("marshal failed,err:%s\n",err)
		return
	}
	fmt.Printf("json:%v\n",string(data))
}