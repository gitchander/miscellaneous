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
	drawPoints()
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

// https://en.wikipedia.org/wiki/Polar_coordinate_system
func polarToCartesian(r, φ float64) (x, y float64) {
	sin, cos := math.Sincos(φ)
	x = r * cos
	y = r * sin
	return
}

func cartesianToPolar(x, y float64) (r, φ float64) {
	φ = math.Atan2(y, x)
	r = math.Hypot(x, y)
	return
}

// Vogel’s formula
// α - angle
func formulaV(n int, c, α float64) (r, φ float64) {

	nf := float64(n)

	r = c * math.Sqrt(nf)
	φ = α * nf

	return
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
	var sizeX, sizeY = 900, 900

	dc := gg.NewContext(sizeX, sizeY)

	dc.DrawRectangle(0, 0, float64(sizeX), float64(sizeY))
	//dc.SetRGB(0, 0, 0) // black
	dc.SetRGB(1, 1, 1) // white
	dc.Fill()

	var (
		centerX = float64(sizeX) / 2
		centerY = float64(sizeY) / 2
	)

	circleRadius := c * 0.5

	for i := 0; i < n; i++ {
		r, φ := formulaV(i, c, angleRad)
		x, y := polarToCartesian(r, φ)
		x += centerX
		y += centerY
		dc.DrawCircle(x, y, circleRadius)
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
		c        = 10.0
		n        = 1500
	)

	//size := image.Point{X: 512, Y: 512}
	//size := image.Point{X: 1024, Y: 1024}
	size := image.Point{X: 900, Y: 900}
	r := image.Rect(0, 0, size.X, size.Y)
	m := image.NewRGBA(r)

	center := Pt2f(float64(size.X), float64(size.Y)).DivScalar(2)

	var ps []Point2f

	for i := 0; i < n; i++ {
		r, φ := formulaV(i, c, angleRad)
		x, y := polarToCartesian(r, φ)

		p := center.Add(Pt2f(x, y))
		ps = append(ps, p)
	}

	var pal []color.Color
	switch 1 {
	case 0:
		pal = palette.Plan9
	case 1:
		pal = palette.WebSafe
	}
	pal = clonePalette(pal)
	shufflePalette(pal)

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
			m.Set(x, y, pal[index%len(pal)])
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
