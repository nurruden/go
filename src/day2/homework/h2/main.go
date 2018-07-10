package main

import (
	"fmt"
	"sort"
)

func main(){
	// 数字
	// ages := []int{2,1,5,66,55,23,5,78,98,13}
	// sort.Ints(ages)
	// for index,value := range ages {
	// 	fmt.Println(index,value)
	// }

    // 字符串
	names := []string{"Hello", "World", "private", "folders", "Users", "workspace"}
	sort.Strings(names)
	for _, value := range names {
    	fmt.Println(value)
}
}
