package main

import "fmt"

type Haier struct {
	name string
	price float64
}
type xiaotiane struct {
	name string
	price float64
}
type xiyiji interface {
	wash()
}
func (h Haier) wash(){
	fmt.Println("海尔洗衣机洗衣服")
}
func main() {
	var a xiyiji
	haier := Haier{
		name:  "海尔洗衣机",
		price: 1500.45,
	}
	a=haier
	fmt.Println(a)  //{海尔洗衣机 1500.45}

	fmt.Println(haier)  //{海尔洗衣机 1500.45}

}
