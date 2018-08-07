package main

import "fmt"

func modify(a int){
	a = 1000
	fmt.Printf("2.Address of a:%v,value of a:%v\n",&a,a)
}

func modify2(a *int) {
	*a = 1000
	fmt.Printf("4.Address of a:%v,value of a:%v\n",&a,a)
}

func main(){

	var b int =100
	fmt.Printf("1,Address of b:%v,value of b:%v\n",&b,b)
	modify(b)
	var p *int = &b
	fmt.Printf("3.Address of p:%v,value of p:%v\n",&p,p)
	fmt.Printf("Address of p:%v\n",&p)
	modify2(p)
	fmt.Printf("b=%d\n",b)
}