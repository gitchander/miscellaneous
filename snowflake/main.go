package main

import (
	"bytes"
	"flag"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
)

const tau = 2 * math.Pi

func main() {

	var p Params

	var (
		filename   string
		strSize    string
		angleMin   float64
		angleMax   float64
		angleSteps int
	)

	flag.StringVar(&filename, "filename", "result.png", "result image filename")
	flag.StringVar(&strSize, "size", "256x256", "image size")
	flag.IntVar(&(p.Order), "order", 15, "order")
	flag.IntVar(&(p.SmoothOrder), "smooth", 4, "smooth order")
	flag.Float64Var(&(p.Scale), "scale", 50, "Scale")
	flag.Float64Var(&(p.SF.Xf), "factor_x", 0.5, "factor x")
	flag.Float64Var(&(p.SF.Yf), "factor_y", -0.5, "factor y")
	flag.Float64Var(&(p.SF.Wf), "factor_w", 4.0, "factor w")
	flag.Float64Var(&(angleMin), "angle_min", 0.0, "angle min (in tau)")
	flag.Float64Var(&(angleMax), "angle_max", 1.0, "angle max (in tau)")
	flag.IntVar(&(angleSteps), "angle_steps", 1000000, "angle steps")

	flag.Parse()

	imageSize, err := parseSize(strSize)
	checkError(err)

	p.Size = imageSize

	angleMin *= tau
	angleMax *= tau

	p.Angle = Range{
		Min:  angleMin,
		Max:  angleMax,
		Step: (angleMax - angleMin) / float64(angleSteps),
	}

	g := image.NewGray(image.Rectangle{Max: p.Size})

	if p.SmoothOrder <= 1 {
		renderSimple(g, p)
	} else {
		renderSmooth(g, p)
	}

	err = saveImagePNG(g, filename)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Params struct {
	Size        image.Point
	Order       int
	SmoothOrder int
	Scale       float64
	Angle       Range
	SF          Snowflake
}

type Range struct {
	Min  float64
	Max  float64
	Step float64
}

type Snowflake struct {
	Xf float64
	Yf float64
	Wf float64
}

func saveImagePNG(m image.Image, filename string) error {
	var b bytes.Buffer
	err := png.Encode(&b, m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b.Bytes(), 0666)
}

func renderBase(p Params, set func(x, y float64)) {

	var n = p.Order

	var (
		xs = make([]float64, n)
		ys = make([]float64, n)
		ws = make([]float64, n)
	)

	dx, dy, dw := 1.0, 1.0, 1.0
	for i := 0; i < n; i++ {
		xs[i] = dx
		ys[i] = dy
		ws[i] = dw

		dx *= p.SF.Xf
		dy *= p.SF.Yf
		dw *= p.SF.Wf
	}

	angle := p.Angle.Min
	for angle < p.Angle.Max {

		x, y := 0.0, 0.0
		for i := 0; i < n; i++ {
			sin, cos := math.Sincos(ws[i] * angle)
			x += xs[i] * cos
			y += ys[i] * sin
		}

		set(x, y)

		angle += p.Angle.Step
	}
}

func renderSimple(g *image.Gray, p Params) {

	cg := color.Gray{Y: 255}

	var (
		scaleFactor = p.Scale
		center      = p.Size.Div(2)
	)

	set := func(x, y float64) {

		v := image.Point{
			X: round(scaleFactor * x),
			Y: round(scaleFactor * y),
		}.Add(center)

		g.Set(v.X, v.Y, cg)
	}

	renderBase(p, set)
}

func renderSmooth(g *image.Gray, p Params) {

	var (
		size   = p.Size
		m      = p.SmoothOrder
		sizeUp = size.Mul(m)
	)

	ssv := make([][]bool, sizeUp.X)
	for i := range ssv {
		ssv[i] = make([]bool, sizeUp.Y)
	}

	var (
		scaleFactor = float64(m) * p.Scale
		center      = sizeUp.Div(2)
	)

	set := func(x, y float64) {
		v := image.Point{
			X: round(scaleFactor * x),
			Y: round(scaleFactor * y),
		}.Add(center)

		if (0 <= v.X) && (v.X < sizeUp.X) {
			if (0 <= v.Y) && (v.Y < sizeUp.Y) {
				ssv[v.X][v.Y] = true
			}
		}
	}

	renderBase(p, set)

	mm := m * m
	for x := 0; x < size.X; x++ {
		mX := m * x
		for y := 0; y < size.Y; y++ {
			mY := m * y
			var sum int
			for dX := 0; dX < m; dX++ {
				iX := mX + dX
				for dY := 0; dY < m; dY++ {
					iY := mY + dY
					if ssv[iX][iY] {
						sum++
					}
				}
			}
			v := float64(sum) / float64(mm) // [0..1]
			gY := round(v * 255)            // [0..255]
			cg := color.Gray{Y: uint8(gY)}
			g.SetGray(x, y, cg)
		}
	}
}

func round(x float64) int {
	return int(math.Round(x))
}
