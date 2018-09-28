package main

import "fmt"

func main() {
	//使用make函数创建一个map，并向其中放入2组数据
	jsonMap := make(map[int]string)

	jsonMap[1] = "emp_id"
	jsonMap[100] = "salary"
	fmt.Println(jsonMap)

	//简短声明一个map，向其中立即放入2组数据。ps.注意格式
	student := map[string]string{
		"ChangSan": "guangzhou",
		"liSi":     "beijing",
	}
	student["wangWu"] = "shanghai"
	fmt.Println(student)

	//按照指定的key在map中获取value
	name := "wangWu" //name:=student["wangWu"]
	fmt.Println("key:", name, "value:", student[name])

	inputStu := "wangWu"
	value, flag := student[inputStu]
	if flag == true {
		fmt.Println("从map中获取到了，值为", value)
	} else {
		fmt.Println("map中未找到：", inputStu)
	}

	//遍历map
	for k, v := range student {
		fmt.Println("key:", k, "value:", v)
	}

	//删除map中的元素，使用delete(m map[Type]Type1, key Type)
	delete(student, "ChangSan")
	fmt.Println()
	for k, v := range student {
		fmt.Println("key:", k, "value:", v)
	}

	/**
	invalid operation: newMap == student (map can only be compared to nil
	判断map是否相同只能遍历比较，"=="操作符只能判空
	map == nil
	*/
	/*newMap:=map[string]string{}
	if newMap == student{

	}*/

	if student != nil {
		fmt.Println("not null!")
	}
}
