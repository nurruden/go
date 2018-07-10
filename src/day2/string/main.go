package main

import (
	"fmt"
)

func main(){
	var str string
	// string不可以修改 str[0] = "d"
	str = "abc咋仨"
	// 8个字节
	var b []byte = []byte(str)
	// rune = int32,32个字节
    var chars []rune = []rune(str)

	fmt.Printf("b = %v len(str)=%d\n",b,len(str))
	// c整数对应的字符， %s 进行基本的字符串
	fmt.Printf("%c\n",97)
	fmt.Printf("chars count:%d\n",len(chars))
}