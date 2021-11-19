package main

import "fmt"

type km float64
type mile float64

func main() {
	type speed uint

	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1)

	var parisToLandom km = 465
	var distanceInMiles mile

	// distanceInMiles = parisToLandom / 0.621  vai dar erro pois sao tipos diferentes
	distanceInMiles = mile(parisToLandom) / 0.621

	fmt.Println(distanceInMiles)
}
