package mscore

import (
	"mvstick/utils/optional"
)

type optPoint = optional.Optional[Point]

type BrickMover struct {
	prevPos optPoint
	b       *Brick
}

func NewBrickMover(b *Brick) *BrickMover {
	return &BrickMover{
		b: b,
	}
}

func (bm *BrickMover) Move(step int) {
	bm.prevPos = optional.MakePresent[Point](bm.b.Pos)
	bm.b.Move(step)
}

func (bm *BrickMover) UndoMove() {
	if bm.prevPos.Present {
		bm.prevPos.Present = false
		bm.b.Pos = bm.prevPos.Value
	}
}
