package main

import "fmt"

/**
数组的遍历、引用
*/
func main() {

	nums := [...]int{9, 0, 3, 5, 2, 7}
	ints := nums
	ints[0] = 2
	//数组是值类型，当原数组赋值给新数组，再对新数组操作，不会改变原数组的属性
	//fmt.Println(nums)
	//fmt.Println(ints)

	fmt.Println("数组nums的大小为： ", len(nums))

	//for i := 0; i < len(nums); i++ {
	//	fmt.Println(nums[i])
	//}

	a := [...]float64{21, 78, 90.12, 88.23}

	//使用range函数遍历
	//for s, r := range a {
	//	fmt.Printf("%d  %.2f", s, r)
	//	fmt.Println()
	//}

	//忽略索引，使用"_"
	for _, j := range a {
		fmt.Println(j)
	}

}
