package main

import "fmt"
/**
 结构体的属性 开头小写表示私有,只有当前包可以使用
开头大写表示公开访问,

 */
type student struct {
	name  string
	age   int
	hobdy []string
}

func main() {
	//结构体初始化
	//实例化结构体的第一种方法
	s := student{"李耀楚", 18, []string{"游泳", "学习"}}
	fmt.Println(s.age)

	//第二种方法
	s2 := student{}
	s2.age = 19
	s2.name = "liyaochu"
	s2.hobdy = []string{"piano"}
	fmt.Println(s2)

	//第三种方法   得到结构体的指针
	s3 := new(student)
	fmt.Println(s3)
	s3.age = 17
	s3.name = "river"
	s3.hobdy = []string{"piano"}
	fmt.Println(s3)

	//第四种方法    得到结构体的指针
	s4 := &student{}

	fmt.Println(s4)
	s4.age = 17
	s4.name = "river"
	s4.hobdy = []string{"piano"}

	//键值对初始化
	s5 := &student{"李耀楚", 18, []string{"游泳", "学习"}}

	fmt.Println(s5)

}
