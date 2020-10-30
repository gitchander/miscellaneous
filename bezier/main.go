package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {

	var c Config

	flag.StringVar(&(c.Filename), "filename", "result.png", "result image file name (*.png)")
	flag.Int64Var(&(c.Seed), "seed", -1, "seed")
	flag.StringVar(&(c.ColorBG), "bg", "#fff", "background color")
	flag.StringVar(&(c.ColorFG), "fg", "random", "foreground color")
	flag.IntVar(&(c.Points), "points", 30, "number of points")
	flag.IntVar(&(c.Curves), "curves", 1, "number of curves")

	flag.Parse()

	err := runConfig(c)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func runConfig(c Config) error {

	if c.Seed == -1 {
		c.Seed = randSeed()
		fmt.Println("seed:", c.Seed)
	}

	r := rand.New(rand.NewSource(c.Seed))

	var clBackground color.Color
	if c.ColorBG == "random" {
		clBackground = randColor(r)
		fmt.Println("background color:", clBackground)
	} else {
		c, err := parseHexColor(c.ColorBG)
		if err != nil {
			return err
		}
		clBackground = c
	}

	var clForeground color.Color
	var randomForeground bool
	if c.ColorFG == "random" {
		randomForeground = true
	} else {
		c, err := parseHexColor(c.ColorFG)
		if err != nil {
			return err
		}
		clForeground = c
	}

	size := image.Point{X: 512, Y: 512}
	m := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))

	src := image.NewUniform(clBackground)
	draw.Draw(m, m.Bounds(), src, image.ZP, draw.Src)

	bcs := calcBC(c.Points)

	for i := 0; i < c.Curves; i++ {

		if randomForeground {
			clForeground = randColor(r)
		}

		ps := make([]image.Point, c.Points)
		var bounds = m.Bounds()
		for i := range ps {
			ps[i] = image.Point{
				X: r.Intn(bounds.Dx()),
				Y: r.Intn(bounds.Dy()),
			}
		}

		drawCurve(m, bcs, ps, clForeground)
	}

	return saveImage(m, c.Filename)
}

type Config struct {
	Filename string
	Seed     int64
	ColorBG  string
	ColorFG  string
	Points   int // max 67 for bin koef element uint64
	Curves   int
}

func saveImage(m image.Image, filename string) error {

	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = buf.WriteTo(file)
	return err
}

func drawCurve(m *image.RGBA, bcs []int, ps []image.Point, c color.Color) {
	const (
		tn     = 10000
		t_coef = 1 / float64(tn-1)
	)
	for i := 0; i < tn; i++ {
		t := float64(i) * t_coef
		p := calcBezier(ps, bcs, t)
		m.Set(p.X, p.Y, c)
	}
}

// Binomial coefficients
func calcBC(n int) []int {
	bcs := make([]int, n)
	var m = 1
	for i := 0; i < n; i++ {
		b := 1
		bcs[0] = 1
		for k := 1; k < m; k++ {
			s := 0
			if k < m-1 {
				s = bcs[k]
			}
			bcs[k] = b + s
			b = s
		}
		m++
	}
	return bcs
}

func calcBezier(ps []image.Point, bcs []int, t float64) image.Point {

	if t <= 0 {
		return ps[0]
	}
	if t >= 1 {
		return ps[len(ps)-1]
	}

	n := len(ps)

	var (
		a = 1.0
		b = power(1-t, n-1)

		av = t
		bv = 1 / (1 - t)
	)

	const useInt = true // !!!

	if useInt {
		var x, y int
		for i := 0; i < n; i++ {
			w := a * b * float64(bcs[i])
			x += round(w * float64(ps[i].X))
			y += round(w * float64(ps[i].Y))
			a *= av
			b *= bv
		}
		return image.Point{
			X: x,
			Y: y,
		}
	} else {
		var x, y float64
		for i := 0; i < n; i++ {
			w := a * b * float64(bcs[i])
			x += w * float64(ps[i].X)
			y += w * float64(ps[i].Y)
			a *= av
			b *= bv
		}
		return image.Point{
			X: round(x),
			Y: round(y),
		}
	}
}

func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}

// a ^ n
func power(a float64, n int) float64 {
	r := 1.0
	for i := 0; i < n; i++ {
		r *= a
	}
	return r
}

func randSeed() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63()
}
