package point2f

import (
	"image"
	"math"

	"github.com/gitchander/miscellaneous/attractor/utils"
)

type Point2f struct {
	X, Y float64
}

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

//------------------------------------------------------------------------------
func PolarToPoint2f(angle, radius float64) Point2f {
	sin, cos := math.Sincos(angle)
	return Point2f{
		X: cos,
		Y: sin,
	}.MulScalar(radius)
}

func PointLerp(a, b Point2f, t float64) Point2f {
	return Point2f{
		X: utils.Lerp(a.X, b.X, t),
		Y: utils.Lerp(a.Y, b.Y, t),
	}
}
