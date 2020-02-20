package main

import (
	"fmt"
	"net"
)

func main(){
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 30000,
	})

	if err != nil {
		fmt.Println("127.0.0.1:30000 失败", err)
		return
	}
	defer listen.Close()

	//循环收发数据
	for{
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("接收消息失败", err)
			return
		}
		fmt.Printf("接受到%v的消息:%v\n",addr,string(buf[:n]))
		//回复消息
		n, err = listen.WriteToUDP([]byte("收到了"), addr)
		if err != nil {
			fmt.Println("接收消息失败", err)
			return
		}
	}
}
