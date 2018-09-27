package main

import "fmt"

func main() {
	num := 11

	if num%2 == 0 {
		fmt.Print("num 是偶数。")
	} else {
		fmt.Println("num 是奇数。")
	}

	//if statement; condition.注意：需要加";"
	if a := 10; a%2 == 0 {
		fmt.Print("a 是偶数。")
	} else {
		fmt.Print("a 是奇数。")
	}
}
