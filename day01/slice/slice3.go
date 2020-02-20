package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	b := a[:]
	b[0] = 100
	fmt.Println(a[0]) //100
	c := a[2:5]
	fmt.Println(c) //[3 4 5]
	fmt.Println(len(c))// 3
	fmt.Println(cap(c))// 5    3, 4, 5, 6, 7 从3开始数到最后

	d:=c[2:5]
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))

	/**
	 &:求内存地址
	 *:取内存对应的值
	 */
	A:=[3]int{1,2,3}
	ModifyArray(&A)
	fmt.Println(A)
}

func ModifyArray(a1*[3]int){
	a1[0]=100
}
