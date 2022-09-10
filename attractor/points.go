package attractor

import (
	"math"
	"math/rand"

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
	rr := radius * radius
	if rr == 0 {
		return Point2f{}
	}
	for {
		p := Point2f{
			X: (1 - r.Float64()*2) * radius,
			Y: (1 - r.Float64()*2) * radius,
		}
		if ((p.X * p.X) + (p.Y * p.Y)) < rr {
			return p
		}
	}
}
