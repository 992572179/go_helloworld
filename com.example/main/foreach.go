package main

import "fmt"

/**
for循环的使用
*/
func main() {

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
		if i >= 6 {
			//满足条件时，break之后的语句不会被执行，且跳出循环。
			break
		}
	}

	fmt.Println()

	for j := 0; j <= 10; j++ {
		if j%2 == 0 {
			//满足条件时，continue之后的语句不会被执行，而是进入下一次循环。
			continue
		}
		fmt.Printf("%d ", j)
	}
}
