package main

import (
	"encoding/json"
	"fmt"
)

//json序列化
type Student struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Gender int `json:"gender"`
}

func main() {
	student := Student{
		ID:     "1",
		Name:   "李耀楚",
		Gender: 0,
	}
	v, err := json.Marshal(student)
	if err!=nil {
		fmt.Println("json序列化出错了")
		fmt.Println(err)
	}
	fmt.Println(v) //byte[]
	fmt.Printf("%#v",string(v))

	str:=string(v)
	s := &Student{}
	//
	json.Unmarshal([]byte(str),s)
	fmt.Println(s)
	fmt.Printf("%T",s)
}
