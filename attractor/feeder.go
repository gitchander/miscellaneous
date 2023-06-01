package attractor

import (
	"math/rand"

	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
	"github.com/gitchander/miscellaneous/attractor/utils/random"
)

type Feeder interface {
	Feed(Point2f) Point2f
}

// ------------------------------------------------------------------------------
type multiFeeder struct {
	fs []Feeder
	r  *rand.Rand
}

var _ Feeder = &multiFeeder{}

func newMultiFeeder(fs []Feeder) *multiFeeder {
	return &multiFeeder{
		fs: fs,
		r:  random.NewRandNow(),
	}
}

func (p *multiFeeder) Feed(a Point2f) Point2f {
	i := p.r.Intn(len(p.fs))
	return p.fs[i].Feed(a)
}

func MultiFeeder(fs ...Feeder) Feeder {
	return newMultiFeeder(fs)
}
