package main

import (
	"time"
)

const (
	durationMin = 10 * time.Millisecond
	durationMax = 1000 * time.Millisecond
)

func cropInt(x int, min, max int) int {
	if min > max { // It is an empty interval.
		return 0
	}
	if x < min {
		x = min
	}
	if x > max {
		x = max
	}
	return x
}

type stepDuration struct {
	d        time.Duration
	min, max time.Duration
	gain     time.Duration // growth, increase, increment
}

func (p *stepDuration) crop() {
	if p.d < p.min {
		p.d = p.min
	}
	if p.d > p.max {
		p.d = p.max
	}
}

func (p *stepDuration) Faster() {
	p.d -= p.gain
	p.crop()
}

func (p *stepDuration) Slower() {
	p.d += p.gain
	p.crop()
}

func (p *stepDuration) SetDuration(d time.Duration) {
	p.d = d
	p.crop()
}

func (p *stepDuration) Duration() time.Duration {
	return p.d
}
