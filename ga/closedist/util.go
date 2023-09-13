package main

import (
	"math"
	"math/rand"
)

func sumN(n int) int {
	return n * (n + 1) / 2
}

func sqr(a float64) float64 {
	return a * a
}

func sqrt(a float64) float64 {
	return math.Sqrt(a)
}

func cloneSlice[T any](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	return b
}

func randomDifferentIndexes(r *rand.Rand, n int) (i, j int) {
	if n < 2 {
		panic("invalid n")
	}
	for i == j {
		i = r.Intn(n)
		j = r.Intn(n)
	}
	return
}

func clampFloat64(x float64, min, max float64) float64 {
	if min > max {
		return 0
	}
	if x < min {
		x = min
	}
	if x > max {
		x = max
	}
	return x
}

func mod(x, y int) int {
	m := x % y
	if m < 0 {
		m += y
	}
	return m
}

func minFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func maxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func middle(a, b float64) float64 {
	return (a + b) / 2
}

func normalize(x float64, min, max float64) float64 {
	return (x - min) / (max - min)
}
