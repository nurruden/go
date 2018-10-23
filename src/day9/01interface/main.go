package main

import (
	"reflect"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age int
	Gender string
}

func TestValue(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type()

	switch t.Kind() {
	case reflect.Struct:
		filedNum := t.NumField()
		fmt.Printf("Filed num:%v\n",filedNum)
		for i := 0;i < filedNum;i++{
			filed := t.Field(i)
			vfiled := v.Field(i)

			fmt.Printf("filed[%d], name:%s, json key:%s,val:%v\n",i,filed.Name,filed.Tag.Get("json"),vfiled.Interface())
		}
	}
}

func main() {
	var user User
	user.Name = "felix"
	user.Age = 100
	user.Gender = "male"
	TestValue(user)
	fmt.Printf("USER:%#v\n",user)
}