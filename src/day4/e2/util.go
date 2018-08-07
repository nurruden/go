package main

import (
	"fmt"
)

func (s *Student) Print0(){
	fmt.Printf("Student:%#v",*s)
}