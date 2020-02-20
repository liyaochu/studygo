package main

import "fmt"

//匿名字段
type people struct {
	name string
	age int
	int
	string
}

func main() {
	p := people{name:"李耀楚", age:18}
   fmt.Println(p.int)
}
