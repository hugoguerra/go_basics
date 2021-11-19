package main

import (
	"fmt"
	"sync"
)

func printParameter(s string, wgp *sync.WaitGroup) {
	fmt.Println(s)
	wgp.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go printParameter("This is my test", &wg)

	wg.Wait()
}
