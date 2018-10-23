package main

import (
	"reflect"
	"fmt"
)

type S struct {
	F string `species:"gopher" color:"blue" json:"f"`
}

func main(){
	s :=S{}
	st := reflect.TypeOf(s)
	filed := st.Field(0)
	fmt.Println(filed.Tag.Get("color"),filed.Tag.Get("species"),filed.Tag.Get("json"))
}