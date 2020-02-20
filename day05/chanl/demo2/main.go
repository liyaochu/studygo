package main

import "fmt"

//无缓冲通道和有缓冲通道
func receiver(ch chan bool){
   ret:=<-ch   //阻塞
   fmt.Println(ret)
}
func main() {
	bools := make(chan bool) //无缓冲
	//bools := make(chan bool,1) //有缓冲
	bools<-true
     go receiver(bools)  //不写 go 就会死锁
	 //bools<-false  //这个就是利用有缓冲通道,上面写了之后取了,才能放,无缓冲就会deadlock!
	fmt.Println("mian函数 结束")
}
