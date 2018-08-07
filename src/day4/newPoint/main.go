package main

import (
	"fmt"
)

type User struct {
	name string
	age int
}

// 切片和map只能用make来初始化
func main(){
	var p *int = new(int)
	*p = 1000
	fmt.Printf("p:%v address:%v\n",*p,p)

	var pUser *User = new(User)
	pUser.age =100
	pUser.name ="aa"
	fmt.Printf("user:%v\n",*pUser)
	test2()
}

func test2(){
	var p *[]int = new([]int)
	*p = make([]int,10)
	(*p)[0] = 100
	(*p)[1] = 200
	fmt.Printf("p:%#v\n",*p)

	var p1 *map[string]int = new(map[string]int)
	*p1 = make(map[string]int,10)
	(*p1)["11"] = 100
	(*p1)["22"] = 200
	fmt.Printf("p1:%#v\n",*p1)
}