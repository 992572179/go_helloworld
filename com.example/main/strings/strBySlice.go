package main

import (
	"fmt"
	"unicode/utf8"
)

//返回字符串的长度
func getStrLen(s string) {
	fmt.Printf("length of this str is:%d\n", utf8.RuneCountInString(s))
}

func main() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	byteSlice2 := []byte{67, 97, 102, 195, 169}
	str := string(byteSlice)
	fmt.Println(str)
	str = string(byteSlice2)
	fmt.Println(str)

	defStr := "hello, world"
	getStrLen(defStr)
	fmt.Println(len(defStr))
}
