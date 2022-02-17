package learning

type Point struct {
	X int
	Y int
}

func Pt(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}
