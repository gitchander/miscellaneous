package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	sorta "github.com/gitchander/miscellaneous/sort_algo"
	"github.com/gitchander/miscellaneous/sort_algo/util"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

func main() {
	//calcAndPrint()
	calcAndPlot()
}

func calcAndPrint() {

	var samples = []sample{
		{
			name: "BubbleSort",
			f:    sorta.BubbleSort,
			ps: []para{
				{10, 1000},
				{100, 100},
				{1000, 100},
				{10000, 10},
			},
		},
		{
			name: "CombSort",
			f:    sorta.CombSort,
			ps: []para{
				{10, 1000},
				{100, 1000},
				{1000, 100},
				{10000, 100},
				{100000, 10},
				{1000000, 10},
			},
		},
		{
			name: "QuickSort",
			f:    sorta.QuickSort,
			ps: []para{
				{10, 1000},
				{100, 1000},
				{1000, 100},
				{10000, 100},
				{100000, 10},
				{1000000, 10},
			},
		},
	}

	r := util.RandomerTimeNow()
	for _, s := range samples {
		fmt.Println(s.name, ":")
		for _, p := range s.ps {
			d := calc(p, s.f, r)
			fmt.Println(p.n, d)
		}
	}
}

func calcAndPlot() {

	var samples = []sample{
		{
			name: "BubbleSort",
			f:    sorta.BubbleSort,
			ps: []para{
				{10, 100},
				{50, 100},
				{100, 100},
				{200, 100},
				{300, 100},
				{400, 100},
				{500, 10},
				{600, 10},
				{700, 10},
				{800, 10},
				{900, 10},
				{1000, 10},
			},
		},
		{
			name: "InsertionSort",
			f:    sorta.InsertionSort,
			ps: []para{
				{10, 100},
				{50, 100},
				{100, 100},
				{200, 100},
				{300, 100},
				{400, 100},
				{500, 10},
				{600, 10},
				{700, 10},
				{800, 10},
				{900, 10},
				{1000, 10},
				{1200, 10},
				{1400, 10},
			},
		},
		{
			name: "CombSort",
			f:    sorta.CombSort,
			ps: []para{
				{10, 1000},
				{100, 1000},
				{500, 100},
				{1000, 100},
				{5000, 10},
			},
		},
		{
			name: "ShellSort",
			f:    sorta.ShellSort,
			ps: []para{
				{10, 1000},
				{100, 1000},
				{500, 100},
				{1000, 100},
				{2000, 10},
				{5000, 10},
			},
		},
		{
			name: "QuickSort",
			f:    sorta.QuickSort,
			ps: []para{
				{10, 1000},
				{100, 1000},
				{500, 100},
				{1000, 100},
				{3000, 10},
				{5000, 10},
			},
		},
	}

	p := plot.New()

	p.X.Label.Text = "N"
	p.Y.Label.Text = "Compare count"

	var ps []plot.Plotter

	type item struct {
		name  string
		value [2]plot.Thumbnailer
	}
	var items []item

	r := util.RandomerTimeNow()
	for i, sample := range samples {
		xys := calcXYs(&sample, r)
		l, s, err := plotter.NewLinePoints(xys)
		checkError(err)
		l.Color = plotutil.Color(i)
		l.Dashes = plotutil.Dashes(i)
		s.Color = plotutil.Color(i)
		s.Shape = plotutil.Shape(i)
		ps = append(ps, l, s)
		items = append(items,
			item{
				name:  sample.name,
				value: [2]plot.Thumbnailer{l, s},
			})
	}

	p.Add(ps...)
	for _, item := range items {
		v := item.value[:]
		p.Legend.Add(item.name, v[0], v[1])
	}

	wt, err := p.WriterTo(512, 512, "png")
	checkError(err)

	var buf bytes.Buffer
	_, err = wt.WriteTo(&buf)
	checkError(err)

	err = ioutil.WriteFile("result.png", buf.Bytes(), 0644)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func calcXYs(s *sample, r util.Randomer) (xys plotter.XYs) {
	for _, p := range s.ps {
		d := calc(p, s.f, r)
		xys = append(xys,
			struct{ X, Y float64 }{
				X: float64(p.n),
				Y: float64(d),
			})
	}
	return
}

type sample struct {
	name string
	f    func(sort.Interface)
	ps   []para
}

type para struct {
	n, k int
}

func calc(p para, f func(sort.Interface), r util.Randomer) int {
	var sum int
	for i := 0; i < p.k; i++ {
		sum += step(p.n, f, r)
	}
	return sum / p.k
}

func step(n int, f func(sort.Interface), r util.Randomer) (count int) {

	var a = make([]int, n)
	for i := range a {
		a[i] = i
	}
	util.Scramble(r, sort.IntSlice(a))

	data := &intSlice{vs: a}

	f(data)

	if !sort.IsSorted(sort.IntSlice(a)) {
		panic("array isn't sorted!")
	}

	return data.count
}

type intSlice struct {
	vs    []int
	count int
}

func (p *intSlice) Len() int {
	return len(p.vs)
}

func (p *intSlice) Less(i, j int) bool {
	p.count++
	vs := p.vs
	return vs[i] < vs[j]
}

func (p *intSlice) Swap(i, j int) {
	vs := p.vs
	vs[i], vs[j] = vs[j], vs[i]
}
