package main

import "fmt"

func main(){
	var a int = 100
	var b *int = &a
	var c *int = b
	*c = 200
	fmt.Printf("Address of a is %v, value of a is:%v\n",&a,a)
	fmt.Printf("Address of b is %v, value of c is:%v\n",&b,b)
    fmt.Printf("Address of b is %v, value of c is:%v\n",&c,c)

}