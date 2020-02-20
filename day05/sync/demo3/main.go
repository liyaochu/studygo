package main

import (
	"fmt"
	"sync"
)

var onlyonce sync.Once

func f1(a int){
	fmt.Println(a)
}
func close(x int ) func(){
	return func() {
		f1(x)
	}
}
//利用闭包完成 onlyonce.Do(f)  f参数必须是空函数
func main() {
	f := close(10)
	onlyonce.Do(f)
}
