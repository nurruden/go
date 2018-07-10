package main

import (
	"fmt"
)

func main(){
	var str string
	// str = "abcdefgh"
    str = "哈哈啦啦的"
	bytes := []byte(str)
	
	// var int i
	// i = 0
	// var i = 0
	for i :=0; i < len(str)/2; i++ {
		// fmt.Printf("%c ",str[i])


		// var tmp = bytes[i]
		// bytes[i] = bytes[len(bytes)-i-1]
		// bytes[len(bytes)-i-1] = tmp
		bytes[i],bytes[len(bytes)-i-1] = bytes[len(bytes)-i-1],bytes[i]
	}
	str = string(bytes)
	fmt.Printf("reverse string: %s\n",str)
}