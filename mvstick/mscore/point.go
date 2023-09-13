package mscore

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func Pt(x, y int) Point {
	return Point{X: x, Y: y}
}
