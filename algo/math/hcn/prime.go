package main

import (
	"fmt"
)

type PrimeFactor struct {
	Prime  int
	Factor int
}

func pf(prime, factor int) PrimeFactor {
	return PrimeFactor{
		Prime:  prime,
		Factor: factor,
	}
}

// https://www.geeksforgeeks.org/print-all-prime-factors-of-a-given-number/
// https://siongui.github.io/2017/05/09/go-find-all-prime-factors-of-integer-number/#footnote-1

func PrimeFactors(n int) (pfs []int) {

	if n <= 0 {
		return pfs
	}

	{
		p := 2
		for (n % p) == 0 {
			pfs = append(pfs, p)
			n /= p
		}
	}

	for p := 3; p*p <= n; p += 2 {
		for (n % p) == 0 {
			pfs = append(pfs, p)
			n /= p
		}
	}
	if n > 2 {
		pfs = append(pfs, n)
	}

	return pfs
}

func testFactors() {
	pfs := PrimeFactors(551350800) // [2 2 2 2 3 3 3 3 5 5 7 11 13 17]
	fmt.Println(pfs)
}

// n = (p ^ factor) * rest
func calcFactor(n, p int) (factor, rest int) {
	factor = 0
	for (n % p) == 0 {
		n /= p
		factor++
	}
	return factor, n
}
