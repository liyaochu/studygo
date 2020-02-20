package main

import (
	"fmt"
	"net"
)

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:20000")

	if err != nil {
		fmt.Println("连接服务端失败 ", err)
		return
	}
	defer conn.Close()
	//向server端发消息

	_, err = conn.Write([]byte("一夜"))
	if err != nil {
		fmt.Println("发送消息失败 ", err)
		return
	}
}
