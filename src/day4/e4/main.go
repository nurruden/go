package main

import "fmt"

func main(){
	b := 255
	a := &b
	fmt.Printf("Address of b is %v\n",a)
	fmt.Printf("Value of b is %v\n",*a)
	*a++
	fmt.Printf("New value of b is %v\n",b)
}
