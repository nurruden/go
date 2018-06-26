package main

import "fmt"

func test_str1() {
	var b string = "hello\n\n\n"
	var c = "hello"
	//如何不换行
	fmt.Printf("b=%v and c=%s\n",b,c)

}

func test_str2() {
	var b string = `
	床前明月光
	疑是地上霜
	`
	fmt.Printf("b %s\n",b)

}

func test_char() {
	//rune 代表char类型
	var c rune
	var d rune
	c = 20320
	d = '好'
	fmt.Printf("c=%c\n",c)
	fmt.Printf("d=%c\n",d)
}


func main() {
	test_str1()
	test_str2()
	test_char()
}