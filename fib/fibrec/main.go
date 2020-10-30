package main

import (
	"fmt"
	"math/big"
)

func main() {
	for i := 0; i < 1000; i++ {
		//Fi, _ := fib(i)
		Fi, _ := fibBig(i)
		fmt.Printf("F[%d] = %v\n", i, Fi)
	}
}

func fib1(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib1(n-2) + fib1(n-1)
}

func fib2(n int) (prev, next int) {
	if n == 0 {
		return 0, 1
	} else {
		prev, next = fib2(n - 1)
		return next, prev + next
	}
}

// this function returns (F[n], F[n+1])
func fib(n int) (Fn, Fn1 int) {

	if n < 0 {
		panic("negative number")
	}

	if n == 0 {
		Fn = 0
		Fn1 = 1
		return
	}

	Fn1, Fn = fib(n - 1)
	Fn1 = Fn + Fn1

	return
}

func fibBig(n int) (Fn, Fn1 *big.Int) {

	if n < 0 {
		panic("negative number")
	}

	if n == 0 {
		Fn = big.NewInt(0)
		Fn1 = big.NewInt(1)
		return
	}

	Fn1, Fn = fibBig(n - 1)
	Fn1.Add(Fn, Fn1)

	return
}
