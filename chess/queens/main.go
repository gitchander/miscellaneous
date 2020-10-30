package main

import (
	"flag"
	"fmt"
	"time"
)

// Eight queens puzzle
// https://en.wikipedia.org/wiki/Eight_queens_puzzle

type position struct {
	X, Y int
}

func main() {

	var (
		n    int
		draw bool
	)

	flag.IntVar(&n, "size", 8, "size of chess board")
	flag.BoolVar(&draw, "draw", false, "draw all positions")

	flag.Parse()

	start := time.Now()
	queens := make([]position, n)

	var count int
	var f = func() {
		if draw {
			printBoart(queens)
		}
		count++
	}
	iter(queens, 0, f)

	fmt.Printf("chess board size %dx%d:\n", n, n)
	fmt.Println("count =", count)
	fmt.Println(time.Since(start))
}

func iter(queens []position, i int, f func()) {
	n := len(queens)
	if i == n {
		f()
		return
	}

loopx:
	for x := 0; x < n; x++ {
		p := position{X: x, Y: i}
		queens[i] = p
		for _, q := range queens[:i] {
			if p.X == q.X {
				continue loopx
			}
			if p.Y == q.Y {
				continue loopx
			}
			if (p.Y - p.X) == (q.Y - q.X) {
				continue loopx
			}
			if (p.Y + p.X) == (q.Y + q.X) {
				continue loopx
			}
		}
		iter(queens, i+1, f)
	}
}

func printBoart(queens []position) {
	n := len(queens)
	line := make([]byte, n)
	for _, q := range queens {
		for i := 0; i < n; i++ {
			line[i] = '-'
		}
		line[q.X] = 'Q'
		fmt.Println(string(line))
	}
	fmt.Println()
}
