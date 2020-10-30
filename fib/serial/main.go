package main

import (
	"flag"
	"fmt"
	"math/big"
)

// Fibonacci Series

func main() {

	var (
		fBig    = flag.Bool("big", false, "use big.Int numbers")
		fNumber = flag.Int("number", 12, "number of terms")
	)

	flag.Parse()

	n := *fNumber

	if *fBig {
		next := nextFibBig()
		for i := 0; i < n; i++ {
			fmt.Printf("F[%d] = %v\n", i+1, next())
		}
		return
	}

	next := nextFib()
	for i := 0; i < n; i++ {
		fmt.Printf("F[%d] = %v\n", i+1, next())
	}
}

// closure fib

func nextFib() func() int {
	var a, b int = 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func nextFibBig() func() *big.Int {
	var (
		a = big.NewInt(0)
		b = big.NewInt(1)
		c = new(big.Int)
	)
	return func() *big.Int {
		c.Add(a, b)
		a.Set(b)
		b.Set(c)
		return a
	}
}

func fibN(n int) *big.Int {
	var (
		a = big.NewInt(0)
		b = big.NewInt(1)
		c = new(big.Int)
	)
	for i := 0; i < n; i++ {
		c.Add(a, b)
		a.Set(b)
		b.Set(c)
	}
	return a
}
