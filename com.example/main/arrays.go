package main

import "fmt"

/**
不同风格的数组声明、赋值
*/
func main() {

	//数组的声明、赋值
	var nums [5]int
	nums[0] = 1
	nums[2] = 5
	fmt.Print(nums)
	fmt.Println()

	//小数默认科学记数法打印输出
	var coins [5]float64
	coins[3] = 0.000014
	coins[0] = 0.00000000001
	fmt.Print(coins)
	fmt.Println()

	//简短声明
	trees := [10]int{1, 3, 5, 7}
	fmt.Print(trees)
	fmt.Println()

	//用"..."忽略数组长度(容量)，让编译器计算长度
	clients := [...]string{"z3", "l4", "w5"}
	fmt.Print("数组容量为： ", cap(clients))
}
