package main

import (
	"math/rand"
)

type Individ struct {
	ps       []Point2f
	distance float64
	fitness  float64
}

func NewIndivid(ps []Point2f, distance float64) *Individ {
	d := &Individ{
		ps:       ps,
		distance: distance,
	}
	d.calcFitness()
	return d
}

func RandIndivid(r *rand.Rand, n int, distance float64) *Individ {
	ps := make([]Point2f, n)
	for i := range ps {
		ps[i] = randPoint2f(r)
	}
	return NewIndivid(ps, distance)
}

func RandIndividBySeed(seed int64, n int, distance float64) *Individ {
	r := newRandSeed(seed)
	ps := make([]Point2f, n)
	for i := range ps {
		ps[i] = randPoint2f(r)
	}
	return NewIndivid(ps, distance)
}

func RandIndividGridBySeed(seed int64, n int, distance float64) *Individ {
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
	return NewIndivid(ps, distance)
}

func (a *Individ) Clone() *Individ {
	if a != nil {
		return &Individ{
			ps:       clonePoint2fSlice(a.ps),
			distance: a.distance,
			fitness:  a.fitness,
		}
	}
	return nil
}

func (a *Individ) Fitness() float64 {
	return a.fitness
}

// MSE - Mean squared error
// https://en.wikipedia.org/wiki/Mean_squared_error
func (d *Individ) calcFitnessMSE() {

	x0 := d.distance

	n := len(d.ps)
	var sum float64
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			x := Distance(d.ps[i], d.ps[j])
			delta := (x - x0)
			sum += delta * delta
		}
	}

	m := n * (n - 1) / 2
	mse := sum / float64(m)

	d.fitness = mse
}

func (d *Individ) calcFitness() {
	d.calcFitnessMSE()
}

func (d *Individ) Range(f func(i int, p Point2f) bool) {
	for i, p := range d.ps {
		if !f(i, p) {
			return
		}
	}
}

func (d *Individ) PointsChan() <-chan Point2f {
	ps := make(chan Point2f)
	go func() {
		for _, p := range d.ps {
			ps <- p
		}
		close(ps)
	}()
	return ps
}

func (d *Individ) Random(r *rand.Rand) {
	if randBool(r) {
		d.randomFull(r)
	} else {
		d.randomShort(r)
	}
}

func (d *Individ) randomFull(r *rand.Rand) {
	i := r.Intn(len(d.ps))
	d.ps[i] = randPoint2f(r)
	d.calcFitness()
}

func (d *Individ) randomShort(r *rand.Rand) {

	delta := 0.01
	halfDelta := delta / 2

	dx := halfDelta - r.Float64()*delta
	dy := halfDelta - r.Float64()*delta

	i := r.Intn(len(d.ps))
	p := d.ps[i]

	p.X = clampFloat64(p.X+dx, 0, 1)
	p.Y = clampFloat64(p.Y+dy, 0, 1)

	d.ps[i] = p

	d.calcFitness()
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
