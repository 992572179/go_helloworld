package main

import "fmt"

/**
GO中函数的声明及使用
*/
func main() {
	//addAndPrint(3,5)
	//a,b:=3,5
	//addAndPrint(a,b)

	l, w := 5.2, 3.1
	fmt.Print(rectPropsByReturn(l, w))

	fmt.Print("\n")

	//函数rectPropsByReturn为我们返回两个值，但可以只取一个值，另一个用"_"空白符代替
	area, _ := rectPropsByReturn(10.8, 5.6)
	fmt.Print(area)
}

func addAndPrint(a int, b int) {
	fmt.Print("两个数相加的和为：\t", a+b)
}

//多返回值函数
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

//在函数声明时指定area和perimeter为返回值，即函数最后一行可以直接return，自动返回
func rectPropsByReturn(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}
