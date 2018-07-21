package main

import (
	"fmt"
)

func main(){
	var mapSlice []map[string]int
	mapSlice = make([]map[string]int,5)
	fmt.Println("before map init")
	for index,value := range mapSlice{
		fmt.Printf("index:%d value:%v\n",index,value)
	}
	fmt.Println()
	mapSlice[0] = make(map[string]int, 10) 
	mapSlice[0]["a"] = 1000 
	mapSlice[0]["b"] = 2000 
	mapSlice[0]["c"] = 3000 
	mapSlice[0]["d"] = 4000 
	mapSlice[0]["e"] = 5000
	fmt.Println("after map init")
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value) 
	}
}