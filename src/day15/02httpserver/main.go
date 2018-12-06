package main

import (
	"fmt"
	"net"
)

func main(){
	fmt.Println("Start server...")
	listen,err := net.Listen("tcp","0.0.0.0:50000")
	if err != nil {
		fmt.Println("listen failed,err:",err)
	}
	for {
		conn,err:=listen.Accept()
		if err != nil {
			fmt.Println("Accept failed,err:",err)
			continue
		}
		go processes(conn)
	}
}

func processes(conn net.Conn){
	defer conn.Close()
	for {
		buf := make([]byte,512)
		_,err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read err:",err)
			return
		}
		fmt.Println("read: ",string(buf))
	}
}