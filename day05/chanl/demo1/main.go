package main

import "fmt"

func main() {
	//channel 是一个引用类型
	var ch1 chan  int
	var ch2 chan  string
	fmt.Println("ch1",ch1)
	fmt.Println("ch2",ch2)
	ch3:=make(chan int,1) //1代表容量,1个缓冲区

	//发送和接收都用这个符号<-
	ch3<-10  //把10发送到ch3中
	//<-ch3 //从管道中接收,直接丢弃
	ret:=<-ch3
	fmt.Println(ret)
	//关闭通道,从关闭的管道里拿值是可以的
	close(ch3)
}
