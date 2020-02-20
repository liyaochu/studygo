package main

import "fmt"

/**
defer 时机     1.返回值
              2.运行defer
              3.运行return
 */
func f1() int{
	x:=5
	defer func() {
		x++
	}()
	return x //1.返回值=5 2.x++ 3.ret 5==>5
}
func f2() (x int){

	defer func() {
		x++
	}()
	return 5   //1.返回值=x 2.x++ 3.ret==>6
}
func f3() (y int){
	x:=5
	defer func() {
		x++
	}()
	return x   //1.返回值=y 2.x++ 3.ret==>5
}

func f4() (x int){
	defer func(x int) {
		x++
	}(x)
	return 5   //1.返回值=x 2.x++(函数内部的x和外面的x不是一个变量) 3.ret==>5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

}
