package main

import (
	"math/rand"
	"fmt"
	"time"
)


func main(){
	var number int
	rand.Seed(time.Now().UnixNano())
	number = rand.Intn(100)
	fmt.Printf("Pleas enter a number between 0 and 100\n")
	for {
		var input int
		fmt.Scanf("%d",&input)
		var flag bool = false
		switch {
		case number > input:
			fmt.Printf("Need a bigger number\n")
		case number == input:
			fmt.Printf("Correct\n")
			flag = true
		case number < input:
			fmt.Printf("Need a smaller number\n")
		}
		if flag {
            break
		}

	}
}