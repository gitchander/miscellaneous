package main

var (
	pointUp    = Point{X: 0, Y: -1}
	pointDown  = Point{X: 0, Y: +1}
	pointLeft  = Point{X: -1, Y: 0}
	pointRight = Point{X: +1, Y: 0}
)

type Point struct {
	X, Y int
}

func Pt(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

// Eq reports whether 'a' and 'b' are equal.
func (a Point) Eq(b Point) bool {
	return a == b
}

func (a Point) Add(b Point) Point {
	return Point{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a Point) Sub(b Point) Point {
	return Point{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a Point) Mul(k int) Point {
	return Point{
		X: a.X * k,
		Y: a.Y * k,
	}
}

func (a Point) Div(k int) Point {
	return Point{
		X: a.X / k,
		Y: a.Y / k,
	}
}

type Rect struct {
	Min, Max Point
}

func (r Rect) Center() Point {
	return r.Min.Add(r.Max).Div(2)
}
