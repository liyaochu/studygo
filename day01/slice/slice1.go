package main

import "fmt"

func main() {
	//切片是引用类型
	a:=[]int{1,2,3}
	fmt.Println(a,len(a),cap(a))
	//var c []int
    c:=make([]int,3,3)
	b:=a
	copy(c,a)
	b[0]=100
	fmt.Println(a)//[100 2 3]
    fmt.Println(c) //[1 2 3]

	fmt.Println(b)//[100 2 3]
}
