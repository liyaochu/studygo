package main

import (
	"fmt"
	"time"
)

//select 多路复用器
func method(ch chan string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("method:%d", i)
		time.Sleep(time.Second)
	}
}
func method2(ch chan string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("method2:%d", i)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan string, 100)
	ch1 := make(chan string, 100)
	go method(ch)
	go method2(ch)
	for {
		select {
		case ret := <-ch:
			fmt.Println(ret)
		case ret := <-ch1:
			fmt.Println(ret)
		default:
			fmt.Println("取不到值")
			time.Sleep(time.Second)
		}
	}

}
