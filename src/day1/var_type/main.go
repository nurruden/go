package main

import "fmt"

func main(){
	var (
		a int 
		b string
		c bool
		d int = 8
		e string = "hello"
	)
	_=c
	//布尔值怎么表示
	fmt.Printf("%d\n%s\n%d\n%s\n",a,b,d,e)
}