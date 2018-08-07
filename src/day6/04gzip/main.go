package main

import (
	"os"
	"fmt"
	"io"
	// "bufio"
	"compress/gzip"
)

func main(){
	filename  := "/Users/allanyang/Desktop/copy.gz"
	
	file,err := os.Open(filename)
	if err != nil{
		fmt.Printf("open %s failed.err:%v\n",filename,err)
		return
	}
	defer file.Close()
	reader,err := gzip.NewReader(file)
	if err != nil{
		fmt.Printf("gzip %s failed.err:%v\n",filename,err)
		return
	}
	var content []byte
	var buf[4096]byte
	for {
		n,err := reader.Read(buf[:])
		if err != nil && err!= io.EOF{
			fmt.Printf("read %s failed,err:%v",filename,err)
			return

		}
		if n > 0 {
			validBuf := buf[0:n]
		// fmt.Printf("%s\n",string(validBuf))
		    content = append(content,validBuf...)
		}
		if err == io.EOF{
			break
		}
		
	}
	fmt.Printf("content:%s\n",content)
}