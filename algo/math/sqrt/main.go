package main

import (
	"fmt"
	//"math"
)

// sqrt - square root of x.
func sqrt(a float64) float64 {
	x := a / 2
	for i := 0; i < 30; i++ {
		x = (x + a/x) / 2
	}
	return x
}

// cbrt - cube root of x.
func cbrt(a float64) float64 {
	x := 1.0
	for i := 0; i < 30; i++ {
		y := sqrt(a / x)
		x = (x + y + y) / 3
	}
	return x
}

func main() {

	// fmt.Println(sqrt(2))
	fmt.Println(sqrt(9))

	// fmt.Println(math.Pow(5, 1.0/3.0))
	fmt.Println(cbrt(27))
}
