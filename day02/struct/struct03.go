package main

import "fmt"

type car struct {
	name string
	age int
}

//自己实现一个构造函数
func newCar(name string,age int) *car{
	return  &car{name,age}
}
func main() {
	c := newCar("保时捷", 25)
	fmt.Println(c)
}
