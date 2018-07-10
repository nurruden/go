package main

import (
	"fmt"
)

func main(){
	darr:= [...]int{57,89,1,2,3,4,5,6,7}
	dslice := darr[2:5]
	fmt.Println("array before:",darr)
	for i := range dslice{
		dslice[i]++
	}
	fmt.Println("After arrary: ",darr)
}