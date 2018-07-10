package main

import (
	"fmt"
)
func subtactOne(numbers []int){
	for i := range numbers {
		numbers[i]-=2
	}

}
func main(){
	nos := []int{8,7,6}
	fmt.Println("Before function call:",nos)
	subtactOne(nos)
	fmt.Println("After function call:",nos)
}