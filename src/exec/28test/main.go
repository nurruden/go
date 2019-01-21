package main

import "fmt"

func change(num *int) {
	fmt.Println(num)    //0xc42000e1f8
	fmt.Println(*num)   //100
	*num = 1000
	fmt.Println(num)    //0xc42000e1f8
	fmt.Println(*num)   //1000
}

type User struct {
	Name string
	Age  int
	mess string
}

func NewUser(name string, age int, mess string) *User {
	return &User{Name:name,Age:age,mess:mess}
}
func (this *User) init(name string, age int, mess string) {
	this.Name = name
	this.Age = age
	this.mess = mess
}

func (this User) GetName() string {
	return this.Name
}


func main() {
	num := 100
	fmt.Println(&num)    //0xc42000e1f8
	change(&num)
	fmt.Println(&num)    //0xc42000e1f8
	fmt.Println(num)     //1000
	//user := NewUser("suoning", 18, "lover")
	//fmt.Println(user, user.mess, user.Name, user.Age)

	var user User
	user.init("nick", 18, "man")
	//(&user).init("nick", 18, "man")
	name := user.GetName()
	fmt.Println(name)
}
