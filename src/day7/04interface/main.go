package main

import "fmt"

type Empolyee interface{

	Calc() float32

}

type Developer struct {
	Name string
	Base float32
}

func (d *Developer) Calc() float32{
	return d.Base
}
type PM struct {
	Name string
	Base float32
	Option float32
}

func (p *PM) Calc() float32{
	return p.Base + p.Option
}

type YY struct {
	Name string
	Base float32
	Option float32
	Rate float32
}

func (p *YY) Calc() float32{
	return p.Base + p.Option +p.Rate
}

type EmployeeMgr struct {
	employeelist []Empolyee
}

func (e *EmployeeMgr) Calc() float32{
	var sum float32
	for _,v := range e.employeelist{
		sum += v.Calc()
	}
	return sum
}

func (e *EmployeeMgr) AddEmployee(d Empolyee){
	e.employeelist = append(e.employeelist,d)
}

func main(){
	var e = &EmployeeMgr{}

	dev := &Developer{
		Name:"develop",
		Base:10000,
	}
	e.AddEmployee(dev)

	pm := &PM{
		Name:"pm",
		Base:10000,
		Option:12000,
	}
	e.AddEmployee(pm)


	yy := &YY{
		Name:"operate",
		Base:10000,
		Option:12000,
		Rate:1.2,
	}
	e.AddEmployee(yy)

	sum := e.Calc()
	fmt.Printf("sum:%f\n",sum)
	}