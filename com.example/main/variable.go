package main

import "fmt"

/**
变量的声明
*/
func main() {
	//var age = 18
	//fmt.Print("age is :" ,age)

	//声明多个变量
	var id, name, address = 1001, "z3", "guangzhou"
	fmt.Println(id, name, address)

	//简短声明，但是要求：操作符左边的所有变量都有初始值
	className, stuNo := "A class", 1002
	fmt.Println(className, stuNo)
	fmt.Println(len(className))

	//简短声明但不赋予初始值，运行报错
	//assignment count mismatch: 2 = 1
	//height,fruit := 185
	//fmt.Print(height,fruit)
}
