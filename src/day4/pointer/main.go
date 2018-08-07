package main

import (
	"fmt"
)

func testPointer(){
	var a *int
	var b int = 200
	// &寻址，寻找值对应的地址
	a = &b
	
	fmt.Printf("value of A:%v\n",a)
	fmt.Printf("address of b %v\n",&b)
	fmt.Printf("address of a %v\n",&a)
	// 取址，取地址对应的值
	*a = 300
	fmt.Printf("Value of b %v\n",*a)
	fmt.Printf("Value of b %v\n",b)
	
	// var c int = 300
	// a = &c
	// fmt.Printf("a=%v\n",*a)
}

func main(){
	var b int32
	b = 156
	var a *int32
	fmt.Printf("Address of a :%v\n",a)
	a = &b
	fmt.Printf("%v\n",a)
	fmt.Printf("%v\n",*a)
	testPointer()
}