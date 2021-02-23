package main

import (
	"fmt"
)

func main() {
	loc := [2]int{-15, 16}
	for i := loc[0]; i < loc[1]; i++ {
		fmt.Printf("%3d: %3d\n", i, mod(i, -7))
	}
}

// https://en.wikipedia.org/wiki/Modular_arithmetic

func mod(a, b int) int {

	m := a % b

	if (m < 0) && (b > 0) {
		m += b
	}

	if (m > 0) && (b < 0) {
		m += b
	}

	return m
}
