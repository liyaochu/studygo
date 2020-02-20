package main

import "fmt"

//类型断言 x.(int)
//type nullinterface interface {}  写空接口这样写很繁琐,不需要用type去定义直接可以定义一个

func showType(x interface{}){
	//我们用这个函数可以接收任意变量
	v ,ok:= x.(int)
	if !ok {

	}else {
     fmt.Println("这是一个int类型",v)
	}
	v2 ,ok:= x.(string)
	if !ok {
		fmt.Println("不是string")
	}else {
		fmt.Println("这是一个int类型",v2)
	}
}

func showType1(x interface{}){
	//我们用这个函数可以接收任意变量
	switch x.(type) {
	case string:
		fmt.Println("这是一个int类型")
	case int:
		fmt.Println("这是一个int类型")
	case byte:
		fmt.Println("这是一个int类型")
	case bool:
		fmt.Println("这是一个int类型")
		
	}
}
func main() {
	//var a interface{}
	//a=100
	showType(100)
	showType1(100)
}
