package main

import (
	"fmt"
)

// termins: Neighborhood, Neighbors

// https://en.wikipedia.org/wiki/Moore_neighborhood
// https://ru.wikipedia.org/wiki/%D0%9E%D0%BA%D1%80%D0%B5%D1%81%D1%82%D0%BD%D0%BE%D1%81%D1%82%D1%8C_%D0%9C%D1%83%D1%80%D0%B0

//------------------------------------------------------------------------------
// Moore neighborhood
//
// MooreNeighborhood(1):
//
// +---+---+---+
// |   |   |   |
// +---+---+---+
// |   | P |   |
// +---+---+---+
// |   |   |   |
// +---+---+---+
//
// MooreNeighborhood(2):
//
// +---+---+---+---+---+
// |   |   |   |   |   |
// +---+---+---+---+---+
// |   |   |   |   |   |
// +---+---+---+---+---+
// |   |   | P |   |   |
// +---+---+---+---+---+
// |   |   |   |   |   |
// +---+---+---+---+---+
// |   |   |   |   |   |
// +---+---+---+---+---+

func MooreNeighborhood(n int) []Point {
	var ps []Point
	for y := -n; y <= n; y++ {
		for x := -n; x <= n; x++ {
			if (x != 0) || (y != 0) {
				ps = append(ps, Pt(x, y))
			}
		}
	}
	return ps
}

func printMooreNeighborhood(n int) {
	for y := -n; y <= n; y++ {
		var bs []byte
		for x := -n; x <= n; x++ {
			var b byte = ' '
			if (x != 0) || (y != 0) {
				b = '*'
			}
			bs = append(bs, b)
		}
		fmt.Println(string(bs))
	}
}

//------------------------------------------------------------------------------
// Von Neumann neighborhood

// VonNeumannNeighborhood(1):
//
//     +---+
//     |   |
// +---+---+---+
// |   | P |   |
// +---+---+---+
//     |   |
//     +---+
//
// VonNeumannNeighborhood(2):
//
//         +---+
//         |   |
//     +---+---+---+
//     |   |   |   |
// +---+---+---+---+---+
// |   |   | P |   |   |
// +---+---+---+---+---+
//     |   |   |   |
//     +---+---+---+
//         |   |
//         +---+

func VonNeumannNeighborhood(n int) []Point {
	var ps []Point
	for y := -n; y <= n; y++ {
		for x := -n; x <= n; x++ {
			if (absInt(x) + absInt(y)) <= n {
				if (x != 0) || (y != 0) {
					ps = append(ps, Pt(x, y))
				}
			}
		}
	}
	return ps
}

func printVonNeumannNeighborhood(n int) {
	for y := -n; y <= n; y++ {
		var bs []byte
		for x := -n; x <= n; x++ {
			var b byte = ' '
			if (absInt(x) + absInt(y)) <= n {
				if (x != 0) || (y != 0) {
					b = '*'
				}
			}
			bs = append(bs, b)
		}
		fmt.Println(string(bs))
	}
}

func absInt(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

// ------------------------------------------------------------------------------
func Neighbors(p Point, neighborhood []Point) []Point {
	neighbors := make([]Point, len(neighborhood))
	for i, d := range neighborhood {
		neighbors[i] = p.Add(d)
	}
	return neighbors
}
