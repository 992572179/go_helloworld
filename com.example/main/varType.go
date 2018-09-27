package main

import (
	"fmt"
	"unsafe"
)

/**
变量的类型
*/
func main() {

	var num int
	balance := 1800000000001111111
	fmt.Printf("type of num is %T, size of a is %d", num, unsafe.Sizeof(num))
	fmt.Printf("\ntype of balance is %T, size of b is %d", balance, unsafe.Sizeof(balance))

}
