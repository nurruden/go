package main

import (
	
	"fmt"
	"io/ioutil"
)

func main(){
	filename  := "/Users/allanyang/gofile/src/day5/homework/image_thumb.go"
   
	content,err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Printf("open %s failed.err:%v\n",filename,err)
		return
	}
	fmt.Printf("content:%s",string(content))
}