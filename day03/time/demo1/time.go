package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//fmt.Println(now)
	fmt.Printf("%#v\n",now)
	now.Year()
	now.Month()
	now.Day()
	now.Second()
	now.Minute()
	now.Nanosecond()

	//时间戳
	now.Unix() //时间戳
	now.UnixNano() //纳秒
	timestamp(111111111111)
	//定时器
	//tickDemo()

	//日期格式化
	//时间格式数据=>>字符串数据
	formatDemo()
}

func timestamp(timestamp int64){
	unix := time.Unix(timestamp, 0)  //将时间戳转化成时间格式
	fmt.Println(unix)
}

func tickDemo(){
	tick := time.Tick(time.Second) //定义一个间隔一秒钟的定时器
	for i:=range tick{
		fmt.Println(i) //每秒都会执行
	}
}
//格式化模板
//其他语言一般是yyyy-mmm-dd,但是go是2006-01-02 15:04:05
func formatDemo(){
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006-01-02"))

}