package main

import ("fmt")

func main(){
	var a [6]int
	var b [3]int=[3]int{1,2,3}
	c:=[3]int{4,5,6}
	d:=[...]int{7,8}
	e:=[3]int{2:10}
	f:=[3]int{10}
	a[0] = 100
	fmt.Printf("a=%v\n",a)
	fmt.Printf("b=%v\n",b)
	fmt.Printf("c=%v\n",c)
	fmt.Printf("d=%v\n",d)
	fmt.Printf("e=%v\n",e)
	fmt.Printf("f=%v\n",f)
	for index,val := range a{
		fmt.Printf("%d,%d \n",index,val)
	}
}