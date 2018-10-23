package main

import "fmt"

type Animal interface {
	Eat()
	Talk()
	Run()
}

type Dog struct {
	name string
}

func (d *Dog) Eat(){
	fmt.Printf("%s is eating\n",d.name)
}

func (d *Dog) Talk(){
	fmt.Printf("%s is talking\n",d.name)
}

func (d *Dog) Run(){
	fmt.Printf("%s is runing\n",d.name)
}


type Pig struct {
	name string
}

func (d *Pig) Eat(){
	fmt.Printf("%s is eating\n",d.name)
}

func (d *Pig) Talk(){
	fmt.Printf("%s is talking\n",d.name)
}

func (d *Pig) Run(){
	fmt.Printf("%s is runing\n",d.name)
}


func main(){
	var dog = &Dog {
		name:"aaa",
	}

	var a Animal
	fmt.Printf("A:%v DOG:%v\n",a,dog)
	a = dog
	a.Eat()
	a.Run()
	a.Talk()

	var pig = &Pig {
		name:"bbb",
	}
	a=pig
	a.Eat()
	a.Run()
	a.Talk()
}