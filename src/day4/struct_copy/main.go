package main

import (
	"fmt"
	"time"
)

type User struct{
	s1 [10000000]int64
	s2 [10000000]int64
	s3 [10000000]int64
	s4 [10000000]int64
}

func (u User) set(){
	for i:=0;i<len(u.s1);i++{
		u.s1[i] = 1
		u.s2[i] = 1
		u.s3[i] = 1
		u.s4[i] = 1
	}
}
func (u *User) set1(){
	for i:=0;i<len(u.s1);i++{
		u.s1[i] = 1
		u.s2[i] = 1
		u.s3[i] = 1
		u.s4[i] = 1
	}
}
func main(){
	var u *User = new(User)
	start := time.Now().UnixNano()
	u.set1()
	end := time.Now().UnixNano()
	fmt.Printf("cost: %d ns\n",(end-start)/1000)
}