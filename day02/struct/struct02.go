package main

import "fmt"

type stu struct {
	name string
}
func main() {
	s := stu{"李耀楚"}
	s1:=s    //s copy了一份给 s1
	s1.name="张瑞文"
	fmt.Println(s1.name)  //张瑞文
	fmt.Println(s.name)  //李耀楚

	s2:=&s  //将s的地址赋值给了s2
	s2.name="娜扎"
	fmt.Println(s.name)  //娜扎
	fmt.Println(s2.name) //娜扎
}
