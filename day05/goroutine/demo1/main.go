package main

import (
	"fmt"
	"sync"
)
//一个 goroutine 对应一个函数,函数结束了 goroutine就结束了

 var wait sync.WaitGroup  //实现优雅的等待
func hello() {
	fmt.Println("Hello Goroutine!")
	wait.Done()
}
func main() {
	defer  fmt.Println("哈哈哈")
	wait.Add(1)
    go hello()  //创建一个goroutine 在新的goroutine中执行hello函数
    fmt.Println("main goroutine done!")
   // time.Sleep(time.Second)
	wait.Wait()
}