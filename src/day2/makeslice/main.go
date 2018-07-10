package main

import (
	"fmt"
)


func main(){
	var a[]int
	a = make([]int,5,5)
	fmt.Printf("%v\n",a)
	fmt.Printf("%v len:%d cap: %d\n",a,len(a),cap(a))
	a[0] = 100
	fmt.Printf("%v len:%d cap: %d,addr:%p\n",a,len(a),cap(a),a)
	a = append(a,200)
	fmt.Printf("%v len:%d cap: %d,addr:%p\n",a,len(a),cap(a),a)
}