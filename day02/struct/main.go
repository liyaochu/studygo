package main

import "fmt"
//自定义类型 是一个新类型
type NewInt int
//类型别名
type MyInt=int
func main() {
	var a NewInt
	fmt.Println(a)
	fmt.Printf("%T\n",a) //main.NewInt

	var b MyInt

	fmt.Println(b)
	fmt.Printf("%T\n",b) //int
}
