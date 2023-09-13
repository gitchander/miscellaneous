package mscore

import (
	"errors"
	"fmt"
	"strings"

	ivl "mvstick/utils/interval"
)

var (
	ErrBrickIndex       = errors.New("Invalid brick index")
	ErrBrickID          = errors.New("Invalid brick id")
	ErrBrickCollision   = errors.New("brick collision")
	ErrBrickOutOfBounds = errors.New("brick out of bounds")
)

type Grid struct {
	size       Point
	fixedCells []Point
	bs         []*Brick
}

func (g *Grid) Clone() *Grid {

	return nil
}

func (g *Grid) Solved() bool {
	if len(g.bs) > 0 {
		b := g.bs[0]
		switch {
		case b.Ori == OriHorizontal:
			return (b.Pos.X + b.Len) == g.size.X
		case b.Ori == OriVertical:
			return (b.Pos.Y + b.Len) == g.size.Y
		}
	}
	return false
}

func (g *Grid) Print() {
	fmt.Println("size =", g.size)
	for i, b := range g.bs {
		letter, _ := indexToLetter(i)
		fmt.Printf("%q: pos = %s, ori = %s, len = %d\n", letter, b.Pos, b.Ori, b.Len)
	}
}

func (g *Grid) getBrickByIndex(i int) (*Brick, error) {
	if (0 <= i) && (i < len(g.bs)) {
		return g.bs[i], nil
	}
	return nil, ErrBrickIndex
}

func (g *Grid) getBrickByLetter(letter byte) (*Brick, error) {
	i, ok := letterToIndex(letter)
	if !ok {
		return nil, fmt.Errorf("There is no brick by letter %q", letter)
	}
	return g.getBrickByIndex(i)
}

func (g *Grid) MoveBrick(m Move) error {
	return g.MoveBrickByIndex(m.BrickIndex, m.Offset)
}

func (g *Grid) MoveBrickByIndex(brickIndex int, step int) error {
	bp, err := g.getBrickByIndex(brickIndex)
	if err != nil {
		return err
	}
	return g.moveBrick(bp, step)
}

func (g *Grid) MoveBrickByLetter(letter byte, step int) error {
	bp, err := g.getBrickByLetter(letter)
	if err != nil {
		return err
	}
	err = g.moveBrick(bp, step)
	if err != nil {
		return fmt.Errorf("brick %q %w", letter, err)
	}
	return nil
}

func (g *Grid) moveBrick(b *Brick, step int) error {

	m := NewBrickMover(b)
	m.Move(step)

	err := g.checkBrick(b)
	if err != nil {
		m.UndoMove()
		return err
	}

	return nil
}

func (g *Grid) checkBrick(b *Brick) error {

	err := g.checkBrickBounds(b)
	if err != nil {
		return err
	}

	err = g.checkBrickCollision(b)
	if err != nil {
		return err
	}

	err = g.checkBrickFixedCollision(b)
	if err != nil {
		return err
	}

	return nil
}

func (g *Grid) checkBrickBounds(b *Brick) error {

	var (
		shi = ivl.Ivl(0, g.size.X)
		svi = ivl.Ivl(0, g.size.Y)

		bI = b.makeInterval()
	)

	switch b.Ori {
	case OriHorizontal:
		if svi.Contains(b.Pos.Y) && bI.In(shi) {
			return nil
		}
	case OriVertical:
		if shi.Contains(b.Pos.X) && bI.In(svi) {
			return nil
		}
	}

	return ErrBrickOutOfBounds
}

func (g *Grid) checkBrickCollision(b *Brick) error {
	for _, bj := range g.bs {
		if bj != b {
			if collision(bj, b) {
				return ErrBrickCollision
			}
		}
	}
	return nil
}

func (g *Grid) checkBrickFixedCollision(b *Brick) error {
	for _, fixedCell := range g.fixedCells {
		if collisionBrickCell(b, fixedCell) {
			return ErrBrickCollision
		}
	}
	return nil
}

func (g *Grid) Printable() string {
	const emptyRune = '-'
	rrs := make([][]byte, g.size.Y)
	for y := range rrs {
		rs := make([]byte, g.size.X)
		for i := range rs {
			rs[i] = emptyRune
		}
		rrs[y] = rs
	}
	for i, b := range g.bs {
		var (
			letter, _ = indexToLetter(i)
			pos       = b.Pos
		)
		for k := 0; k < b.Len; k++ {
			switch b.Ori {
			case OriHorizontal:
				rrs[pos.Y][pos.X+k] = letter
			case OriVertical:
				rrs[pos.Y+k][pos.X] = letter
			}
		}
	}
	var b strings.Builder
	for _, rs := range rrs {
		b.WriteString(string(rs))
		b.WriteByte('\n')
	}
	return b.String()
}

