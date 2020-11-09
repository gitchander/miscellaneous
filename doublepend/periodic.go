package main

import (
	"time"
)

func runPeriodic(period time.Duration, f func() bool) {
	t := time.Now()
	for {
		if !f() {
			return
		}
		// calc sleep
		t = t.Add(period)
		d := t.Sub(time.Now())
		if d > 0 {
			time.Sleep(d)
		}
	}
}
