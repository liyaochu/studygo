package main

import "fmt"



type MyInt int
//可以给任意类型追加方法
func(m * MyInt) sayHi(){
	fmt.Println("hello")
}
//不能给别的包定义的类型追加方法
func(a*int) sayHello(){

}

func main() {
	var a MyInt

	a.sayHi()


}
