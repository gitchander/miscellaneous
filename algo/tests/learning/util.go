package learning

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LineTest(m1, m2, p Point) int {
	var (
		d1 = (p.X - m1.X) * (m2.Y - m1.Y)
		d2 = (p.Y - m1.Y) * (m2.X - m1.X)
	)
	return d1 - d2
}

// Обчислює квадрат відстані між двома точками.
func squareDistance(m1, m2 Point) int {
	var (
		dx = m1.X - m2.X
		dy = m1.Y - m2.Y
	)
	return dx*dx + dy*dy
}
