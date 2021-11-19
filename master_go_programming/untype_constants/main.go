package main

import "fmt"

func main() {
	const a float64 = 3.5 // typed constant

	const b = 5.7 // untyped constant

	const c float64 = a * b

	const str = "Hello" + "Go!"

	const d = 5 > 6

	fmt.Println(d)

	// const x int = 5
	// const y float64 = 2.2 * x // error compile int <> float | typed constants

	// untyped constant not get error at differents types
	const x = 5
	const y = 2.2 * x
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", x)

}
