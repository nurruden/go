package main

import ( 

    // "sort"
	"fmt"
)

func initMap1(){
	var user map[string]int = make(map[string]int,500)
	user["abc"] = 001
	user["user01"] = 002
	user["user02"] = 003
	fmt.Printf("user：%v",user["user001"])
}

func initMap2(){
	var m map[string]int
	m = map[string]int{
		"user003":001,
		"user002":002,
	}
	m["user03"]=003
	fmt.Printf("user：%v",m["user01"])
}
var whiteUser map[int]bool = map[int]bool {
	32423:true,
	3483943:true,
	1:true,
}

func isWhiteUser(userId int) bool{
	_,ok := whiteUser[userId]
	return ok
}

func testtransver(){
	var m map[string]int 
	m =map[string]int{
		"user001":1,
		"user002":2,
	}
	m["user003"]=3
	for key,value := range m{
		fmt.Printf("key:%s value:%d\n",key,value)
	}
}

func testdelete() {

	var user map[string]int = make(map[string]int, 5000) //map只有容量,没有长度

	user["abc"] = 38
	user["user01"] = 10001
	user["user02"] = 10002

	for key, value := range user {
		fmt.Printf("key : %v ,value:%v \n", key, value)
	}

	delete(user, "user01")

}
func testMapCopy() {

	var user map[string]int = make(map[string]int, 5000) //map只有容量,没有长度

	user["abc"] = 38
	user["user01"] = 10001
	user["user02"] = 10002

	fmt.Printf("%v\n", user)

	b := user
	b["abc"] = 28
	fmt.Printf("%v\n", user)

}

func testMapSlice(){
	var s []map[string]int
	s = make([]map[string]int,5)
	for k,v := range s{
		fmt.Printf("index:%d,value:%v\n",k,v)
	}
	s[0] = make(map[string]int,16)
	s[0]["abc"] = 100

	for key,value := range s[0]{
		fmt.Printf("key:%s,value:%v\n",key,value)
	}
	
}

func main() {
	// initMap1()

	// userId := 1
	// if isWhiteUser(userId) {
	// 	fmt.Printf("is true\n")
	// } else {
	// 	fmt.Printf("is not\n")
	// }

	// testtransver()
	testMapSlice()

}