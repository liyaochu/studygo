package main

import (
	"fmt"
	"reflect"
)
//通过反射修改值
func update(x interface{}){
	v:= reflect.ValueOf(x)
	kind := v.Kind()
	if kind==reflect.Ptr {
		//取到对应的指针再去修改值
		v.Elem().SetInt(111)
	}

}

func main() {
   var a int64 = 100
	update(&a)
	fmt.Println(a)


}
