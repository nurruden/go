package main

import ("fmt")

func modify(b [3]int){
	b[0]=1000
	return
}

func main(){
	var a[3]int
	a[0]=10
	b := a
	b[0]=20
	fmt.Println(a,b)
	modify(a)
	fmt.Println(a)
}