package main

import (
	"fmt"
)

func printChar(s string) {
	/*
		/rune是GO语言的内建类型，一个rune代表一个代码点
		无论占用多少字节，都可以用一个rune来表示
	*/
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

//打印字符对应的AscII码
func printAscII(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d ", s[i])
	}
}

//打印字符对应的16进制数
func printHex(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

/**
for range遍历字符串
*/
func forRangeStr(s string) {
	for _, rune := range s {
		fmt.Printf("%c", rune)
	}

}

func main() {

	s := "Васил"
	printChar(s)
	fmt.Println()
	printAscII(s)
	fmt.Println()
	printHex(s)

	fmt.Println()
	name := "Señor"
	forRangeStr(name)
}
