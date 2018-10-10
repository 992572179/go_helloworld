package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("goroutine...")
}

func main() {
	//在函数或者方法前面，加上"go"关键字
	go hello()
	//主协程休眠一秒
	time.Sleep(1 * time.Second)
	fmt.Println("main ...")
}
