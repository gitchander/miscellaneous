package main

import (
	"fmt"
	"log"
	"math"
)

const epsilon = 1e-13

//------------------------------------------------------------------------------
// sqrt - square root of s.
func sqrt(s float64) float64 {
	a := s
	for {
		// a * b = s
		// b = s / a
		b := s / a

		//fmt.Printf("(%.20f, %.20f)\n", a, b)

		middle := (a + b) / 2
		if math.Abs(a-b) < epsilon {
			return middle
		}
		a = middle
	}
	return a
}

// cbrt - cube root of s.
func cbrt(s float64) float64 {
	a := s
	for {
		// a * b * b = s
		// b * b = s / a
		// b = sqrt(s / a)
		b := sqrt(s / a)

		middle := (a + b + b) / 3
		if math.Abs(a-b) < epsilon {
			return middle
		}
		a = middle
	}
	return a
}

func cubeFloat64(a float64) float64 {
	return a * a * a
}

func testSqrt() {
	// fmt.Println(sqrt(2))
	fmt.Println(sqrt(2))
	//fmt.Println(math.Sqrt(8))

	// fmt.Println(math.Pow(5, 1.0/3.0))
	fmt.Println(cbrt(cubeFloat64(3)))
}

//------------------------------------------------------------------------------
// sqrt int
//------------------------------------------------------------------------------
func sqrtIntMath(s int) (int, bool) {
	if s < 0 {
		return 0, false
	}
	a := int(math.Round(math.Sqrt(float64(s))))
	if (a * a) == s {
		return a, true
	}
	return 0, false
}

func sqrtInt(s int) (int, bool) {
	if s < 0 {
		return 0, false
	}
	a := s
	for {
		if (a * a) == s {
			return a, true
		}

		b := s / a

		if absInt(a-b) <= 1 {
			return 0, false
		}
		a = (a + b) / 2 // middle
	}
}

func absInt(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func squareInt(a int) int {
	return a * a
}

func testSqrtInt() {

	fmt.Println(sqrtInt(squareInt(678567954)))
	return

	for i := 0; i < 1_000_000; i++ {

		a1, ok1 := sqrtIntMath(i)
		a2, ok2 := sqrtInt(i)

		if ok1 != ok2 {
			log.Fatalf("sqrt(%d): (a1(%t) != a2(%t))", i, ok1, ok2)
		}

		if a1 != a2 {
			log.Fatalf("sqrt(%d): (a1(%d) != a2(%d))", i, a1, a2)
		}
	}
}

func main() {
	//testSqrt()
	testSqrtInt()
}
