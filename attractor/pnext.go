package attractor

import (
	"math/rand"

	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
	"github.com/gitchander/miscellaneous/attractor/utils/random"
)

type PtNext struct {
	ps []Point2f
	p  Point2f
	t  float64 // [0..1]
	r  *rand.Rand
}

var _ Nexter = &PtNext{}

func NewPtNext(ps []Point2f, t float64) *PtNext {
	return &PtNext{
		ps: ps,
		t:  clamp(t, 0, 1),
		r:  random.NewRandNow(),
	}
}

func (v *PtNext) Randomize(min, max Point2f) {
	v.p = Point2f{
		X: random.RandInterval(v.r, min.X, max.X),
		Y: random.RandInterval(v.r, min.Y, max.Y),
	}
}

func (v *PtNext) Next() Point2f {
	i := v.r.Intn(len(v.ps))
	v.p = PointLerp(v.p, v.ps[i], v.t)
	return v.p
}
