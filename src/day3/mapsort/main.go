package main

import (
	"fmt"
	"sort"
)

func main(){
	var a map[string]int = make(map[string]int,10)
	for i :=0;i<10;i++ {
		keya := fmt.Sprintf("key%d",i)
		fmt.Print(keya)
		a[keya] = i
	}
	var keys[]string
	for key,_:=range a{
		keys = append(keys,key)
	}
	fmt.Print(keys)
	sort.Strings(keys)
	for _,key :=range keys{
		fmt.Printf("key:%s=%d\n",key,a[key])
	}
}