func ParseLines(lines []string, emptyRune rune) (*Grid, error) {

	m := make(map[rune]*Brick)

	size := Point{
		X: 0,
		Y: len(lines),
	}
	var fixedCells []Point

	for y, line := range lines {

		rs := []rune(line)
		size.X = maxInt(size.X, len(rs))

		for x, r := range rs {

			if r == emptyRune {
				continue
			}
			if r == 'x' {
				fixedCells = append(fixedCells, Pt(x, y))
				continue
			}

			b, ok := m[r]
			if !ok {
				m[r] = &Brick{
					Ori: OriUnknown,
					Pos: Pt(x, y),
					Len: 1,
				}
				continue
			}

			switch b.Ori {
			case OriUnknown:
				{
					if (b.Pos.Y == y) && ((b.Pos.X + b.Len) == x) {
						b.Ori = OriHorizontal
						b.Len++
					} else if (b.Pos.X == x) && ((b.Pos.Y + b.Len) == y) {
						b.Ori = OriVertical
						b.Len++
					} else {
						return nil, fmt.Errorf("invalid %q brick %q", b.Ori, r)
					}
				}
			case OriHorizontal:
				{
					if (b.Pos.Y == y) && ((b.Pos.X + b.Len) == x) {
						b.Len++
					} else {
						return nil, fmt.Errorf("invalid %q brick %q", b.Ori, r)
					}
				}
			case OriVertical:
				{
					if (b.Pos.X == x) && ((b.Pos.Y + b.Len) == y) {
						b.Len++
					} else {
						return nil, fmt.Errorf("invalid %q brick %q", b.Ori, r)
					}
				}
			}
		}
	}

	for r, b := range m {
		if (b.Ori == OriUnknown) || (b.Len < 2) {
			return nil, fmt.Errorf("invalid %q brick %q", b.Ori, r)
		}
	}

	rs := make([]rune, 0, len(m))
	for r := range m {
		rs = append(rs, r)
	}

	runeSlice(rs).Sort()

	bs := make([]*Brick, len(rs))
	for i, r := range rs {
		bs[i] = m[r]
	}

	g := &Grid{
		size:       size,
		fixedCells: fixedCells,
		bs:         bs,
	}

	return g, nil
}

//------------------------------------------------------------------------------

// o empty cell
// x wall (fixed obstacle)
// A primary piece (red car)
// B - Z all other pieces

func ParseGrid6x6(s string) (*Grid, error) {

	m := make(map[rune]*Brick)

	size := Point{X: 6, Y: 6}
	var fixedCells []Point

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {

			r := rune(s[y*size.X+x])

			if r == 'o' { // empty cell
				continue
			}
			if r == 'x' {
				fixedCells = append(fixedCells, Pt(x, y))
				continue
			}

			b, ok := m[r]
			if !ok {
				m[r] = &Brick{
					Ori: OriUnknown,
					Pos: Pt(x, y),
					Len: 1,
				}
				continue
			}

			switch b.Ori {
			case OriUnknown:
				{
					if (b.Pos.Y == y) && ((b.Pos.X + b.Len) == x) {
						b.Ori = OriHorizontal
						b.Len++
					} else if (b.Pos.X == x) && ((b.Pos.Y + b.Len) == y) {
						b.Ori = OriVertical
						b.Len++
					} else {
						return nil, fmt.Errorf("invalid %q brick %q", b.Ori, r)
					}
				}
			case OriHorizontal:
				{
					if (b.Pos.Y == y) && ((b.Pos.X + b.Len) == x) {
						b.Len++
					} else {
						return nil, fmt.Errorf("invalid %q brick %q", b.Ori, r)
					}
				}
			case OriVertical:
				{
					if (b.Pos.X == x) && ((b.Pos.Y + b.Len) == y) {
						b.Len++
					} else {
						return nil, fmt.Errorf("invalid %q brick %q", b.Ori, r)
					}
				}
			}
		}
	}

	for r, b := range m {
		if (b.Ori == OriUnknown) || (b.Len < 2) {
			return nil, fmt.Errorf("invalid %q brick %q", b.Ori, r)
		}
	}

	rs := make([]rune, 0, len(m))
	for r := range m {
		rs = append(rs, r)
	}

	runeSlice(rs).Sort()

	bs := make([]*Brick, len(rs))
	for i, r := range rs {
		bs[i] = m[r]
	}

	g := &Grid{
		size:       size,
		fixedCells: fixedCells,
		bs:         bs,
	}

	return g, nil
}

//------------------------------------------------------------------------------

func gridGetPoses(g *Grid) (poses []int) {
	poses = make([]int, len(g.bs))
	for i, b := range g.bs {
		poses[i] = b.getMutablePos()
	}
	return poses
}

func gridSetPoses(g *Grid, poses []int) {
	for i, b := range g.bs {
		b.setMutablePos(poses[i])
	}
}
