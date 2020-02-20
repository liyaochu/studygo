package main

import (
	"fmt"
)


func main() {
	ch := make(chan int, 1)

	for i:=0;i<10 ; i++ {
		select {
		case ch <-i : //尝试往 ch中些数据第一次就写进去了,第二次这里就写不了,需要走下面的case /
		case ret:=<-ch:
			fmt.Println(ret) //02468
			
			
		}
	}

}
