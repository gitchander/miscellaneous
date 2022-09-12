package attractor

import (
	"math/rand"

	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
	"github.com/gitchander/miscellaneous/attractor/utils/random"
)

type PsFeeder struct {
	ps []Point2f
	t  float64 // [0..1]
	r  *rand.Rand
}

var _ Feeder = &PsFeeder{}

func NewPsFeeder(ps []Point2f, t float64) *PsFeeder {
	return &PsFeeder{
		ps: ps,
		t:  clamp(t, 0, 1),
		r:  random.NewRandNow(),
	}
}

func (p *PsFeeder) Feed(a Point2f) Point2f {
	i := p.r.Intn(len(p.ps))
	return LerpPoint2f(a, p.ps[i], p.t)
}
