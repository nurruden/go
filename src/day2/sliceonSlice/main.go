package main

import (
	"fmt"
)

func main(){
	var a []int
	var b [6]int = [6]int{1,2,3,4,5,6}
	a = b[4:5]
	fmt.Printf("len(a):%d cap:%d\n",len(a),cap(a))
	c := a[0:1]
	fmt.Printf("c=%v\n",c)
}