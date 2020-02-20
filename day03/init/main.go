package main

import "fmt"

//init 示例
//先做全局的声明,然后 init
var today="星期四"

const week  = 7

type student struct {
	Name string
}
//包导入的时候会自动执行,(多用来做初始化执行)
func init(){
	fmt.Println("包被初始化")
	fmt.Println(week)
}
func main() {
	fmt.Println("这是面函数")
}
