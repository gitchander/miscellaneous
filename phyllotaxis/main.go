package main

import (
	"log"
	"math"

	"github.com/fogleman/gg"
)

// http://algorithmicbotany.org/papers/abop/abop-ch4.pdf
// https://www.youtube.com/watch?v=KWoJgHFYWxY

func main() {

	var (
		angle = DegToRad(360 / sqr(math.Phi))
		c     = 3.0
		n     = 20000
	)

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
		r, φ := formulaV(i, c, angle)
		x, y := polarToCartesian(r, φ)
		x += centerX
		y += centerY
		dc.DrawCircle(x, y, circleRadius)
	}

	dc.SetRGB(0, 0, 0) // black
	//dc.SetRGB(1, 1, 1) // white
	dc.Fill()

	err := dc.SavePNG("result.png")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const tau = 2 * math.Pi

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

// Vogel’s formula
// α - angle
func formulaV(n int, c, α float64) (r, φ float64) {

	nf := float64(n)

	r = c * math.Sqrt(nf)
	φ = α * nf

	return
}
