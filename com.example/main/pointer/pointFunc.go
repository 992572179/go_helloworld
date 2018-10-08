package main

import "fmt"

/**
带指针的函数
*/
func main() {
	a := 25
	fmt.Println(&a)

	fmt.Println(modifyPointer(&a))
}

func modifyPointer(p *int) int {
	*p = 128
	return *p
}
