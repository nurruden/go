package main

import (
	"fmt"
)

func main(){
	veggies := []string{"potatoes","toma","bri"}
	fruits := []string{"orange","apple"}
	copy(veggies,fruits)
	fmt.Printf("a:%v b:%v\n",veggies,fruits)

}