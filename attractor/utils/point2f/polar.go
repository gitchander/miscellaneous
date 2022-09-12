package point2f

import "math"

func PolarToPoint2f(angle, radius float64) Point2f {
	sin, cos := math.Sincos(angle)
	return Point2f{
		X: cos,
		Y: sin,
	}.MulScalar(radius)
}

func Point2fToPolar(p Point2f) (angle, radius float64) {
	angle = math.Atan2(p.Y, p.X)
	radius = math.Hypot(p.Y, p.X)
	return
}
