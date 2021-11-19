package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	for i := 1; i < 3; i++ {
		fmt.Println("f1() i=", i)
		time.Sleep(time.Second)
	}

	fmt.Println("Finish f1() execution")
	wg.Done()
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

	var wg sync.WaitGroup

	wg.Add(1)

	go f1(&wg)
	fmt.Println("No. of Goroutines after f1():", runtime.NumGoroutine())

	f2()

	wg.Wait()
	fmt.Println("Finish main execution")
}
