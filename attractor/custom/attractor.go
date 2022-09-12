package main

import (
	"math"
	"math/rand"

	attr "github.com/gitchander/miscellaneous/attractor"
	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
	"github.com/gitchander/miscellaneous/attractor/utils/random"
)

//------------------------------------------------------------------------------
type Attractor1 struct {
	A, B, C, D float64
}

var _ attr.Feeder = &Attractor1{}

func (t *Attractor1) Feed(p Point2f) Point2f {
	return Point2f{
		X: math.Sin(t.A * p.Y * t.B),
		Y: math.Sin(t.C*p.X+t.C) + math.Cos(t.D*p.Y),
	}.DivScalar(2)
}

func randAttractor1(r *rand.Rand) *Attractor1 {

	const (
		//d = 1.0
		//	d = math.Pi

		d = 4

		min = -d
		max = +d
	)

	return &Attractor1{
		A: random.RandInterval(r, min, max),
		B: random.RandInterval(r, min, max),
		C: random.RandInterval(r, min, max),
		D: random.RandInterval(r, min, max),
	}
}

//------------------------------------------------------------------------------
type attractor2 struct {
	r          *rand.Rand
	a, b, c, d float64
}

var _ attr.Feeder = &attractor2{}

func (t *attractor2) Feed(p Point2f) Point2f {

	// if (fRandom.Next() > 0.5)
	// {
	//         _x= sin(a * y) + cos(b* x + d*y);
	//         _y= sin(c * x) + cos(d * y );
	// }
	// else
	// {
	//         _x= -y * sin(y * x + x*sin(y*d + c*sin(a*x)));
	//         _y= x * sin(b * x) + a*b * sin(c*y);
	// }

	var q Point2f

	if t.r.Float64() > 0.5 {
		q = Point2f{
			X: math.Sin(t.a*p.Y) + math.Cos(t.b*p.X+t.d*p.Y),
			Y: math.Sin(t.c*p.X) + math.Cos(t.d*p.Y),
		}
	} else {
		q = Point2f{
			X: -p.Y * math.Sin(p.Y*p.X+p.X*math.Sin(p.Y*t.d+t.c*math.Sin(t.a*p.X))),
			Y: p.X*math.Sin(t.b*p.X) + t.a*t.b*math.Sin(t.c*p.Y),
		}
	}

	return q.DivScalar(2)
}

func randAttractor2(r *rand.Rand) *attractor2 {

	const (
		//d = 1.0
		//	d = math.Pi

		d = 4

		min = -d
		max = +d
	)

	return &attractor2{
		r: random.NewRandNow(),

		a: random.RandInterval(r, min, max),
		b: random.RandInterval(r, min, max),
		c: random.RandInterval(r, min, max),
		d: random.RandInterval(r, min, max),
	}
}

//------------------------------------------------------------------------------

type attractor3 struct {
	r          *rand.Rand
	a, b, c, d float64
}

var _ attr.Feeder = &attractor3{}

func (t *attractor3) Feed(p Point2f) Point2f {

	var q Point2f

	//--------------------------------------------------------------------------
	// Found: 2010.12.21
	if t.r.Float64() > 0.5 {
		q = Point2f{
			X: math.Sin(t.a*p.Y) + math.Cos(t.b*p.X+t.d*p.Y),
			Y: math.Sin(t.c*p.X) + math.Cos(t.d*p.Y),
		}
	} else {
		q = Point2f{
			X: -p.Y,
			Y: math.Sin(t.b * p.X),
		}
	}
	//--------------------------------------------------------------------------
	// if t.r.Float64() > 0.5 {
	// 	q = Point2f{
	// 		X: math.Sin(t.a*p.Y) + math.Cos(t.b*p.X+t.b*p.Y),
	// 		Y: math.Sin(t.c*p.X) + math.Cos(t.d*p.Y+t.d*p.X),
	// 	}
	// } else {
	// 	q = Point2f{
	// 		X: -p.Y,
	// 		Y: p.X + p.Y,
	// 	}
	// }
	//--------------------------------------------------------------------------

	return q.DivScalar(2)
}

func randAttractor3(r *rand.Rand) *attractor3 {

	const (
		d = 4

		min = -d
		max = +d
	)

	return &attractor3{
		r: random.NewRandNow(),

		a: random.RandInterval(r, min, max),
		b: random.RandInterval(r, min, max),
		c: random.RandInterval(r, min, max),
		d: random.RandInterval(r, min, max),
	}
}
