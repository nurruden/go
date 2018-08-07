package main

import (
	"os"
	"fmt"
	// "io"
	// "bufio"
)

func isFileExists(filename string)bool {
	_,err := os.Stat(filename)
	if os.IsNotExist(err){
		return false
	}
    return true
}

func main(){
	filename  := "/Users/allanyang/Desktop/copy12"
	var file *os.File
	var err error
	if isFileExists(filename){
        file,err = os.OpenFile(filename,os.O_APPEND|os.O_WRONLY,0777)
	}else{
		file,err = os.Create(filename)
	}
	if err != nil{
		fmt.Printf("open %s failed.err:%v\n",filename,err)
		return
	}
	defer file.Close()
	n,err:=file.WriteString("testtest")
	if err !=nil{
		fmt.Printf("write failed,err:%v\n",err)
		return
	}
	fmt.Printf("write %d succ\n",n)
}