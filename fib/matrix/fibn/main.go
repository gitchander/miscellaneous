package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gitchander/miscellaneous/fib/matrix"
)

// N-th number of Fibonacci Series
// https://en.wikipedia.org/wiki/Fibonacci_number

func main() {

	var c Config

	flag.BoolVar(&(c.UseBig), "big", false, "use big.Int numbers")
	flag.IntVar(&(c.Number), "number", 12, "number of Fibonacci Series")

	flag.Parse()

	runConfig(c)
}

type Config struct {
	Number int // Index of Fibonacci number
	UseBig bool
}

func runConfig(c Config) {

	n := c.Number

	if c.UseBig {
		runBig(n)
		return
	}

	f := matrix.FibN(n)
	fmt.Printf("F[%d] = %v\n", n, f)
}

func runBig(n int) {
	start := time.Now()
	f := matrix.FibNBig(n)
	dur := time.Since(start)
	if dur > 1*time.Microsecond {
		fmt.Println("calc duration:", dur)
	}

	bitLen := f.BitLen()
	if bitLen < 1000000 {
		fmt.Printf("F[%d] = %v\n", n, f)
	}
	fmt.Printf("bitLen = %d\n", bitLen)
}
