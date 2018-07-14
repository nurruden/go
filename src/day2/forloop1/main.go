package main

import (
	"fmt"

)

// func main(){
// 	for{
// 		fmt.Printf("Hello")
// 		break
// 	}
// }

func main() { 
	for i := 0;i <= 10; { // initialisation and post are omitted 
	fmt.Printf("%d\n", i)
	i += 2
	} }