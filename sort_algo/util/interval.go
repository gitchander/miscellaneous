package util

type Interval struct {
	Min, Max int // [Min ... Max-1]
}

func Ivl(min, max int) Interval {
	if min > max {
		return Interval{} // It's an empty interval
	}
	return Interval{Min: min, Max: max}
}

func (i Interval) Empty() bool {
	return i.Min >= i.Max
}

func (i Interval) Middle() int {
	return (i.Min + i.Max) / 2
}

// Length is less or equal of zero if an interval is empty.
func (i Interval) Len() int {
	return i.Max - i.Min
}

func (a Interval) Equal(b Interval) bool {
	return (a == b) || (a.Empty() && b.Empty())
}

func (i Interval) Contains(x int) bool {
	return (i.Min <= x) && (x < i.Max)
}

// func (a Interval) Intersects(b Interval) bool { }
// func (a Interval) Intersection(b Interval) Interval {}

func (a Interval) Split(n int) []Interval {

	n = minInt(n, a.Len())
	if n <= 0 {
		return []Interval{}
	}

	vs := make([]Interval, n)

	quo, rem := quoRem(a.Len(), n)
	b := Interval{Min: a.Min}
	for i := range vs {
		b.Max = b.Min + quo
		if rem > 0 {
			b.Max++
			rem--
		}
		vs[i] = b
		b.Min = b.Max
	}

	return vs
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a - quo*b
	return
}
