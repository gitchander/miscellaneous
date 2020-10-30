package rational

import (
	"fmt"
	"strconv"
)

// Rational number
// https://en.wikipedia.org/wiki/Rational_number

type Rational struct {
	p, q int // p / q
}

var zero = RatInt(0)

func Zero() Rational {
	return zero
}

func Rat(p, q int) Rational {
	return Rational{
		p: p,
		q: q,
	}.normal()
}

func RatInt(x int) Rational {
	return Rational{
		p: x,
		q: 1,
	}
}

func (a Rational) Split() (p, q int) {
	p = a.p
	q = a.q
	return
}

func (a Rational) Equal(b Rational) bool {
	return (a.p * b.q) == (a.q * b.p)
}

func (a Rational) Add(b Rational) Rational {
	return Rational{
		p: (a.p * b.q) + (b.p * a.q),
		q: (a.q * b.q),
	}.normal()
}

func (a Rational) Sub(b Rational) Rational {
	return Rational{
		p: (a.p * b.q) - (b.p * a.q),
		q: (a.q * b.q),
	}.normal()
}

func (a Rational) Mul(b Rational) Rational {
	return Rational{
		p: a.p * b.p,
		q: a.q * b.q,
	}.normal()
}

func (a Rational) Div(b Rational) Rational {
	return Rational{
		p: a.p * b.q,
		q: a.q * b.p,
	}.normal()
}

func (a Rational) normal() Rational {

	var (
		p = a.p
		q = a.q
	)

	if q == 0 {
		panic("invalid rational: q == 0")
	}
	if p == 0 {
		return zero
	}

	// move sign to p
	if q < 0 {
		q = -q
		p = -p
	}

	var negative bool
	if p < 0 {
		negative = true
		p = -p
	}

	d := gcd_Euclidean(p, q)
	if d > 1 {
		p /= d
		q /= d
	}

	if negative {
		p = -p
	}

	return Rational{
		p: p,
		q: q,
	}
}

func (a Rational) String() string {
	if a.q == 1 {
		return strconv.Itoa(a.p)
	}
	return fmt.Sprintf("%d/%d", a.p, a.q)
}

// gcd - Greatest common divisor
// https://en.wikipedia.org/wiki/Greatest_common_divisor

func gcd_Euclidean(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func gcd_Remainder(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
