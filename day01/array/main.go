package main

import "fmt"

func main() {
	a:=[5]int{1,3,5,7,8}
	num:=0
	for i:=0;i<len(a) ; i++ {
		num+=a[i]
	}
	fmt.Println(num)
}
