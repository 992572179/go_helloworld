package main

import "fmt"

//可变参数
func find(num int, nums ...int) {
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "索引为：", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println("未找到.")
	}
}

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
}
