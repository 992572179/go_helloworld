package main

import "fmt"

//声明一个结构体
type Student struct {
	id int
	//name string
	//address string
	name, address string
}

//声明一个匿名结构体，使用var关键字，而不是type
var Department struct {
	deptName, manageBy string
}

func main() {

	//使用属性名创建一个结构体变量
	stu1 := Student{
		id:      1001,
		name:    "changSan",
		address: "xi'an",
	}
	//不使用属性创建，但参数需要对应
	stu2 := Student{1002, "liSi", "shanghai"}

	fmt.Println(stu1)
	fmt.Println(stu2)
	fmt.Println(stu2.address)

	var stu3 Student
	stu3.id = 1003
	stu3.name = "tokyo"
	fmt.Println("stu3:", stu3)
}
