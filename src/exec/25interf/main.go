package main

import (
	"fmt"
)

type Human struct {
	name string
	age int
}

type Element interface{

}

type List[] Element


func main(){

	list := make(List,3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Human{"Dennis",70}

	for index,element := range list {
		if value,ok := element.(int);ok {
			fmt.Printf("list[%d] is an int and its value is %d\n",index,value)
		}else if value,ok := element.(string);ok {
			fmt.Printf("list[%d] is an string and its value is %s\n",index,value)
		}else if value,ok := element.(Human); ok {
			fmt.Printf("list[%d] is an Person and its value is %v\n",index,value)
		}else {
			fmt.Println("list[%d] is a different type", index)
		}
	}

}