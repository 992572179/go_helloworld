package main

import "fmt"

/**
类型转换
*/
func main() {

	var num int64 = 21
	var balance float32 = 23.9

	//invalid operation: num + balance (mismatched types int and float32)
	//强数据类型，int不能和float相加求和
	//s:=num+balance
	//fmt.Print(s)

	//须将float32转化为int，再进行求和计算
	s := num + int64(balance)
	fmt.Print(s)

}
