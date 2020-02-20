package main

import "fmt"
//从终端输入可以可以获取到
func main() {
	var (
		name string
		age int
	)
	//fmt.Scan(&name,&age)
	//fmt.Scanf("name:%s age:%d",&name,&age)
	fmt.Scanln(&name,&age)
	fmt.Println(name,age)
}
