package main

// [ Min ... Max-1 ]
type Interval struct {
	Min, Max int
}

func (v Interval) Empty() bool {
	return v.Min >= v.Max
}
