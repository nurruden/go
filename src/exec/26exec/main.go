package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	School string
	SpendMoney float64
}

type Employee struct {
	Human
	Company string
	EarnMoney float64
}

func(h Human)SayHi(){
	fmt.Printf("Hi, I am %s, you can reach me on phone %s\n",h.name,h.phone)
}

func(h Human)Sing(lyrics string){
	fmt.Println("I am singing  ",lyrics)
}

func(e Employee)SayHi(){
	fmt.Printf("Hi, I am %s, I work at %s,you can reach me on phone %s \n",e.name,e.Company,e.phone)
}

type Men interface{
	SayHi()
	Sing(lyrics string)
}

func main(){
	Allan := Employee{Human{"Allan",30,"33-22-11"},"Juniper",30000}
	Felix := Employee{Human{"Felix",30,"43-22-11"},"Juniper",30000}
	Horton := Student{Human{"Horton",40,"33-55-33"},"PKU",4000}
	Nora := Student{Human{"Nora",30,"11-22-33"},"TSU",35000}

	var i Men
	i = Allan
	fmt.Println("This is empoyee: Allan")
	i.SayHi()
	i.Sing("love love love")

	i = Horton
	fmt.Println("This is student: Horton")
	i.SayHi()
	i.Sing("hate hate")

	x := make([]Men, 4)
	x[0],x[1],x[2],x[3]=Allan,Felix,Horton,Nora

	for _,value := range x {
		value.SayHi()
	}

	fmt.Println("********************************************************************")
	var y float64 = 3.4
	v := reflect.ValueOf(y)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}