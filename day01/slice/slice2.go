package main

import "fmt"

func main() {
	//切片怎么删除元素 删除上海,没有直接的delete
	a:=[]string{"北京","上海","杭州","深圳"}
	strings := append(a[:1], a[2:]...)
	fmt.Println(strings)

}
