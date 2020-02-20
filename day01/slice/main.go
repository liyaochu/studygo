package main

import "fmt"

func main() {

	//声明切片的方式第一种:直接声明
	s:=[]int{1,2,3}
	fmt.Println(s)
	//从数组得到切片
    a:=[3]int{1,2,3}
    var c []int
    c =a[0:2]
    fmt.Printf("c:%T",c)

    x:=[...]string{"上海","广州","北京","南京","苏州","长沙"}
    y:=x[1:4] //不包含4的索引[广州 北京 南京]
    fmt.Println(y)
    //切片y的长度 是3
    fmt.Println(len(y))
    //切片的容量 5 因为是从广州切的 就从底层数组最大放多少容量 广州->长沙
	fmt.Println(cap(y))
}
