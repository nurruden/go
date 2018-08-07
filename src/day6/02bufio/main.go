package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func main(){
	filename  := "/Users/allanyang/gofile/src/day5/homework/image_thumb.go"
	var content []byte
	file,err := os.Open(filename)
	if err != nil{
		fmt.Printf("open %s failed.err:%v\n",filename,err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	
	var buf[4096]byte
	for {
		n,err := reader.Read(buf[:])
		if err != nil && err!= io.EOF{
			fmt.Printf("read %s failed,err:%v",filename,err)
			return

		}
		if err == io.EOF{
			break
		}
		validBuf := buf[0:n]
		// fmt.Printf("%s\n",string(validBuf))
		content = append(content,validBuf...)
	}
	fmt.Printf("content:%s\n",content)
}