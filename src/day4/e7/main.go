package main

import (
	"fmt"
)

type User struct {
	name string
	age int
}


func main(){
	 
    var p *[]User = new([]User)
    *p = make([]User,10)
    (*p)[0].name = "a"
	(*p)[0].age = 1
	(*p)[1].name = "b"
	(*p)[1].age = 1
	fmt.Printf("%v,%v\n",&p,*p)
	fmt.Printf("%v\n",(*p)[0].name)

}
 