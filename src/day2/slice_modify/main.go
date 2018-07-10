package main

import (
	"fmt"
)


func main(){
	var a[]int

	var b[6]int = [6]int{1,2,3,4,5,6}
	a = b[0:2]
    a[0] = 1000
	fmt.Printf("a=%v\n",a)
	fmt.Printf("b=%v\n",b)


}