package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:30000")
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
	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("发送消息失败 ", err)
		return
	}
	fmt.Println("",string(buf[:n]))
}
