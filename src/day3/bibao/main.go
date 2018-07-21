package main

import (
	"fmt"
)

func main(){
	var f= Adder()
	fmt.Print(f(1),"-")
	fmt.Print(f(20),"-")
	fmt.Print(f(300),"-")
}

func Adder() func(int) int{
	var x int
	return func(d int) int{
		x += d
		return x
	}
}