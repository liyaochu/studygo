package main

import "fmt"

//实现接口时,使用值接收者和指针接收者的区别

type Cat struct {
	name string
}
type animal interface {
	eat()
	move()
}

//cat实现 animal接口  (值接收)
/*func (c Cat) eat() {
	fmt.Println("猫吃鱼")
}
func (c Cat) move() {
	fmt.Println("猫在跑")
}*/
////cat实现 animal接口  (指针接收)
func (c *Cat)eat(){
//
}
func (c *Cat) move() {
	fmt.Println("猫在跑")
}
func main() {
	var a animal
	//cat := Cat{"汤姆"}
	//a = cat   //实现animal接口的是*cat 类型,此时cat是值类型,不能直接存值类型
	//获取一个指针
	tom:=&Cat{name:"杰克"} //*cat
	a=tom
	tom.eat()  //(*tom).eat
	tom.move() //(*tom).move
	fmt.Println(a)
}
