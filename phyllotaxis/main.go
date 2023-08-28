package main

import (
	"bytes"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"io/ioutil"
	"log"
	"math"

	"github.com/fogleman/gg"
)

// http://algorithmicbotany.org/papers/abop/abop-ch4.pdf
// https://www.youtube.com/watch?v=KWoJgHFYWxY

const tau = 2 * math.Pi

const (
	goldAngleRad = tau / (math.Phi * math.Phi)
	goldAngleDeg = 360 / (math.Phi * math.Phi)
)

func main() {
	//drawPoints()
	drawAreas()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func DegToRad(deg float64) (rad float64) {
	return deg * (tau / 360)
}

func RadToDeg(rad float64) (deg float64) {
	return rad * (360 / tau)
}

func sqr(a float64) float64 {
	return a * a
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

func squareDistance(a, b Point2f) float64 {
	return sqr(a.X-b.X) + sqr(a.Y-b.Y)
}

func drawAreas() {

	var (
		angleRad = goldAngleRad

		// c        = 10.0
		// n        = 1500

		c = 40.0
		n = 200
	)

	//size := image.Point{X: 512, Y: 512}
	//size := image.Point{X: 1024, Y: 1024}
	size := image.Point{X: 900, Y: 900}
	r := image.Rect(0, 0, size.X, size.Y)
	m := image.NewRGBA(r)

	center := Pt2f(float64(size.X), float64(size.Y)).DivScalar(2)

	var ps []Point2f

	for i := 0; i < n; i++ {
		p := formulaV(i, c, angleRad)
		t := PolarToCartesian(p)
		t = t.Add(center)
		ps = append(ps, t)
	}

	var pal []color.Color
	switch 1 {
	case 0:
		pal = palette.Plan9
	case 1:
		pal = palette.WebSafe
	case 2:
		// pal = make([]color.Color, 6)
		// for i := range pal {
		// 	pal[i] = color.Gray{Y: 255}
		// }
		// pal[0] = color.Gray{Y: 0}

		pal = make([]color.Color, 6) // 6, 8, 13
		d := 255 / (len(pal))

		for i := range pal {
			pal[i] = color.Gray{Y: 255 - byte(i*d)}
		}
		//pal[0] = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	}
	pal = clonePalette(pal)
	shufflePalette(pal)
	pal = pal[:6]

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
				c := color.Gray{Y: 0}
				m.Set(x, y, c)
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

func clonePalette(a []color.Color) []color.Color {
	b := make([]color.Color, len(a))
	copy(b, a)
	return b
}
