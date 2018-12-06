package main

import "fmt"

type Circle struct {
	redius float64
}

func (c Circle) getArea() float64{
	return 3.14 * c.redius
}

func main(){
	var c1 Circle
	c1.redius = 100.00
	fmt.Printf("Area of Circle(c1) = %.2f \n",c1.getArea())
}