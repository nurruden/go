package main

import (
	"fmt"
)

type Animal struct{
	Name string
	Age int
}

type People struct{
	Animal
	Sex string
}

func main(){
	var p People
	p.Age =18
	p.Name = "US"
	p.Sex = "female"
	fmt.Printf("name%s,age%d\n",p.Name,p.Age)
}