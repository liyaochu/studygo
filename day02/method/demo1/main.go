package main

import "fmt"

/**
 1.方法就是某一个具体的类型才能调用的函数
2.函数是谁都可以调用

 */
type people struct {
	name string
	age int
}
//这个势函数
//func dream(){
//	fmt.Println("我的梦想是不想上班")
//}

//这个是方法
//函数指定接收者之后是方法
//func (p people)dream(){
//	p.age=20
//	fmt.Printf("%s的梦想是不想上班",p.name)
//}
//指针可以修改数据
func (p *people)dream(){
	p.age=20
	fmt.Printf("%s的梦想是不想上班",p.name)
}
func main() {
	//调用方法首先要实例化接收者,然后通过接收者调用函数
	p := people{"李耀楚", 18}
	p.dream()
	fmt.Println(p.age)
}
