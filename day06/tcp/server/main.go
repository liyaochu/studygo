package main

import (
	"fmt"
	"net"
)

//1.接收k客户端请求建立连接
//2.创建goroutine处理连接
func process(conn net.Conn){
	defer conn.Close()
	//从连接中获取数据
	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("结束客户端的消息失败 失败", err)
		return
	}
	fmt.Println("接收客服端信息",string(buf[:n]))
}
func main() {
	//首先监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("127.0.0.1:20000 失败", err)
		return
	}
	defer listen.Close()
	//接收客户端请求
	for {
		//建立连接
		conn, err := listen.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println("连接 失败", err)
			continue
		}
		//2.创建goroutine处理连接
      go process(conn)
	}

}
