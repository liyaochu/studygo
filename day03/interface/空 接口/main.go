package main

import (
	"fmt"
	"time"
)
func showType(i interface{}){
fmt.Printf("%T\n",i)
}



//空接口什么类型都可以
func main() {
	 var a interface{}
	 a=8
	 fmt.Println(a)
	 a="hahha"
	 fmt.Println(a)
	 a=false
	 fmt.Println(a)
	 a= struct {
		 name string
	 }{"huahua"}
	fmt.Println(a)

	showType(1)
	showType("哈哈")
	showType(time.Now())

	//定义一个值为空的interface的map
	m := make(map[string]interface{},  100)
	m["李耀楚"]=100
	m["瑞文"]=true
	m["老葛"]=99.99

}
