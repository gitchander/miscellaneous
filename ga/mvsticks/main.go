package main

import (
	"fmt"
	"image"
	"log"
	"sort"
)

func _() {
	var r image.Rectangle
	_ = r
}

func main() {

	lines := []string{
		"BBCCH-",
		"I--JHK",
		"IAAJ-K",
		"LDDD--",
		"L-MEE-",
		"FFMGG-",
	}
	g, err := ParseLines(lines, '-')
	checkError(err)

	g.Print()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/*

"BBCCH-"
"I--JHK"
"IAAJ-K"
"LDDD--"
"L-MEE-"
"FFMGG-"

*/

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func PointByXY(x, y int) Point {
	return Point{X: x, Y: y}
}

func ParseLines(lines []string, fillRune rune) (*Grid, error) {

	m := make(map[rune]*Brick)

	var (
		dx = 0
		dy = len(lines)
	)

	for y, line := range lines {
		rs := []rune(line)
		dx = maxInt(dx, len(rs))
		for x, r := range rs {
			if r != fillRune {
				b, ok := m[r]
				if !ok {
					m[r] = &Brick{
						Pos: PointByXY(x, y),
						Ori: Unknown,
						Len: 1,
					}
				} else {
					switch b.Ori {
					case Unknown:

						if (b.Pos.Y == y) && ((b.Pos.X + b.Len) == x) {
							b.Ori = Horizontal
							b.Len++
						} else if (b.Pos.X == x) && ((b.Pos.Y + b.Len) == y) {
							b.Ori = Vertical
							b.Len++
						} else {
							return nil, fmt.Errorf("invalid %q brick %q", Unknown, r)
						}

					case Horizontal:
						if (b.Pos.Y != y) || ((b.Pos.X + b.Len) != x) {
							return nil, fmt.Errorf("invalid %q brick %q", Horizontal, r)
						} else {
							b.Len++
						}

					case Vertical:
						if (b.Pos.X != x) || ((b.Pos.Y + b.Len) != y) {
							return nil, fmt.Errorf("invalid %q brick %q", Vertical, r)
						} else {
							b.Len++
						}
					}
				}
			}
		}
	}

	for r, b := range m {
		if b.Ori == Unknown {
			return nil, fmt.Errorf("invalid %q brick %q", Unknown, r)
		}
	}

	rs := make([]rune, 0, len(m))
	for r := range m {
		rs = append(rs, r)
	}

	sort.Sort(RuneSlice(rs))

	bs := make([]*Brick, len(rs))
	for i, r := range rs {
		bs[i] = m[r]
	}

	g := &Grid{
		size: Point{
			X: dx,
			Y: dy,
		},
		rs: rs,
		bs: bs,
	}

	return g, nil
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Orientation int

const (
	Unknown Orientation = iota
	Horizontal
	Vertical
)

func (o Orientation) String() string {
	switch o {
	case Unknown:
		return "unknown"
	case Horizontal:
		return "horizontal"
	case Vertical:
		return "vertical"
	default:
		return fmt.Sprintf("Orientation(%d)", int(o))
	}
}

type Brick struct {
	Pos Point
	Ori Orientation
	Len int
}

type RuneSlice []rune

var _ sort.Interface = RuneSlice(nil)

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Grid struct {
	size Point
	rs   []rune // id list
	bs   []*Brick
}

func (g *Grid) Print() {

	fmt.Println("size =", g.size)

	for i, b := range g.bs {
		fmt.Printf("%q: pos = %s, ori = %s, len = %d\n", g.rs[i], b.Pos, b.Ori, b.Len)
	}
}

// func (g *Grid) MoveBrick(i int, d int) bool {

// 	if (i < 0) || (len(g.bs) <= i) {
// 		return false
// 	}

// 	b := g.bs[i]

// 	if b.Ori == Horizontal {
// 		b.Pos.X += d
// 	} else if b.Ori == Vertical {
// 		b.Pos.Y += d
// 	}

// 	return false
// }

func (g *Grid) CollisionByIndex(i int) bool {
	bi := g.bs[i]
	for j, bj := range g.bs {
		if i != j {
			if collision(bi, bj) {
				return true
			}
		}
	}
	return false
}

func collision(a, b *Brick) bool {
	if a.Ori == Horizontal {

		aH := Interval{
			Min: a.Pos.X,
			Max: a.Pos.X + a.Len,
		}

		if b.Ori == Horizontal {

			bH := Interval{
				Min: b.Pos.X,
				Max: b.Pos.X + b.Len,
			}

			if a.Pos.Y == b.Pos.Y {
				return aH.Overlaps(bH)
			}

		} else if b.Ori == Vertical {

			bV := Interval{
				Min: b.Pos.Y,
				Max: b.Pos.Y + b.Len,
			}

			if aH.Contains(b.Pos.X) && bV.Contains(a.Pos.Y) {
				return true
			}

		}
	} else if a.Ori == Vertical {

		// aV := Interval{
		// 	Min: a.Pos.Y,
		// 	Max: a.Pos.Y + a.Len,
		// }

		if b.Ori == Horizontal {
			// ...
		} else if b.Ori == Vertical {
			// ...
		}
	}
	return false
}
