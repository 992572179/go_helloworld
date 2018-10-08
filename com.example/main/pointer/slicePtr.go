package main

import "fmt"

func main() {
	a := [...]int{33, 44, 55}
	modify(&a)
	fmt.Println(a)
	modify2(a[:])
	fmt.Println(a)
}

func modify(arr *[3]int) {
	//(*arr)[0] = 90
	//(*a)[x]-->a[x]
	arr[0] = 90
}

/**
使用切片代替指针数组
*/
func modify2(sli []int) {
	sli[0] = 99
}
