package main

import (
	"fmt"
)


func main(){
	var a[]int
	var c[]int
	var d[]int
	var b[6]int = [6]int{1,2,3,4,5,6}
	a = b[0:2]
	c = b[:]
	d = b[2:]
	fmt.Printf("a=%v\n",a)
	fmt.Printf("a=%v\n",c)
	fmt.Printf("a=%v\n",d)
}