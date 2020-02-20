package main

import (
	"fmt"
	"time"
)

func printTime(t time.Time){
	//"2006-01-02 15:04:05"
	format := t.Format("2006-01-02 15:04:05")
	fmt.Println(format)
}
//求两个时间的差值一种方式
func calcTime(){
	starttime := time.Now().UnixNano() / 1000
	fmt.Println("红酥手 黄藤酒 ")
	time.Sleep(time.Millisecond*30)
	endtime := time.Now().UnixNano() / 1000
	fmt.Println(starttime-endtime)
}
//求两个时间的差值二种方式
func calcTime2(){
	starttime := time.Now()
	fmt.Println("红酥手 黄藤酒 ")
	time.Sleep(time.Millisecond*30)

	fmt.Println("耗费了",time.Since(starttime))
}
func main() {
	now := time.Now()
	printTime(now)
	calcTime2()
}
