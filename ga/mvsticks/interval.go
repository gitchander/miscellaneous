package main

type Interval struct {
	Min int
	Max int
}

//var ZI Interval // zero interval

func (v Interval) Empty() bool {
	return v.Min >= v.Max
}

func (v Interval) notEmpty() bool {
	return v.Min < v.Max
}

func (a Interval) Width() int {
	if a.notEmpty() {
		return a.Max - a.Min
	}
	return 0
}

func (v Interval) Contains(d int) bool {
	return v.notEmpty() && (v.Min <= d) && (d < v.Max)
}

func (a Interval) Overlaps(b Interval) bool {
	return a.notEmpty() && b.notEmpty() &&
		(a.Min < b.Max) && (b.Min < a.Max)
}

func (v Interval) Add(d int) Interval {
	return Interval{
		Min: v.Min + d,
		Max: v.Max + d,
	}
}

func (v Interval) Sub(d int) Interval {
	return Interval{
		Min: v.Min - d,
		Max: v.Max - d,
	}
}
