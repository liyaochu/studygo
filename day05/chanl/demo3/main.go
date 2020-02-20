package main

import "fmt"

//接收值判断通道是否关闭
func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ints := make(chan int, 100)
	go send(ints)
	//利用for循环去通道接收值
	//for   {
    //   ret , ok :=<-ints
    //   //可以用这个方式去判断管道是否有值
    //   if !ok{
    //   	break
	//   }
    //   fmt.Println(ret)
	//}

	//利用for range 去通道取值
	for res:= range ints{
		fmt.Println(res)
	}
}
