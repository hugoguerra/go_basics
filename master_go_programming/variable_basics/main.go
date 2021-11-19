package main

import (
	"fmt"
)

// y := 19 -> non-declaration statement outside function body | nao pode declarar variables fora da funcao

func main() {
	var age int = 10
	var ageV2 = 20
	fmt.Println("Age is: ", age)
	fmt.Println("Age V2 is: ", ageV2)

	var name = "Dan"
	fmt.Println("Name is: ", name)

	// we use short declaration(:= , =) when we know the initial value
	s := "Learning Golang"
	fmt.Println(s)

	car, cost := "Audi", 5000
	fmt.Println(car, cost)

	car, year := "BMW", 1990
	_ = year

	var opened = false
	opened, file := true, "text.txt"
	_, _ = opened, file

	// we use multiple declaration for better readability
	var (
		salary   float64
		firsName string = "Oi"
		gender   bool
	)

	fmt.Println(salary, firsName, gender)

	var a, b, c int

	// multiple assignments
	a, b, c = 1, 2, 3
	println(a, b, c)

	// expression in a short declaration
	sum := 2 + 5.5
	fmt.Println(sum)

	// x := 10 -> variaveis locais nao podem deixar sem uso, usar o _

}
