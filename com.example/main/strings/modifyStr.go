package main

import "fmt"

/**
一个字符串一旦声明就不能做更改
可将其转化为rune切片，操作切片，再转回字符串
*/
func transferStr(s []rune) string {
	s[0] = 'k'
	s[1] = 'o'
	return string(s)
}
func main() {
	s := "hello"
	fmt.Println(transferStr([]rune(s)))
}
