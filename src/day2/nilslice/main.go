package main

import (
	"fmt"
)

func main() {
	var a []int
	if a == nil{
		fmt.Printf("a is null\n")
	}
	fmt.Printf("a:%v,addr:%p\n",a,a)
	a = append(a,2,2,3)

	fmt.Printf("a:%v,addr:%p\n",a,a)
	var b []int = []int{1,2,3}
	a = append(a,b...)
	fmt.Printf("a:%v,addr:%p\n",a,a)
}