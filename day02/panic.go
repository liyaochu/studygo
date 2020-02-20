package main

import "fmt"

func main() {

	defer func() {
		err := recover() //尝试将函数从当前异常恢复过来
		fmt.Println("err==",err)
	}()
	 var a [] int
	 a[0]=100
	 /**
	  这个上面写法不对,没有初始化slice
	  */
	 fmt.Println("这是main函数")
}
