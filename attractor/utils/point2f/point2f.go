package point2f

import (
	"image"

	"github.com/gitchander/miscellaneous/attractor/utils"
)

type Point2f struct {
	X, Y float64
}

func _() {
	image.Pt(0, 0)
}

// Pt2f is shorthand for Point2f{X, Y}.
func Pt2f(x, y float64) Point2f {
	return Point2f{
		X: x,
		Y: y,
	}
}

func (p Point2f) Point() image.Point {
	return image.Point{
		X: utils.Round(p.X),
		Y: utils.Round(p.Y),
	}
}

func (p Point2f) SetPoint(q image.Point) {
	p.X = float64(q.X)
	p.Y = float64(q.Y)
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
