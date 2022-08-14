package main

import (
	"fmt"
	"time"
)

func main() {
	go foo()
	go bar()
	time.Sleep(time.Millisecond * 100)
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
