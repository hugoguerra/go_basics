package main

import "fmt"

type second uint
type duration second

type authorizationDomain struct {
}

func main() {
	var x int
	y := 10
	fmt.Println(20 < y || x == 0)

	var a = fmt.Sprintf("%d\n", 34234)
	fmt.Sprintf("%T\n", a)

	var d duration
	fmt.Printf("%T\n", d)

	v := 20
	fmt.Println(v)
	goto todo

todo:
	fmt.Println("Estou no todo")
}
