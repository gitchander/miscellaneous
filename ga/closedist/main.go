package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/fogleman/gg"
)

var errInvalidSize = errors.New("invalid size format")

func main() {
	run()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func run() {

	var (
		points         int
		size           string
		populationSize int
		mutations      int
		generations    int
		initSeed       int64
		distance       float64
	)

	flag.IntVar(&points, "points", 10, "number of points")
	flag.StringVar(&size, "size", "512x512", "image size")
	flag.IntVar(&populationSize, "population", 50, "population size")
	flag.IntVar(&mutations, "mutations", 30, "number of mutations")
	flag.IntVar(&generations, "generations", 1000, "number of generations")
	flag.Int64Var(&initSeed, "init-seed", -1, "initial seed")
	flag.Float64Var(&distance, "distance", 0.1, "distance")

	flag.Parse()

	imageSize, err := parseSize(size)
	checkError(err)

	start := time.Now()

	if initSeed < 0 {
		initSeed = time.Now().UnixNano()
		fmt.Printf("init-seed %d\n", initSeed)
	}

	r := newRandSeed(initSeed)

	startIndivid := RandIndivid(r, points, distance)
	//startIndivid := RandIndividGridBySeed(seed, points, distance)
	fmt.Println("begin-fitness:", startIndivid.Fitness())

	generation := make([]*Individ, 0, (populationSize + mutations))
	generation = append(generation, startIndivid)

	destFilename := "result.png"

	err = drawIndivid(imageSize, startIndivid.Range, destFilename)
	checkError(err)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for i := 0; i < generations; i++ {

		// Mutations
		ng := len(generation)

		for j := 0; j < mutations; j++ {

			k := r.Intn(ng)
			m := generation[k].Clone()

			//----------------------------------------------
			m.Random(r)
			//----------------------------------------------

			generation = append(generation, m)
		}

		//-------------------------------------------------------------------
		// Selection

		byFitness(generation).Sort()

		if len(generation) > populationSize {
			generation = generation[:populationSize]
		}

		//-------------------------------------------------------------------

		select {
		case <-(ticker.C):
			{
				bestIndivid := generation[0]
				err = drawIndivid(imageSize, bestIndivid.Range, destFilename)
				checkError(err)
				fmt.Println("fitness:", bestIndivid.Fitness())
			}
		default:
		}
	}

	bestIndivid := generation[0]
	err = drawIndivid(imageSize, bestIndivid.Range, destFilename)
	checkError(err)

	fmt.Println("end-fitness:", bestIndivid.Fitness())
	fmt.Println("work duration:", time.Since(start))
}

func parseSize(s string) (image.Point, error) {
	var zp image.Point
	vs := strings.Split(s, "x")
	if len(vs) != 2 {
		return zp, errInvalidSize
	}
	x, err := strconv.Atoi(vs[0])
	if err != nil {
		return zp, fmt.Errorf("invalid size format: %s", err)
	}
	y, err := strconv.Atoi(vs[1])
	if err != nil {
		return zp, fmt.Errorf("invalid size format: %s", err)
	}
	p := image.Point{
		X: x,
		Y: y,
	}
	return p, nil
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

func drawIndivid(size image.Point, rangePoints func(f func(i int, p Point2f) bool),
	filename string) error {

	dc := gg.NewContext(size.X, size.Y)

	var (
		clBackground   = ColorRGB{R: 1, G: 1, B: 1}
		clCircleFill   = ColorRGB{R: 0.5, G: 0.8, B: 0.5}
		clCircleStroke = ColorRGB{R: 0, G: 0.5, B: 0}
	)

	setColorRGB(dc, clBackground)
	dc.Clear()

	dc.SetRGB(0.2, 0.3, 0.8)
	dc.Translate(float64(size.X)/2, float64(size.Y)/2)

	sw := minFloat64(float64(size.X), float64(size.Y)) / 2
	dc.Scale(sw, sw)

	const (
		sk = 0.8
		//sk = 1.0
	)
	dc.Scale(sk, sk)

	dc.SetLineWidth(sw * 0.01)

	// dc.NewSubPath()
	// rangePoints(
	// 	func(i int, p Point2f) bool {
	// 		dc.LineTo(p.X, p.Y)
	// 		return true
	// 	})
	// dc.ClosePath()

	// const fillPath = true
	// if fillPath {
	// 	dc.SetRGB(0.4, 0.6, 0.9)
	// 	dc.FillPreserve()
	// }
	// dc.SetRGB(0.2, 0.3, 0.8)
	// dc.Stroke()

	//
	var (
		//circleRadius = 0.006
		//circleRadius = 0.008
		circleRadius = 0.02
	)

	var (
		minX, maxX float64
		minY, maxY float64
	)

	rangePoints(
		func(i int, p Point2f) bool {
			if i == 0 {
				minX, maxX = p.X, p.X
				minY, maxY = p.Y, p.Y
			} else {
				minX = minFloat64(minX, p.X)
				maxX = maxFloat64(maxX, p.X)

				minY = minFloat64(minY, p.Y)
				maxY = maxFloat64(maxY, p.Y)
			}
			return true
		},
	)

	center := Point2f{
		X: middle(minX, maxX),
		Y: middle(minY, maxY),
	}

	var (
		w  = maxFloat64(maxX-minX, maxY-minY)
		hw = w / 2
	)

	rangePoints(
		func(i int, p Point2f) bool {

			var (
				x = (p.X - center.X) / hw
				y = (p.Y - center.Y) / hw
			)

			dc.DrawCircle(x, y, circleRadius)
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
