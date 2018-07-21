package main

import (
	"fmt"
)

func main() {
	// // first init
	// var user map[string]int = make(map[string]int)
	// user["abc"] = 38
	// fmt.Printf("abc=%v\n",user)

	// second init
	a := make(map[string]int) 
	a["steve"] = 12000
    a["jamie"] = 15000
    a["mike"] = 9000
    fmt.Println("a mapa contents:", a)
}