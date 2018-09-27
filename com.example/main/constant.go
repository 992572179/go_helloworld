package main

import (
	"fmt"
	"math"
)

/**
常量的声明
*/
func main() {

	//声明常量用const关键字，且声明后须赋初始值
	const uName = "guest"
	fmt.Println(uName)

	//const balance
	//fmt.Print(balance)

	f := math.Sqrt(4)
	fmt.Println(f)
}
