package main

import (
	"fmt"
	"io"
	"io/ioutil"
)
//ioutil读文件
func main() {
	file, err := ioutil.ReadFile("./day03/x.txt")
	if err==io.EOF{
		fmt.Println("文件读完了")
		return
	}
	if err!=nil {
		fmt.Println("read form failed,err",err)
	}
	fmt.Println(string(file))
}
