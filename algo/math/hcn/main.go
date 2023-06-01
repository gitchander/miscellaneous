package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

func main() {
	testNextHCN_List()
	//testNumberOfDivisors()
	//testFactors()
}

func digitsOfNumber(number int) []int {
	if number == 0 {
		return []int{0}
	}
	var ds []int
	for number > 0 {
		ds = append(ds, number%10)
		number /= 10
	}
	return ds
}

func testNextHCN_List() {
	var x, n int = 1, 0
	for i := 0; i < 100; i++ {
		start := time.Now()
		x, n = NextHCN(x, n)
		fmt.Printf("HCN(%d) = %d (%d divisors, duration %s)\n", i+1, x, n, time.Since(start))
	}
}

func testNextHCN() {
	start := time.Now()
	var (
		i = 61
		x = 245044800
		n = numberOfDivisors(x)
	)
	x, n = NextHCN(x, n)
	fmt.Printf("HCN(%d) = %d (%d divisors)\n", i+1, x, n)
	fmt.Println(time.Since(start))
}

// x - start value
// n - prev number of divisors
func NextHCN(x, n int) (int, int) {
	for i := x; ; i++ {
		if m := numberOfDivisors(i); m > n {
			return i, m
		}
	}
}

//------------------------------------------------------------------------------
// Superior highly composite number
// https://en.wikipedia.org/wiki/Superior_highly_composite_number

//------------------------------------------------------------------------------
// Highly composite number
// https://en.wikipedia.org/wiki/Highly_composite_number

// HCN - Highly composite number
func calcHCN() {
	k := 0
	j := 1
	n := 1000_000_000
	for i := 0; i < n; i++ {
		kk := numberOfDivisors(i)
		if kk > k {
			k = kk
			fmt.Printf("HCN(%2d) = %d\n", j, i)
			j++
		}
	}
}

// ------------------------------------------------------------------------------
var (
	//numberOfDivisors = numberOfDivisorsV1
	//numberOfDivisors = numberOfDivisorsV2
	numberOfDivisors = numberOfDivisorsV3
)

// Divisor function
// https://en.wikipedia.org/wiki/Divisor_function
func numberOfDivisorsV1(a int) int {
	var (
		n = 0
		d = 1
	)
	for d*d < a {
		if a%d == 0 {
			n += 2
		}
		d++
	}
	if d*d == a {
		n++
	}
	return n
}

// https://siongui.github.io/2017/05/09/go-find-all-prime-factors-of-integer-number/

func numberOfDivisorsV2(n int) int {

	v := 1

	f := func(p int) {
		factor := 0
		for (n % p) == 0 {
			n /= p
			factor++
		}
		if factor > 0 {
			v *= factor + 1
		}
	}

	f(2)

	for p := 3; p*p <= n; p += 2 {
		f(p)
	}

	if n >= 2 {
		f(n)
	}

	return v
}

func numberOfDivisorsV3(n int) int {

	v := 1

	{
		p := 2

		factor := 0
		for (n % p) == 0 {
			n /= p
			factor++
		}
		if factor > 0 {
			v *= factor + 1
		}
	}

	for p := 3; p*p <= n; p += 2 {
		factor := 0
		for (n % p) == 0 {
			n /= p
			factor++
		}
		if factor > 0 {
			v *= factor + 1
		}
	}

	if n >= 2 {
		v *= 2
	}

	return v
}

func testNumberOfDivisors() {

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	start := time.Now()
	for i := 1; i < 1000_000; i++ {

		select {
		case <-(ticker.C):
			fmt.Println(i)
		default:
		}

		var (
			have = numberOfDivisorsV3(i)
			want = numberOfDivisorsV1(i)
		)
		if have != want {
			log.Fatalf("invalid NOD: have %d, want %d", have, want)
		}
	}
	fmt.Println(time.Since(start))
}

func allDivisors(a int) []int {
	var (
		ds []int
		d  = 1
	)
	for d*d < a {
		if a%d == 0 {
			ds = append(ds, d, a/d)
		}
		d++
	}
	if d*d == a {
		ds = append(ds, d)
	}
	sort.Ints(ds)
	return ds
}
