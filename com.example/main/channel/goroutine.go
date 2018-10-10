package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello() going to sleep for 2 sec")
	time.Sleep(2 * time.Second)
	fmt.Println("hello goroutine...")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("waiting hello() exec...")
	go hello(done)
	<-done
	fmt.Println("main ....")

}
