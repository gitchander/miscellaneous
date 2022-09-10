package attractor

import (
	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
)

type Nexter interface {
	Next() Point2f
}

//------------------------------------------------------------------------------
type feedNexter struct {
	f Feeder
	p Point2f
}

var _ Nexter = &feedNexter{}

func newFeedNexter(f Feeder, p Point2f) *feedNexter {
	return &feedNexter{
		f: f,
		p: p,
	}
}

func (v *feedNexter) Next() Point2f {
	v.p = v.f.Feed(v.p)
	return v.p
}

func MakeNexter(f Feeder, p Point2f) Nexter {
	return newFeedNexter(f, p)
}
