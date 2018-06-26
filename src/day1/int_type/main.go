package main

import "fmt"

func main(){
	var a int = 10
	// a := 10
	var b int32 =16
	var c int
	c = a + int(b)
	var d int = int(b)
	fmt.Print(a,b,c,d)
	fmt.Println()
	fmt.Printf("%x\n",c)
}