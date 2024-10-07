package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"math"

	"github.com/fogleman/gg"

	"github.com/gitchander/miscellaneous/utils/random"
)

// http://algorithmicbotany.org/papers/abop/abop-ch4.pdf
// https://www.youtube.com/watch?v=KWoJgHFYWxY

func main() {
	//drawPoints()
	drawAreas()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sqr(a float64) float64 {
	return a * a
}

func squareDistance(a, b Point2f) float64 {
	return sqr(a.X-b.X) + sqr(a.Y-b.Y)
}

// Vogel’s formula
// α - angle
func formulaV(n int, c, α float64) Polar {

	nf := float64(n)

	return Polar{
		Rho: c * math.Sqrt(nf),
		Phi: α * nf,
	}
}

func drawPoints() {

	var (
		//angleDeg = 137.5
		angleDeg = 360 / sqr(math.Phi)
	)

	var (
		angleRad = DegToRad(angleDeg)
		c        = 3.0
		n        = 20000
	)
	//log.Println(360 / sqr(math.Phi))

	//var sizeX, sizeY = 512, 512
	//var sizeX, sizeY = 1024, 1024
	//var sizeX, sizeY = 900, 900
	size := image.Point{X: 900, Y: 900}

	dc := gg.NewContext(size.X, size.Y)

	dc.DrawRectangle(0, 0, float64(size.X), float64(size.Y))
	//dc.SetRGB(0, 0, 0) // black
	dc.SetRGB(1, 1, 1) // white
	dc.Fill()

	center := Pt2f(float64(size.X), float64(size.Y)).DivScalar(2)

	circleRadius := c * 0.5

	for i := 0; i < n; i++ {
		p := formulaV(i, c, angleRad)
		t := PolarToCartesian(p)
		t = t.Add(center)
		dc.DrawCircle(t.X, t.Y, circleRadius)
	}

	dc.SetRGB(0, 0, 0) // black
	//dc.SetRGB(1, 1, 1) // white
	dc.Fill()

	err := dc.SavePNG("points.png")
	checkError(err)
}

func imageFill(m draw.Image, c color.Color) {
	draw.Draw(m, m.Bounds(), image.NewUniform(c), image.ZP, draw.Src)
}

func drawAreas() {

	var (
		angleRad = goldAngleRad

		// c = 10.0
		// n = 1500

		c = 30.0
		n = 200

		// c = 48.0
		// n = 200
	)

	//size := image.Point{X: 512, Y: 512}
	//size := image.Point{X: 1024, Y: 1024}
	size := image.Point{X: 650, Y: 650}
	ir := image.Rect(0, 0, size.X, size.Y)
	m := image.NewRGBA(ir)

	var (
		//fillColor = color.White
		fillColor = color.Gray{Y: 255}
	)
	//imageFill(m, fillColor)

	// err := WriteImagePNG("areas.png", m)
	// checkError(err)
	// return

	center := Pt2f(float64(size.X), float64(size.Y)).DivScalar(2)

	var ps []Point2f

	for i := 0; i < n; i++ {
		p := formulaV(i, c, angleRad)
		t := PolarToCartesian(p)
		t = t.Add(center)
		ps = append(ps, t)
	}

	var seed int64
	if false {
		seed = 7788002435780809846
	} else {
		seed = random.NextSeed()
		fmt.Println("random seed", seed)
	}
	pal := randPalette(6, seed)

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			v := Pt2f(float64(x), float64(y))
			index := 0
			d := squareDistance(v, ps[index])
			for i, p := range ps {
				di := squareDistance(p, v)
				if di < d {
					index = i
					d = di
				}
			}
			var (
				//cond = index < 25
				cond = index < 90
				//cond = index < (25 + 4*13)
				//cond = (25 <= index) && (index < (25 + 3*13))
				//cond = (25 <= index) && (index < (25 + 5*13))
				//cond = true
			)

			if cond {
				m.Set(x, y, pal[index%len(pal)])
			} else {
				m.Set(x, y, fillColor)
			}
		}
	}
	err := WriteImagePNG("areas.png", m)
	checkError(err)
}

func WriteImagePNG(filename string, m image.Image) error {
	var b bytes.Buffer
	err := png.Encode(&b, m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b.Bytes(), 0644)
}
