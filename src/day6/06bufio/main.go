package main

import (
	"os"
	"fmt"
	"bufio"
)

func main(){
	filename  := "/Users/allanyang/Desktop/copy12"
	var file *os.File
	var err error

	file,err = os.Create(filename)
	
	if err != nil{
		fmt.Printf("open %s failed.err:%v\n",filename,err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("222")
	writer.Flush()
}