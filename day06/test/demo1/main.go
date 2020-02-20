package main

import "strings"
//用sep 切割 s
func Split(s, sep string) (result[]string ){

	index := strings.Index(s, sep)
	for index >0{
		result = append(result, s[:index])
		s=s[index+1:]
		index=strings.Index(s,sep)
	}
	result = append(result, s)
	return
}

func main() {

}
