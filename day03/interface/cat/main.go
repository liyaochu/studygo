package main

//为什么要用看接口
type Cat struct {
}

func (c Cat) Say() string {
	return "喵喵"
}

type Dog struct {
}

func (d Dog) Say() string {
	return "旺旺"
}

type animal interface {
	Say() string
}

func main() {
	//cat := Cat{}
	//fmt.Println("猫", cat.Say())
	//dog := Dog{}
	//fmt.Println("狗", dog.Say())


}
