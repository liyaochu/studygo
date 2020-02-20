package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./day03/x.txt")
	if err!=nil{
		fmt.Println("open file failed",err)
		return
	}
	defer file.Close()

	//利用缓冲区从文件读取数据
	reader := bufio.NewReader(file)
	for{
		rea, err := reader.ReadString('\n')
		//io.EOF表示文件读完了
		if err==io.EOF{
			fmt.Println("文件读完了")
			return
		}
		if err!=nil{
			fmt.Println("open file failed",err)
			return
		}
		fmt.Println(rea)
	}

}
