package interval

import "image"

func _() {
	var r image.Rectangle
	r.In(image.ZR)
}

type Interval struct {
	Min int
	Max int
}

// ZI is the zero Interval.
var ZI Interval

// MakeInterval
func Ivl(min, max int) Interval {
	return Interval{
		Min: min,
		Max: max,
	}
}

func (x Interval) Empty() bool {
	return x.Min >= x.Max
}

func (x Interval) notEmpty() bool {
	return x.Min < x.Max
}

func (x Interval) Width() int {
	if x.notEmpty() {
		return x.Max - x.Min
	}
	return 0
}

func (x Interval) Contains(d int) bool {
	return x.notEmpty() && (x.Min <= d) && (d < x.Max)
}

func (x Interval) Overlaps(b Interval) bool {
	return x.notEmpty() && b.notEmpty() &&
		(x.Min < b.Max) && (b.Min < x.Max)
}

func (x Interval) Add(d int) Interval {
	return Interval{
		Min: x.Min + d,
		Max: x.Max + d,
	}
}

func (x Interval) Sub(d int) Interval {
	return Interval{
		Min: x.Min - d,
		Max: x.Max - d,
	}
}

// In reports whether every value in 'x' is in 'y'.
func (x Interval) In(y Interval) bool {
	if x.Empty() {
		return true
	}
	return (y.Min <= x.Min) && (x.Max <= y.Max)
}
