package learning

// Polygon

// Intersect

// Обчислює чи можна провести пряму яка однозначно розділить (відокремить) дві множини точок.
func SeparableStraight(ps1, ps2 []Point) bool {

	var (
		cp1 = ConvexPolygon(ps1)
		cp2 = ConvexPolygon(ps2)
	)

	if (len(cp1) == 0) || (len(cp2) == 0) {
		return false
	}

	// Both lengths are 1 or 2.
	if (len(cp1) <= 2) && (len(cp2) <= 2) {
		switch len(cp1) {
		case 1:
			switch len(cp2) {
			case 1:
				return cp1[0] != cp2[0]
			case 2:
				return !pointBetween(cp2[0], cp2[1], cp1[0])
			}
		case 2:
			switch len(cp2) {
			case 1:
				return !pointBetween(cp1[0], cp1[1], cp2[0])
			case 2:
				return !crossingLines(cp1[0], cp1[1], cp2[0], cp2[1])
			}
		}
	}

	if len(cp1) > 2 {
		for _, p := range cp2 {
			if convexPolygonContains(cp1, p) {
				return false
			}
		}
	}
	if len(cp2) > 2 {
		for _, p := range cp1 {
			if convexPolygonContains(cp2, p) {
				return false
			}
		}
	}

	return true
}

func convexPolygonContains(ps []Point, p Point) bool {
	var ok bool
	f := func(m1, m2 Point) bool {
		d := LineTest(m1, m2, p)
		ok = (d >= 0)
		return ok
	}
	walkPolygonLines(ps, f)
	return ok
}

// Walk lines with indexes: (0, 1), (1, 2), (2, 3), ... (n-2, n-1), (n-1, 0)
func walkPolygonLines(ps []Point, f func(m1, m2 Point) bool) {
	n := len(ps)
	if n < 2 {
		return
	}
	i := n - 1
	for j := 0; j < n; j++ {
		if !f(ps[i], ps[j]) {
			return
		}
		i = j
	}
}

// Point p between (m1, m2).
// Повертає істину якщо точка p лежить між двома точками m1 та m2 (включаючи ці точки).
func pointBetween(m1, m2 Point, p Point) bool {

	var (
		min = Point{
			X: minInt(m1.X, m2.X),
			Y: minInt(m1.Y, m2.Y),
		}

		max = Point{
			X: maxInt(m1.X, m2.X),
			Y: maxInt(m1.Y, m2.Y),
		}
	)

	if (p.X < min.X) || (max.X < p.X) {
		return false
	}
	if (p.Y < min.Y) || (max.Y < p.Y) {
		return false
	}

	d := LineTest(m1, m2, p)
	return d == 0
}

// Повертає істину, якщо у двох відрізків є спільні точки (вони перетинаються).
func crossingLines(a1, a2 Point, b1, b2 Point) bool {

	// Point b1 between (a1, a2).
	if pointBetween(a1, a2, b1) {
		return true
	}
	// Point b2 between (a1, a2).
	if pointBetween(a1, a2, b2) {
		return true
	}
	// Point a1 between (b1, b2).
	if pointBetween(b1, b2, a1) {
		return true
	}
	// Point a2 between (b1, b2).
	if pointBetween(b1, b2, a2) {
		return true
	}

	db1 := LineTest(a1, a2, b1)
	db2 := LineTest(a1, a2, b2)

	da1 := LineTest(b1, b2, a1)
	da2 := LineTest(b1, b2, a2)

	if ((db1 * db2) < 0) && ((da1 * da2) < 0) {
		return true
	}

	return false
}
