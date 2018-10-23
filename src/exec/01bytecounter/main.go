package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int,error){
	fmt.Printf("c is:%v\n",*c)
	fmt.Printf("*c is:%v\n",*c)
	*c += ByteCounter(len(p))
	fmt.Printf("c is:%v\n",*c)
	fmt.Printf("*c is:%v\n",*c)
	return len(p), nil
}

func main(){
	var c ByteCounter

	c.Write([]byte("hello"))

	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c,"hello, %s\n",name)
	fmt.Println(c)

	x := 1
	p := &x
	fmt.Printf("p is: %v\n",p)
	*p = 2
	fmt.Printf("p is: %v\n",p)
	fmt.Println(x)
}
