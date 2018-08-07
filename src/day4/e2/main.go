package main

import (
	"fmt"
)



type Student struct{
	Name string
	Age int
}
// 定义值类型为student的getname方法
func (s Student) GetName() string{
    return s.Name
}
// 结构体值类型，用make初始化的是引用类型

// 定义指针类型为student的setname方法
func (s *Student) SetName(name string){
	s.Name=name
}

func main(){
	var s1 Student = Student{
		Name:"S1",
		Age:18,
	}
	name := s1.GetName()
	fmt.Printf("name=%s\n",name)

	s1.SetName("s2")
	name = s1.GetName()
	fmt.Printf("name=%s\n",name)
	s1.Print0()
}