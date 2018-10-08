package main

import "fmt"

/**
指针对象的定义和使用
*/
func main() {
	num := 100.00
	//p为num的内存地址
	var p = &num
	fmt.Printf("%T\n", num)
	fmt.Println(p)

	a := "str"
	b := &a
	//对指针类型的"解引用"
	c := *b
	fmt.Println(b, "\t", c)
}
