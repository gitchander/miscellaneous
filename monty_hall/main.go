package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Monty Hall problem
// https://en.wikipedia.org/wiki/Monty_Hall_problem

func main() {

	var (
		doorsNumber    int
		changeOfChoice bool
		samplesNumber  int

		duration = time.Second
	)

	flag.IntVar(&doorsNumber, "doors", 3, "number of doors")
	flag.BoolVar(&changeOfChoice, "change", false, "change of choice")
	flag.IntVar(&samplesNumber, "samples", 1000000, "number of samples")

	flag.Parse()

	source := generate(doorsNumber, samplesNumber)

	for i := 0; i < doorsNumber-2; i++ {
		source = doChoice(source, changeOfChoice)
	}

	results := winResults(source, duration)

	for result := range results {
		fmt.Printf("%.8f\n", result)
	}
}

func generate(doorsNumber, samplesNumber int) <-chan *Choice {

	output := make(chan *Choice)

	go func() {
		r := getRand()
		for i := 0; i < samplesNumber; i++ {
			output <- randChoice(r, doorsNumber)
		}
		close(output)
	}()

	return output
}

func doChoice(input <-chan *Choice, change bool) <-chan *Choice {

	output := make(chan *Choice)

	go func() {
		r := getRand()
		for c := range input {
			excludeDoor(r, c)
			if change {
				changeOfChoice(r, c)
			}
			output <- c
		}
		close(output)
	}()

	return output
}

type Stats struct {
	guard sync.Mutex
	total int
	wins  int
}

func (p *Stats) AddSample(isWin bool) {
	p.guard.Lock()
	{
		if isWin {
			p.wins++
		}
		p.total++
	}
	p.guard.Unlock()
}

func (p *Stats) Result() (result float64) {
	p.guard.Lock()
	result = float64(p.wins) / float64(p.total)
	p.guard.Unlock()
	return result
}

func winResults(input <-chan *Choice, d time.Duration) <-chan float64 {

	results := make(chan float64)

	done := make(chan struct{})

	var stats Stats

	go func() {
		defer close(done)

		for c := range input {
			isWin := (c.Doors[c.Index].inner == Car)
			stats.AddSample(isWin)
		}
	}()

	go func() {
		defer close(results)

		ticker := time.NewTicker(d)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				results <- stats.Result()
			case <-done:
				results <- stats.Result()
				return
			}
		}
	}()

	return results
}

const (
	Goat = iota
	Car
)

type Door struct {
	inner int // Goat or Car
}

func (d Door) String() string {
	return "?"
}

type Choice struct {
	Index int // Index of player's choice.
	Doors []Door
}

func randDoors(r *rand.Rand, doors []Door) {
	j := r.Intn(len(doors))
	for i := range doors {
		v := Goat
		if i == j {
			v = Car
		}
		doors[i] = Door{inner: v}
	}
}

func randChoice(r *rand.Rand, n int) *Choice {
	doors := make([]Door, n)
	randDoors(r, doors)
	return &Choice{
		Index: r.Intn(len(doors)), // first choice (1/n)
		Doors: doors,
	}
}

func changeOfChoice(r *rand.Rand, c *Choice) {
	var (
		i = c.Index
		n = len(c.Doors)
	)
	if n < 2 {
		return
	}
	for {
		j := r.Intn(n)
		if j != i {
			c.Index = j
			break
		}
	}
}

func excludeDoor(r *rand.Rand, c *Choice) {

	var (
		i = c.Index
		n = len(c.Doors)
	)

	var j int

	for {
		j = r.Intn(n)
		if (j != i) && (c.Doors[j].inner != Car) {
			break
		}
	}

	c.Doors = removeDoor(c.Doors, j)
	if j < i {
		c.Index--
	}
}

type Swapper interface {
	Len() int
	Swap(i, j int)
}

func removeElement(w Swapper, i int) int {
	n := w.Len()
	if n == 0 {
		return 0
	}
	for i < n-1 {
		w.Swap(i, i+1)
		i++
	}
	return n - 1
}

type DoorSlice []Door

func (p DoorSlice) Len() int      { return len(p) }
func (p DoorSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func removeDoor(doors []Door, i int) []Door {
	n := removeElement(DoorSlice(doors), i)
	return doors[:n]
}
