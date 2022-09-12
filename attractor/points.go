package attractor

import (
	"math"
	"math/rand"

	"github.com/gitchander/miscellaneous/attractor/utils"
	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
)

func MakeRegularPoints(n int) []Point2f {
	return RegularPoints(n, 1, -0.25)
}

func RegularPoints(n int, radius float64, phase float64) []Point2f {
	const tau = 2 * math.Pi
	ps := make([]Point2f, n)
	var (
		tauPhase = tau * phase
		du       = tau / float64(n)
	)
	for i := range ps {
		u := float64(i)*du + tauPhase
		ps[i] = PolarToPoint2f(u, radius)
	}
	return ps
}

func RandPointInRadius(r *rand.Rand, radius float64) Point2f {

	switch {
	case math.IsInf(radius, 0):
		return Point2f{}
	case math.IsNaN(radius):
		return Point2f{}
	}

	radius = math.Abs(radius)

	if radius == 0 {
		return Point2f{}
	}

	rr := radius * radius
	for {
		p := Point2f{
			X: utils.Lerp(-1, +1, r.Float64()),
			Y: utils.Lerp(-1, +1, r.Float64()),
		}.MulScalar(radius)
		if ((p.X * p.X) + (p.Y * p.Y)) < rr {
			return p
		}
	}
}
