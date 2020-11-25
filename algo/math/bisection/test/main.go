package main

import (
	"fmt"
	"log"

	"github.com/gitchander/miscellaneous/algo/math/bisection"
)

func main() {

	f := func(x float64) float64 {
		return x*x*x - x + 2
	}

	var (
		// interval
		a = -2.0
		b = +2.0

		tolerance = 1E-15

		nmax = 60
	)

	c, err := bisection.Bisection(f, a, b, tolerance, nmax)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("bisection result:", c)
}
