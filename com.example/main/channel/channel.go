package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("信道为空，需要定义赋值。")
		a = make(chan int)
		fmt.Printf("Type of a is %T:\n", a)
	}

	fmt.Println(".....")
	//简短声明、赋初始值
	b := make(chan float64)
	fmt.Println(b)
}
