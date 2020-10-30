package main

import (
	"math/rand"
)

type Individ struct {
	ps      []Point2f
	genome  []int
	fitness float64
}

func NewIndivid(ps []Point2f) *Individ {
	genome := make([]int, len(ps))
	for i := range genome {
		genome[i] = i
	}
	d := &Individ{
		ps:     ps,
		genome: genome,
	}
	d.calcFitness()
	return d
}

func RandIndividBySeed(seed int64, n int) *Individ {
	r := newRandSeed(seed)
	ps := make([]Point2f, n)
	for i := range ps {
		ps[i] = randPoint2f(r)
	}
	return NewIndivid(ps)
}

func RandIndividGridBySeed(seed int64, n int) *Individ {
	r := newRandSeed(seed)
	ps := make([]Point2f, 0, n)
	d := 1 / float64(n-1)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			p := Point2f{
				X: float64(x) * d,
				Y: float64(y) * d,
			}
			ps = append(ps, p)
		}
	}
	shuffleElements(r, Point2fSlice(ps))
	return NewIndivid(ps)
}

func cloneInts(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func (a *Individ) Clone() *Individ {
	if a != nil {
		return &Individ{
			ps:      a.ps,
			genome:  cloneInts(a.genome),
			fitness: a.fitness,
		}
	}
	return nil
}

func (a *Individ) Fitness() float64 {
	return a.fitness
}

func Fitness(d *Individ) float64 {
	if len(d.genome) == 0 {
		return 0
	}
	var sum float64
	indexPrev := len(d.genome) - 1
	for _, index := range d.genome {
		sum += Distance(d.ps[index], d.ps[indexPrev])
		indexPrev = index
	}
	return sum
}

func (d *Individ) calcFitness() {
	if len(d.genome) == 0 {
		return
	}
	var sum float64
	indexPrev := d.genome[len(d.genome)-1]
	for _, index := range d.genome {
		sum += Distance(d.ps[index], d.ps[indexPrev])
		indexPrev = index
	}
	d.fitness = sum
}

func (d *Individ) Range(f func(i int, p Point2f) bool) {
	for i, index := range d.genome {
		if !f(i, d.ps[index]) {
			return
		}
	}
}

func (d *Individ) PointsChan() <-chan Point2f {
	ps := make(chan Point2f)
	go func() {
		for _, index := range d.genome {
			ps <- d.ps[index]
		}
		close(ps)
	}()
	return ps
}

func randomDifferentIndexes(r *rand.Rand, n int) (i, j int) {
	if n < 2 {
		return
	}
	for i == j {
		i = r.Intn(n)
		j = r.Intn(n)
	}
	return
}

func (d *Individ) RandomExchange(r *rand.Rand) {
	n := len(d.genome)
	i, j := randomDifferentIndexes(r, n)
	if i == j {
		return
	}

	d.genome[i], d.genome[j] = d.genome[j], d.genome[i]

	d.calcFitness()
}

func (d *Individ) RandomInsertion(r *rand.Rand) {
	n := len(d.genome)
	i, j := randomDifferentIndexes(r, n)
	if i == j {
		return
	}

	temp := d.genome[i]
	for i != j {
		k := mod(i+1, n)
		d.genome[i] = d.genome[k]
		i = k
	}
	d.genome[i] = temp

	d.calcFitness()
}

func (d *Individ) RandomFlip(r *rand.Rand) {
	n := len(d.genome)
	i, j := randomDifferentIndexes(r, n)
	if i == j {
		return
	}

	for {
		d.genome[i], d.genome[j] = d.genome[j], d.genome[i]

		i = mod(i+1, n)
		if i == j {
			break
		}

		j = mod(j-1, n)
		if i == j {
			break
		}
	}

	d.calcFitness()
}

func mod(x, y int) int {
	m := x % y
	if m < 0 {
		m += y
	}
	return m
}

type ByFitness []*Individ

func (p ByFitness) Len() int {
	return len(p)
}

func (p ByFitness) Less(i, j int) bool {
	return p[i].fitness < p[j].fitness
}

func (p ByFitness) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
