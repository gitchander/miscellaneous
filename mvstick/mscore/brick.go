package mscore

import (
	"fmt"

	ivl "mvstick/utils/interval"
	"mvstick/utils/varint"
)

type Brick struct {
	Ori Orientation
	Pos Point // Position
	Len int   // Size, Width
}

func (b *Brick) Clone() *Brick {
	if b == nil {
		return nil
	}
	return &Brick{
		Ori: b.Ori,
		Pos: b.Pos,
		Len: b.Len,
	}
}

func (b *Brick) makeInterval() ivl.Interval {
	switch b.Ori {
	case OriHorizontal:
		return ivl.Interval{
			Min: b.Pos.X,
			Max: b.Pos.X + b.Len,
		}
	case OriVertical:
		return ivl.Interval{
			Min: b.Pos.Y,
			Max: b.Pos.Y + b.Len,
		}
	default:
		return ivl.ZI
	}
}

// Horizontal Interval
func (b *Brick) hi() ivl.Interval {
	var (
		have = b.Ori
		want = OriHorizontal
	)
	if have != want {
		err := fmt.Errorf("invalid brick orientation: have %s, want %s", have, want)
		panic(err)
	}
	return ivl.Interval{
		Min: b.Pos.X,
		Max: b.Pos.X + b.Len,
	}
}

// Vertical Interval
func (b *Brick) vi() ivl.Interval {
	var (
		have = b.Ori
		want = OriVertical
	)
	if have != want {
		err := fmt.Errorf("invalid brick orientation: have %s, want %s", have, want)
		panic(err)
	}
	return ivl.Interval{
		Min: b.Pos.Y,
		Max: b.Pos.Y + b.Len,
	}
}

//------------------------------------------------------------------------------

func (b *Brick) getMutablePos() int {
	var pos int
	switch b.Ori {
	case OriHorizontal:
		pos = b.Pos.X
	case OriVertical:
		pos = b.Pos.Y
	}
	return pos
}

func (b *Brick) setMutablePos(pos int) {
	switch b.Ori {
	case OriHorizontal:
		b.Pos.X = pos
	case OriVertical:
		b.Pos.Y = pos
	}
}

func (b *Brick) Move(step int) {
	switch b.Ori {
	case OriHorizontal:
		b.Pos.X += step
	case OriVertical:
		b.Pos.Y += step
	}
}

func collisionBrickCell(b *Brick, cell Point) bool {
	switch b.Ori {
	case OriHorizontal:
		return (b.Pos.Y == cell.Y) && (b.hi().Contains(cell.X))
	case OriVertical:
		return (b.Pos.X == cell.X) && (b.vi().Contains(cell.Y))
	}
	return false
}

func collisionUseMakeInterval(a, b *Brick) bool {
	var (
		aI = a.makeInterval()
		bI = b.makeInterval()
	)
	if a.Ori == OriHorizontal {
		switch b.Ori {
		case OriHorizontal:
			return (a.Pos.Y == b.Pos.Y) && aI.Overlaps(bI)
		case OriVertical:
			return aI.Contains(b.Pos.X) && bI.Contains(a.Pos.Y)
		}
	} else if a.Ori == OriVertical {
		switch b.Ori {
		case OriHorizontal:
			return aI.Contains(b.Pos.Y) && bI.Contains(a.Pos.X)
		case OriVertical:
			return (a.Pos.X == b.Pos.X) && aI.Overlaps(bI)
		}
	}
	return false
}

func collisionUseViHi(a, b *Brick) bool {
	if a.Ori == OriHorizontal {
		switch b.Ori {
		case OriHorizontal:
			return (a.Pos.Y == b.Pos.Y) && a.hi().Overlaps(b.hi())
		case OriVertical:
			return a.hi().Contains(b.Pos.X) && b.vi().Contains(a.Pos.Y)
		}
	} else if a.Ori == OriVertical {
		switch b.Ori {
		case OriHorizontal:
			return a.vi().Contains(b.Pos.Y) && b.hi().Contains(a.Pos.X)
		case OriVertical:
			return (a.Pos.X == b.Pos.X) && a.vi().Overlaps(b.vi())
		}
	}
	return false
}

var (
	collision = collisionUseMakeInterval
	//collision = collisionUseViHi
)

func encodeBrickPos(data []byte, b *Brick) int {
	pos := b.getMutablePos()
	return varint.EncodeInt64(data, int64(pos))
}

func decodeBrickPos(data []byte, b *Brick) int {
	v, n, _ := varint.DecodeInt64(data)
	pos := int(v)
	b.setMutablePos(pos)
	return n
}

func encodePosesLen(bs []*Brick) int {
	return len(bs) * varint.MaxSize
}

func encodePoses(data []byte, bs []*Brick) int {
	n := 0
	for _, b := range bs {
		n += encodeBrickPos(data[n:], b)
	}
	return n
}

func decodePoses(data []byte, bs []*Brick) int {
	n := 0
	for _, b := range bs {
		n += decodeBrickPos(data[n:], b)
	}
	return n
}

func encodePosesBytes(bs []*Brick) []byte {
	data := make([]byte, encodePosesLen(bs))
	n := encodePoses(data, bs)
	return data[:n]
}
