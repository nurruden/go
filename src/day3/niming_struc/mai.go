package main

import (
	"fmt"
)
type Address struct{
	Province string
	City string
}
type User struct{
	Name string
	Sex string
	Age int
	AvatarURL string
	int
	string
	Address
}

func main(){
	var user User
	user.int =100
	user.string = "as"
	user.Address.City = "shijiazhuang"
	user.Address.Province = "hebei"
	fmt.Printf("user:%#v\n",user)
	// fmt.Printf("user.int:%d,user.string%s\n",user.int,user.string)

	user01 := User{
		int:100,
		string:"Ss",
		Address: Address{
			City:"BEIJING",
			Province:"BEIJING",
		},
	}
	fmt.Printf("user:%#v\n",user01)
	// fmt.Printf("city=%s,province=%s\n",user01.Address.City,user01.Address.Province)
}