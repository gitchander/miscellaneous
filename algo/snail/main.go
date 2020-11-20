package main

import (
	"fmt"
)

func makeSnail(xn, yn int) [][]int {

	const (
		Right = iota
		Down
		Left
		Up
	)

	aas := make([][]int, yn)
	for i := range aas {
		aas[i] = make([]int, xn)
	}

	var (
		minX = 0
		maxX = xn
	)

	var (
		minY = 0
		maxY = yn
	)

	var (
		x   = 0
		y   = 0
		dir = Right
	)

	i := 0
	for (minX < maxX) && (minY < maxY) {

		aas[y][x] = i
		i++

		switch dir {
		case Right:
			if x+1 < maxX {
				x++
			} else {
				y++
				minY++
				dir = Down
			}
		case Down:
			if y+1 < maxY {
				y++
			} else {
				x--
				maxX--
				dir = Left
			}
		case Left:
			if x-1 >= minX {
				x--
			} else {
				y--
				maxY--
				dir = Up
			}
		case Up:
			if y-1 >= minY {
				y--
			} else {
				x++
				minX++
				dir = Right
			}
		}
	}

	return aas
}

func printTable(aas [][]int) {
	for _, as := range aas {
		for _, a := range as {
			fmt.Printf("%4d", a)
		}
		fmt.Println()
	}
}

func main() {
	aas := makeSnail(7, 7)
	printTable(aas)
}
