package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter,r*http.Request){
        fmt.Fprintln(w,"hello go")
}
func main() {
	http.HandleFunc("/",sayHello)            //注册漏油,当你访问就注册sayHello函数
	err := http.ListenAndServe(":9090", nil) //建立连接
	if err!=nil{
		fmt.Printf("server failed %v\n",err)
		return
	}

}
