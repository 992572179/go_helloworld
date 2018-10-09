package main

import "fmt"

//字段可以只有类型，没有字段名，称为匿名字段
type Hyper struct {
	int
	string
}

func main() {

	h1 := Hyper{1001, "Jerry"}
	fmt.Println("h1:", h1)

	var h2 Hyper
	//可以直接访问结构体的匿名字段，并赋值
	h2.string = "Tom"
	h2.int = 1002
	fmt.Println("h2:", h2)

}
