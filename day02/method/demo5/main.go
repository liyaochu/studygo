package main

import "fmt"

//结构体模拟继承
type animal struct {
	name string

}
//定义一个动物会吃的方法
func (a*animal)eat(){
fmt.Printf("%s 在犬吠",a.name)
}

type dog struct {
	feet int
	* animal
}
func main() {
	d := dog{
		feet:   4,
		animal: &animal{name: "旺财"},
	}
	//狗也是动物所以也可以吃
	d.eat()
}
