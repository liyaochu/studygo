package main

import (
	"fmt"
	"reflect"
)
//通过反射获取到类型
func reflectType(x interface{}){
	of := reflect.TypeOf(x)
	fmt.Printf("type %v\n",of)
}

func main() {

	reflectType(100)
	reflectType("上海有海")
	reflectType(map[string]int{})
}
