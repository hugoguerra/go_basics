package main

import "fmt"

func main() {
	var a = 10
	var b = 3.3
	fmt.Println(a, b)

	a = int(b)
	fmt.Println(a, b)

	var (
		value int
		price float64
		name  string
		done  bool
	)
	// int initialized with 0
	// float64 initialized with 0
	// string initialized with empty
	// bool initialized with false

	fmt.Println(value, price, name, done)

	x, y := 2., 5
	x = y
	fmt.Println(x, y)
}
