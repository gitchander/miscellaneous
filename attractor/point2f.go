package main

import (
	"math"
)

type Point2f struct {
	X, Y float64
}

func PolarToPoint2f(angle, radius float64) Point2f {
	sin, cos := math.Sincos(angle)
	return Point2f{
		X: cos,
		Y: sin,
	}.MulScalar(radius)
}

func pointLerp(a, b Point2f, t float64) Point2f {
	return Point2f{
		X: lerp(a.X, b.X, t),
		Y: lerp(a.Y, b.Y, t),
	}
}

// Lerp - Linear interpolation
// t= [0, 1]
// (t == 0) => v0
// (t == 1) => v1
func lerp(v0, v1 float64, t float64) float64 {
	return (1.0-t)*v0 + t*v1
}

//func (p Point2f) SetPoint(q image.Point) {
//	p.X = float64(q.X)
//	p.Y = float64(q.Y)
//}

//func (p Point2f) Point() image.Point {
//	return image.Point{
//		X: round(p.X),
//		Y: round(p.Y),
//	}
//}

func (p Point2f) IntXY() (x, y int) {
	x = round(p.X)
	y = round(p.Y)
	return
}

func round(a float64) int {
	return int(math.Floor(a + 0.5))
}

func (p Point2f) Add(q Point2f) Point2f {
	return Point2f{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}

func (p Point2f) Sub(q Point2f) Point2f {
	return Point2f{
		X: p.X - q.X,
		Y: p.Y - q.Y,
	}
}

func (p Point2f) MulScalar(scalar float64) Point2f {
	return Point2f{
		X: p.X * scalar,
		Y: p.Y * scalar,
	}
}

func (p Point2f) DivScalar(scalar float64) Point2f {
	return Point2f{
		X: p.X / scalar,
		Y: p.Y / scalar,
	}
}
