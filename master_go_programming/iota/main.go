package main

import "fmt"

func main() {
	const (
		val1 = iota
		val2 = iota
		val3 = iota
	)

	fmt.Println(val1, val2, val3)

	const (
		val11 = iota
		val22
		val33
	)

	fmt.Println(val11, val22, val33)

	const (
		East = iota
		North
		South
		West
	)

	fmt.Println(West)

	const (
		a = iota * 2
		b
		c
	)

	fmt.Println(a, b, c)

	const (
		x = -(iota + 2)
		_
		y
		z
	)

	fmt.Println(x, y, z)
}
