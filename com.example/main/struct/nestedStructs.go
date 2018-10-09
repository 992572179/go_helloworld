package main

import "fmt"

type Person struct {
	id   int
	name string
	//City是Person中的一个匿名字段
	City
}

type City struct {
	cityName string
	aqi      int64
}

//嵌套结构体
func main() {
	var p Person
	p.name = "Spark"
	p.id = 1004
	//可直接调用，cityName为"提升字段"
	p.cityName = "guangzhou"
	p.aqi = 83
	fmt.Println(p)
}
