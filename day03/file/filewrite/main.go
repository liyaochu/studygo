package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//打开文件支持文件写入
func write(){
	ioutil.WriteFile("x.txt",[]byte("学习go"),1)
}

func main() {
	file, err := os.OpenFile("x.txt", os.O_CREATE|os.O_WRONLY, 0755)
	if err!=nil {
		fmt.Println("open failed,err",err)
		return
	}
	defer  file.Close()

	str:="李耀楚 爱 张瑞文"
	file.Write([]byte("哈哈哈哈\n"))
	file.WriteString(str)

	write()
}
