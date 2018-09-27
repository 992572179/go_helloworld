package main

import "fmt"

func replace(s ...string) {
	s[0] = "GO"
	s = append(s, "playground")
	fmt.Println(s)
}
func main() {
	hello := []string{"hello", "world"}
	replace(hello...)
	fmt.Println(hello)
}
