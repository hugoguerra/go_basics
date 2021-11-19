package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string

	_, _, _, _ = a, b, c, d

	var x = 20
	var y = 15.5
	var z = "Gopher!"

	_, _, _ = x, y, z

	fmt.Println(a, b, c, d, x, y, z)
}
