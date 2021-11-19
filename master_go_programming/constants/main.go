package main

import "fmt"

const hour = 8

func main() {
	const days = 7

	x, y := 5, 0
	// fmt.Println(x / y) // error at runtime
	_, _ = x, y

	const a, b = 7, 0
	// fmt.Println(a / b) // error at compile

	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)

	fmt.Println(min1, min2, min3)

	const (
		val1 = -500
		val2
		val3
	)

	fmt.Println(val1, val2, val3)

	// Constant Rules
	// https://go.dev/blog/constants

	// constant can not get value com a variable
	// can not initiate with math.Pow

	// CONSTANTS RULES

	// 1. You cannot change a constant
	const temp int = 100
	// temp = 50 //compile-time error

	// 2. You cannot initiate a constant at runtime (constants belong to compile-time)
	// const power = math.Pow(2, 3) //error, functions calls belong to runtime

	// 3. You cannot use a variable to initialize a constant
	t := 5
	// error, variables belong to runtime and you cannot initialize a const to runtime values
	// const tc = t

	// 4. You can use a function like len() to initialize a const if it has as argument
	// a constant string literal (not a variable)
	// a string literal is an untyped constant

	const l1 = len("Hello") //ok

	str := "Hello"
	// const l2 = len(str) //error, str is a variable and belongs to runtime

	_, _ = t, str

	// UNTYPED CONSTANTS
	const x = 5
	const y float64 = 1.1

	var v1 int = 5
	var v2 float64 = 1.1

	fmt.Println(x * y)
	// => 5.5, No Error because x is untyped and gets its type when its used first time (float64).

	// fmt.Println(v1 * v2)
	// => Error: invalid operation: v1 * v2 (mismatched types int and float64)
	_, _ = v1, v2
}
