package main

import (
	"fmt"
	"strings"
)

//字符串反转
func main(){
	s1:="hello"
	bytesArray := []byte(s1)
	s2:=""
	for i:=len(bytesArray)-1;i>=0 ;i--  {
		s2=s2+string(bytesArray[i])
	}
	fmt.Println(s2)

	split := strings.Split(s1, "")
	fmt.Println(split)
	a:=[]string{"pythpn","java","go"}
	join := strings.Join(a, "~")
	fmt.Println(join)

}
