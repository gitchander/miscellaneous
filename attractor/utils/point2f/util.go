package point2f

import (
	"github.com/gitchander/miscellaneous/attractor/utils"
)

func LerpPoint2f(p0, p1 Point2f, t float64) Point2f {
	return Point2f{
		X: utils.Lerp(p0.X, p1.X, t),
		Y: utils.Lerp(p0.Y, p1.Y, t),
	}
}
