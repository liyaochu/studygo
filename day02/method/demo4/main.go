package main

import "fmt"

//结构体的嵌套
type address struct {
	province string
	city string
}
type emil struct {
	province string
}
type student struct {
	name string
	age int
	addr  address //嵌套了别的结构体
	em emil
}
func main() {
	s := student{
		name: "李耀楚",
		age:  18,
		addr: address{"上海", "松江"},
		em:emil{"上海"},
	}
	//fmt.Println(s.province) 匿名字段支持简写
	fmt.Println(s.em.province)  //当匿名字段有冲突时,必须显示调用
	fmt.Println(s.addr.province)
}
