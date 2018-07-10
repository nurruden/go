package main

import "fmt"

func main(){
	if num := 10; num % 2 == 0 {
		fmt.Println("Num is even")
	} else {
		fmt.Println("Num is odd")
	}
}