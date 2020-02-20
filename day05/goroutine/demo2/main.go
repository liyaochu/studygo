package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wait sync.WaitGroup //实现优雅的等待
func a() {
	defer wait.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	defer wait.Done()

	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
/**
Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。

Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数
 */
func main() {
	runtime.GOMAXPROCS(1) //设置我的go程序中一个逻辑核心 m:n 设置 n为 1
	wait.Add(2)
	go a()
	go b()
	wait.Wait()
}
