package main

import "time"

type sleeper interface {
	Sleep()
}

func newSleeper(d time.Duration) sleeper {
	if d > 0 {
		return &realSleeper{
			d: d,
			t: time.Now(),
		}
	}
	return dummySleeper{}
}

type dummySleeper struct{}

var _ sleeper = dummySleeper{}

func (dummySleeper) Sleep() {}

type realSleeper struct {
	d time.Duration
	t time.Time
}

var _ sleeper = &realSleeper{}

func (p *realSleeper) Sleep() {
	p.t = p.t.Add(p.d)
	d := p.t.Sub(time.Now())
	if d > 0 {
		time.Sleep(d)
	}
}
