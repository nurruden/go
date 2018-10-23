package main

import "fmt"

type byteCounter int

func (c *byteCounter) Write(p []byte)(int,error) {
	*c += byteCounter(len(p))
	return len(p),nil
}

func main(){
	var c byteCounter

	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "allan"
	fmt.Fprintf(&c, "hello,%s",name)
	fmt.Println(c)
}
