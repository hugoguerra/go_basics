package main

import (
	"fmt"
	"math"
)

var version = "3.1"

func main() {
	// var a float64 = 7.1

	// x, y := true, 3.7
	// a, x = 5.5, false

	// _, _, _ = a, x, y

	// name := "Golang"
	// fmt.Println(name)

	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)

	const secPerDay = 24 * 60 * 60
	const daysYear = 365

	fmt.Printf("There are %v seconds in a yeader", secPerDay*daysYear)

	const x1 int = 10

	// declaring a constant of type slice int ([]int)
	// const m = []int{1: 3, 4: 5, 6: 8} -> -> You cannot declare a slice constant.
	// _ = m

	const a int = 7
	const b float64 = 5.6
	const c = float64(a) * b

	x := 8
	// const xc int = x
	_ = x

	// const noIPv6 = math.Pow(2, 128) -> nao pode passar o valor de uma variable para um constant
	var noIPv6 = math.Pow(2, 128)
	_ = noIPv6
}
