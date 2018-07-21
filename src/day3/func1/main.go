package main

import (
	"fmt"
	"os"
)


func Calc(a, b int) (s1 int, s2 int) { 
	s1 = a + b
	s2 = a - b
	return
	}
func testCalc(){
    sum,sub := Calc(2, 3)
	fmt.Printf("sum=%d,sub=%d\n",sum,sub)
}
func testAdd(){
	sum := Add()
	fmt.Printf("sum=%d\n",sum)

	sum = Add(1)
	fmt.Printf("sum=%d\n",sum)

	sum = Add(1,3,5)
	fmt.Printf("sum=%d\n",sum)
}

func Add(a ...int) int{
	fmt.Printf("func args count:%d\n",len(a))
	var sum int
	for index,arg := range a{
		fmt.Printf("arg[%d]=%d\n",index,arg)
		sum = sum + arg

	}
	return sum
}

func testDefer(){
	// defer fmt.Printf("hello world\n")
	// fmt.Printf("Wula\n")

	file,err := os.Open("/Users/allanyang/northstar/test/test.py")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n",err)
		return
	}

	defer file.Close()
	var buf[9192]byte
	n,err := file.Read(buf[:])
	if err != nil {
		fmt.Printf("read file failed,err:%v\n",err)
		return
	}
	fmt.Printf("read %d bytes succ,conyent:%s\n",n,string(buf[:]))
	return
}


func main() {
	// testAdd()
	testDefer()
	// 有可能a扩容，所以需要给a赋值
	var a[]byte
	a = append(a,'a')
}