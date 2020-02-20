package main

import "fmt"
//指针的初始化
/**
  new 用来初始化值类型指针
  make 用来 初始化 slice map chan
 */
func main() {
	//var a *int   //a 是int类型的指针,但是这个写法是不正确,没有初始化指针
	var a=new(int)  //这种写法才对
	*a = 10
	fmt.Println(a)
	fmt.Println(*a)// invalid memory address or nil pointer dereference

    var c =new([3]int)
	c[0]=1
    fmt.Println(*c) // [1 0 0]
}
