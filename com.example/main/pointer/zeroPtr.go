package main

import "fmt"

/**
变量的初始值以及为其赋值
*/
func main() {
	a := 25
	var b *int
	if b == nil {
		//未赋值，打印为nil
		fmt.Println("b is", b)
		b = &a
		fmt.Println("after initial:", b)
	}
}
