package main

import "fmt"
//switch 的坑
func main() {
	f := func() bool {
		return false
	}
	// switch f()
	switch f(); {
	case true:
		fmt.Println("真")
	case false:
		fmt.Println("假")
	}

	//switch 空 默认返回真
}