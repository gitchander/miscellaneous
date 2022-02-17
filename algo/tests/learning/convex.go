package learning

// Обчислює випуклий багатокутник для множини точок.
func ConvexPolygon(ps []Point) []Point {

	firstIndex, ok := firstConvexPointIndex(ps)
	if !ok {
		return nil
	}

	var cs []Point

	i := firstIndex
	cs = append(cs, ps[i])

	for {
		nextIndex, ok := nextConvexPointIndex(ps, i)
		if !ok || (nextIndex == firstIndex) {
			break
		}
		i = nextIndex
		cs = append(cs, ps[i])
	}

	return cs
}

// Шукає індекс першої точки випуклого багатокутника.
// Це точка з найменшою координатою X. Якщо таких координат
// декілька то береться точка з найменшою координатою Y.
func firstConvexPointIndex(ps []Point) (int, bool) {
	if len(ps) == 0 {
		return 0, false
	}
	firstIndex := 0
	for i := 1; i < len(ps); i++ {
		switch {
		case ps[firstIndex].X > ps[i].X:
			firstIndex = i
		case ps[firstIndex].X == ps[i].X:
			if ps[firstIndex].Y > ps[i].Y {
				firstIndex = i
			}
		}
	}
	return firstIndex, true
}

// Шукає індекс наступої точки випуклого багатокутника.
func nextConvexPointIndex(ps []Point, currIndex int) (int, bool) {
	nextIndex := -1
	var sd int
	for i := range ps {
		if (i == currIndex) || (i == nextIndex) {
			continue
		}
		if nextIndex == -1 {
			nextIndex = i
			sd = squareDistance(ps[currIndex], ps[nextIndex])
			continue
		}
		d := LineTest(ps[currIndex], ps[i], ps[nextIndex])
		if d > 0 {
			nextIndex = i
			sd = squareDistance(ps[currIndex], ps[nextIndex])
		} else if d == 0 {
			nsd := squareDistance(ps[currIndex], ps[i])
			if nsd > sd {
				nextIndex = i
				sd = nsd
			}
		}
	}
	if nextIndex == -1 {
		return 0, false
	}
	return nextIndex, true
}
