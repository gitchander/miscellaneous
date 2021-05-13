package main

import (
	"fmt"
)

type Stat struct {
	Total   int
	Correct int
	Errors  int
}

func (p *Stat) Add(correct bool) {
	p.Total++
	if correct {
		p.Correct++
	} else {
		p.Errors++
	}
}

func (t Stat) String() string {

	perr := percent(float64(t.Errors), float64(t.Total))

	return fmt.Sprintf("total %d, correct %d, errors %d (%.2f %%)", t.Total,
		t.Correct, t.Errors, perr)
}
