package mscore

import (
	"fmt"
)

type Move struct {
	BrickIndex int // Brick index
	Offset     int
}

func MakeMove(brickIndex int, offset int) Move {
	return Move{
		BrickIndex: brickIndex,
		Offset:     offset,
	}
}

func (m Move) String() string {
	letter, _ := indexToLetter(m.BrickIndex)
	return fmt.Sprintf("%c%+d", letter, m.Offset)
}

// back
func (m Move) Reverse() Move {
	return Move{
		BrickIndex: m.BrickIndex,
		Offset:     -m.Offset,
	}
}

func ParseMove(s string) (Move, error) {

	// examples: "A+1", "B-2"

	var zv Move // zero value

	const minMoveSize = 3

	if len(s) < minMoveSize {
		return zv, fmt.Errorf("parse move (%s): invalid length", s)
	}

	brickIndex, ok := letterToIndex(s[0])
	if !ok {
		return zv, fmt.Errorf("parse move (%s): invalid letter", s)
	}

	// var negative bool
	// switch sign := s[1]; sign {
	// case '-':
	// 	negative = true
	// case '+':
	// 	negative = false
	// default:
	// 	return zv, fmt.Errorf("parse move (%s): invalid sign", s)
	// }

	// offset, err := parseInt(s[2:])
	// if err != nil {
	// 	return zv, fmt.Errorf("parse move (%s): invalid offset", s)
	// }

	// if negative {
	// 	offset = -offset
	// }

	offset, err := parseInt(s[1:])
	if err != nil {
		return zv, fmt.Errorf("parse move (%s): invalid offset", s)
	}

	m := MakeMove(brickIndex, offset)
	return m, nil
}
