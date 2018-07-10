package main

import "fmt"

func main(){
	num := 75
	switch {
	case num >=0 && num <= 50:
		fmt.Print("Greater than 0 and less than 50\n")
		fallthrough
	case num >50 && num <=100:
		fmt.Print("Greater than 50 and less than 100\n")
		fallthrough
	default:
		fmt.Print("Not in 0-100\n")
	}
}