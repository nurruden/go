package main

import (
	"fmt"
)

func main(){
	var str string
	// str = "abcba"
    str = "abcsa"
	var chars []rune = []rune(str)
	var tmp = str
	// var int i
	// i = 0
	// var i = 0
	for i :=0; i < len(chars)/2; i++ {
		// fmt.Printf("%c ",str[i])


		// var tmp = bytes[i]
		// bytes[i] = bytes[len(bytes)-i-1]
		// bytes[len(bytes)-i-1] = tmp
		chars[i],chars[len(chars)-i-1] = chars[len(chars)-i-1],chars[i]
	}
	str = string(chars)
	// fmt.Printf("str:%s,tmp:%s",str,tmp)
	if str == tmp {
		fmt.Printf("1\n")
	} else {
		fmt.Printf("2\n")
	}
	// fmt.Printf("reverse string: %s\n",str)
}