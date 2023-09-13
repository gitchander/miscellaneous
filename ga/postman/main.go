package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"path/filepath"
	"time"

	"github.com/fogleman/gg"
)

// https://pdfs.semanticscholar.org/3249/7c1c6c2783d39caf31451a9ef2ffec16ee16.pdf
// https://www.sciencedirect.com/science/article/pii/030505489400070O

func main() {
	postmanProblem()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func postmanProblem() {

	var (
		dir                string
		numberOfPoints     int
		initSeed           int64
		generations        int
		mutations          int
		generationCapacity int
	)

	flag.StringVar(&dir, "dir", "images", "output directory")
	flag.IntVar(&numberOfPoints, "points", 100, "number of points")
	flag.Int64Var(&initSeed, "init-seed", -1, "initial seed")
	flag.IntVar(&generations, "generations", 1000, "number of generations")
	flag.IntVar(&mutations, "mutations", 50, "number of mutations")
	flag.IntVar(&generationCapacity, "gen-cap", 50, "generation capacity")

	flag.Parse()

	r := newRandNow()

	if initSeed < 0 {
		initSeed = r.Int63()
		fmt.Printf("init-seed %d\n", initSeed)
	}

	err := MkdirIfNotExist(dir)
	checkError(err)

	start := time.Now()

	startIndivid := RandIndividBySeed(initSeed, numberOfPoints)
	//startIndivid := RandIndividGridBySeed(initSeed, 20, 20)
	fmt.Println("begin-fitness:", startIndivid.Fitness())

	generation := make([]*Individ, 0, (generationCapacity + mutations))
	generation = append(generation, startIndivid)

	err = drawIndivid(startIndivid.Range, filepath.Join(dir, "source.png"))
	checkError(err)

	for i := 0; i < generations; i++ {

		ng := len(generation)

		for j := 0; j < mutations; j++ {

			k := r.Intn(ng)
			m := generation[k].Clone()

			//----------------------------------------------
			//m.RandomExchange(r)
			//m.RandomInsertion(r)
			m.RandomFlip(r)
			//----------------------------------------------

			generation = append(generation, m)
		}

		byFitness(generation).Sort()

		if len(generation) > generationCapacity {
			generation = generation[:generationCapacity]
		}
	}

	bestIndivid := generation[0]
	err = drawIndivid(bestIndivid.Range, filepath.Join(dir, "genetic.png"))
	checkError(err)

	fmt.Println("end-fitness:", bestIndivid.Fitness())
	fmt.Println("work duration:", time.Since(start))
}

func drawField() {
	r := newRandNow()

	ps := make([]Point2f, 10)
	for i := range ps {
		ps[i] = randPoint2f(r)
	}

	fmt.Println(len(ps))

	const (
		dx = 512
		dy = 512
	)

	dc := gg.NewContext(dx, dy)

	// fill background
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0.2, 0.3, 0.8)
	dc.SetLineWidth(3)
	dc.Scale(dx, dy)

	dc.NewSubPath()
	for _, p := range ps {
		dc.LineTo(p.X, p.Y)

		// if i == 0 {
		// 	dc.MoveTo(p.X, p.Y)
		// } else {
		// 	dc.LineTo(p.X, p.Y)
		// }
	}
	dc.ClosePath()
	dc.Stroke()

	for _, p := range ps {

		dc.DrawCircle(p.X, p.Y, 0.01)
		dc.SetLineWidth(2.0)

		dc.SetRGB(0.5, 0.8, 0.5)
		dc.FillPreserve()

		dc.SetRGB(0, 0.5, 0)
		dc.Stroke()
	}

	dc.SavePNG("test.png")
}

func drawIndivid(rangePoints func(f func(i int, p Point2f) bool), filename string) error {

	//size := image.Point{X: 100, Y: 100}
	//size := image.Point{X: 256, Y: 256}
	size := image.Point{X: 512, Y: 512}

	dc := gg.NewContext(size.X, size.Y)

	var (
		clBackground   = ColorRGB{R: 1, G: 1, B: 1}
		clCircleFill   = ColorRGB{R: 0.5, G: 0.8, B: 0.5}
		clCircleStroke = ColorRGB{R: 0, G: 0.5, B: 0}
	)

	setColorRGB(dc, clBackground)
	dc.Clear()

	dc.SetRGB(0.2, 0.3, 0.8)
	dc.Scale(float64(size.X), float64(size.Y))
	dc.SetLineWidth(2)

	dc.NewSubPath()
	rangePoints(
		func(i int, p Point2f) bool {
			dc.LineTo(p.X, p.Y)
			return true
		})
	dc.ClosePath()

	const fillPath = true
	if fillPath {
		dc.SetRGB(0.4, 0.6, 0.9)
		dc.FillPreserve()
	}
	dc.SetRGB(0.2, 0.3, 0.8)
	dc.Stroke()

	const (
		// circleRadius = 0.01
		// circleRadius = 0.006
		circleRadius = 0.008
	)

	rangePoints(
		func(i int, p Point2f) bool {
			dc.DrawCircle(p.X, p.Y, circleRadius)
			//dc.SetLineWidth(2)

			// if i == 0 {
			// 	setColorRGB(dc, clRed)
			// } else {
			// 	setColorRGB(dc, clCircleFill)
			// }

			setColorRGB(dc, clCircleFill)
			dc.FillPreserve()
			setColorRGB(dc, clCircleStroke)
			dc.Stroke()

			return true
		})

	return dc.SavePNG(filename)
}

func testIntersection() {
	r := newRandNow()
	var (
		a0 = randPoint2f(r)
		a1 = randPoint2f(r)
		b0 = randPoint2f(r)
		b1 = randPoint2f(r)
	)

	size := image.Point{X: 512, Y: 512}

	dc := gg.NewContext(size.X, size.Y)

	p, ok := Intersection(a0, a1, b0, b1)
	fmt.Println(ok)

	var (
		clBackground   = ColorRGB{R: 1, G: 1, B: 1}
		clCircleFill   = ColorRGB{R: 0.5, G: 0.8, B: 0.5}
		clCircleStroke = ColorRGB{R: 0, G: 0.5, B: 0}
	)

	setColorRGB(dc, clBackground)
	dc.Clear()

	dc.SetRGB(0.2, 0.3, 0.8)
	dc.SetLineWidth(3)
	dc.Scale(float64(size.X), float64(size.Y))

	dc.MoveTo(a0.X, a0.Y)
	dc.LineTo(a1.X, a1.Y)

	dc.MoveTo(b0.X, b0.Y)
	dc.LineTo(b1.X, b1.Y)

	dc.Stroke()

	drawPoint(dc, a0, clCircleFill, clCircleStroke)
	drawPoint(dc, a1, clCircleFill, clCircleStroke)
	drawPoint(dc, b0, clCircleFill, clCircleStroke)
	drawPoint(dc, b1, clCircleFill, clCircleStroke)

	if ok {
		drawPoint(dc, p, clRed, clBlack)
	}

	dc.SavePNG("intersection_test.png")
}

func drawPoint(dc *gg.Context, p Point2f, clCircleFill, clCircleStroke ColorRGB) {

	dc.DrawCircle(p.X, p.Y, 0.01)
	dc.SetLineWidth(2.0)

	setColorRGB(dc, clCircleFill)
	dc.FillPreserve()

	setColorRGB(dc, clCircleStroke)
	dc.Stroke()
}
