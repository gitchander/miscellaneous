package attractor

import (
	"fmt"
	"math/rand"

	"github.com/gitchander/miscellaneous/attractor/utils"
	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
	"github.com/gitchander/miscellaneous/attractor/utils/random"
)

type CornerpointSelector int

const (
	CornerpointRandom CornerpointSelector = iota
	CornerpointNotTwice
	CornerpointNotDirectNeighbors
)

func ParseCornerpointSelector(s string) (CornerpointSelector, error) {
	var cps CornerpointSelector
	switch s {
	case "random":
		cps = CornerpointRandom
	case "not-twice":
		cps = CornerpointNotTwice
	case "not-direct-neighbors":
		cps = CornerpointNotDirectNeighbors
	default:
		return 0, fmt.Errorf("Invalid CornerpointSelector %q", s)
	}
	return cps, nil
}

// ------------------------------------------------------------------------------
type PsFeeder struct {
	ps  []Point2f
	t   float64 // [0..1]
	cps CornerpointSelector
	r   *rand.Rand

	indexes []int
}

var _ Feeder = &PsFeeder{}

func NewPsFeeder(ps []Point2f, t float64, cps CornerpointSelector) *PsFeeder {
	return &PsFeeder{
		ps:      ps,
		t:       t,
		cps:     cps,
		r:       random.NewRandNow(),
		indexes: make([]int, 3),
	}
}

func (p *PsFeeder) lastIndex() int {
	return p.indexes[len(p.indexes)-1]
}

func (p *PsFeeder) addNewIndex(index int) {
	utils.RotateDown(intSlice(p.indexes))
	p.indexes[len(p.indexes)-1] = index
}

// Random corner point selection.
func (p *PsFeeder) nextIndex1() int {
	i := p.r.Intn(len(p.ps))
	return i
}

// Cornerpoints cannot be choosen twice in a row.
func (p *PsFeeder) nextIndex2() int {
	if len(p.ps) < 2 {
		return 0
	}
	var index int
	lastIndex := p.lastIndex()
	for {
		index = p.r.Intn(len(p.ps))
		if index != lastIndex {
			break
		}
	}
	p.addNewIndex(index)
	return index
}

// If a cornerpoint is chosen twice, the next selected point
// must not be a direct neighbor.
func (p *PsFeeder) nextIndex3() int {

	var index int

	var (
		beforeLastIndex = p.indexes[len(p.indexes)-2]
		lastIndex       = p.indexes[len(p.indexes)-1]
	)

	if beforeLastIndex == lastIndex {

		var (
			n0 = mod(lastIndex-1, len(p.ps))
			n1 = mod(lastIndex+1, len(p.ps))
		)

		for {
			index = p.r.Intn(len(p.ps))
			if (index != n0) && (index != n1) {
				break
			}
		}
	} else {
		index = p.r.Intn(len(p.ps))
	}

	p.addNewIndex(index)
	return index
}

func (p *PsFeeder) nextIndex() int {
	var index int
	switch p.cps {
	case CornerpointRandom:
		index = p.nextIndex1()
	case CornerpointNotTwice:
		index = p.nextIndex2()
	case CornerpointNotDirectNeighbors:
		index = p.nextIndex3()
	}
	return index
}

func (p *PsFeeder) Feed(a Point2f) Point2f {
	i := p.nextIndex()
	return LerpPoint2f(a, p.ps[i], p.t)
}

// ------------------------------------------------------------------------------
type intSlice []int

var _ utils.Swapper = intSlice(nil)

func (x intSlice) Len() int      { return len(x) }
func (x intSlice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func mod(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}
