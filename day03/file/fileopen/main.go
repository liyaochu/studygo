package main

import (
	"fmt"
	"io"
	"os"
)

//打开和关闭
func main() {
	file, err := os.Open("./day03/x.txt")
	if err!=nil{
	  fmt.Println("open file failed",err)
		return
	}
		defer file.Close()

	//读文件
    var bytes [128]byte //定义一个字节数组
	for {
		read, err := file.Read(bytes[:])  //基于数组得到一个切片
		if err==io.EOF{
			fmt.Println("文件读完了")
			return
		}
		if err!=nil {
			fmt.Println("read form failed,err",err)
		}
		fmt.Printf("读取了%d个字节",read)
		fmt.Println(string(bytes[:]))
	}

}
