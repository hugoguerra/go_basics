package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	for i := 1; i < 3; i++ {
		fmt.Println("f1() i=", i)
	}

	fmt.Println("Finish f1() execution")
}

func f2() {
	for i := 4; i < 8; i++ {
		fmt.Println("f2() i=", i)
	}

	fmt.Println("Finish f2() execution")
}

func main() {

	fmt.Println("Init main execution")
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	go f1()
	fmt.Println("No. of Goroutines after f1():", runtime.NumGoroutine())

	f2()

	time.Sleep(time.Second * 2)
	fmt.Println("Finish main execution")
}
