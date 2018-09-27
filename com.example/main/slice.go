package main

import "fmt"

/**
切片
*/
func main() {

	arrays := [...]int{1, 5, 2, 9, 7}

	//创建索引从[1,3)的切片
	var b = arrays[1:3]
	fmt.Println(b)

	//创建一个数组，并返回切片的引用
	c := []int{6, 7, 8}
	fmt.Println(c)

	//使用make([]T,len,cap)函数创建切片
	d := make([]int, 5, 5)
	fmt.Println(d)

}
