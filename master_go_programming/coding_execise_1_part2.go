package main

import "fmt"

func main() {
	var (
		a int
		b float64
		c bool
		d string
	)

	_, _, _, _ = a, b, c, d

	var (
		x = 20
		y = 15.5
		z = "Gopher!"
	)
	_, _, _ = x, y, z

	fmt.Println(a, b, c, d, x, y, z)
}
