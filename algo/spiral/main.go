package main

import (
	"fmt"
	"strings"
)

func main() {
	sp := CreateSpiral(7, 4)
	w := maxWidth(sp) + 1
	var b strings.Builder
	for _, line := range sp {
		for _, x := range line {
			fmt.Fprintf(&b, "%[1]*[2]d", w, x)
		}
		b.WriteByte('\n')
	}
	fmt.Print(b.String())
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxWidth(ass [][]int) int {
	n := 0
	for _, as := range ass {
		for _, a := range as {
			n = maxInt(n, numberWidth(a))
		}
	}
	return n
}

func numberWidth(a int) int {
	if a == 0 {
		return 1
	}
	const base = 10
	count := 0
	if a < 0 {
		a = -a
		count++
	}
	for a > 0 {
		a /= base
		count++
	}
	return count
}

func CreateSpiral(xn, yn int) [][]int {
	// if n < 1 {
	// 	return [][]int{}
	// }
	spiral := make([][]int, yn)
	for i := range spiral {
		spiral[i] = make([]int, xn)
	}
	var (
		minX, maxX = 0, xn
		minY, maxY = 0, yn
	)
	value := 0
	nextValue := func() int {
		value++
		return value
	}
	//for (minX < maxX) && (minY < maxY) {
	for {
		for x, y := minX, minY; x < maxX; x++ {
			spiral[y][x] = nextValue()
		}
		if minY++; minY == maxY {
			break
		}

		for x, y := (maxX - 1), minY; y < maxY; y++ {
			spiral[y][x] = nextValue()
		}
		if maxX--; minX == maxX {
			break
		}

		for x, y := (maxX - 1), (maxY - 1); x >= minX; x-- {
			spiral[y][x] = nextValue()
		}
		if maxY--; minY == maxY {
			break
		}

		for x, y := minX, (maxY - 1); y >= minY; y-- {
			spiral[y][x] = nextValue()
		}
		if minX++; minX == maxX {
			break
		}
	}
	return spiral
}
