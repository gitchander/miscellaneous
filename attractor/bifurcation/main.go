package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"time"
)

// https://en.wikipedia.org/wiki/Bifurcation_diagram
// https://en.wikipedia.org/wiki/Logistic_map

func main() {

	fmt.Println("Rendering Bifurcation diagram")

	n := 100000

	// var (
	// 	size = image.Pt(4096, 2048)
	// 	lK   = Interval{Min: -2.5, Max: +4.5}
	// 	lV   = Interval{Min: -1, Max: 2}
	// )

	// var (
	// 	size = image.Pt(4096, 2048)
	// 	lK   = Interval{Min: -2, Max: +4}
	// 	lV   = Interval{Min: -0.55, Max: 1.55}
	// )

	var (
		size = image.Pt(1024, 1024)
		lK   = Interval{Min: 2.95, Max: +4}
		lV   = Interval{Min: -0.05, Max: 1.05}
	)

	// var (
	// 	size = image.Pt(1024, 1024)
	// 	lK   = Interval{Min: 3.5, Max: +4}
	// 	lV   = Interval{Min: -0.05, Max: 1.05}
	// )

	//--------------------------------------------------------------------------
	lfX := Interval{Min: 0, Max: float64(size.X - 1)}
	lfY := Interval{Min: 0, Max: float64(size.Y - 1)}
	lv0 := Interval{Min: 0, Max: 1}

	fss := make([][]float64, size.X)
	for y := range fss {
		fss[y] = make([]float64, size.Y)
	}

	var (
		subX = 10
		dfX  = 1.0 / float64(subX)
	)

	r := newRandNow()

	for x := 0; x < size.X; x++ {

		fX0 := float64(x)

		for j := 0; j < subX; j++ {

			fX := fX0 + float64(j)*dfX

			var (
				// fX = float64(x)

				nfX = NormalizeIvl(lfX, fX)
				k   = LerpIvl(lK, nfX)
			)

			lmf := LogisticMapFunc(k)

			v := randValueIn(r, lv0)

			for i := 0; i < n; i++ {

				v = lmf(v)

				var (
					nv = NormalizeIvl(lV, v)
					fY = LerpIvl(lfY, nv)
				)

				if false {
					y := int(math.Round(fY))
					if (0 <= y) && (y < size.Y) {
						fss[x][y] += 1
					}
				} else {
					fY0, fY1 := floorCeil(fY)

					var (
						t0 = 1 - (fY - fY0)
						t1 = 1 - (fY1 - fY)
					)

					var (
						y0 = int(fY0)
						y1 = int(fY1)
					)

					if (0 <= y0) && (y0 < size.Y) {
						fss[x][y0] += t0
					}
					if (0 <= y1) && (y1 < size.Y) {
						fss[x][y1] += t1
					}
				}
			}
		}
	}

	if true {
		for _, fs := range fss {
			var vmax float64
			for _, x := range fs {
				vmax = maxFloat64(vmax, x)
			}
			// normalize
			for i := range fs {
				fs[i] /= vmax
			}
		}
	} else {
		var vmax float64
		for _, fs := range fss {
			for _, x := range fs {
				vmax = maxFloat64(vmax, x)
			}
		}
		// normalize
		for _, fs := range fss {
			for i := range fs {
				fs[i] /= vmax
			}
		}
	}

	//toneCorrection := gammaCorrectionFunc(0.3)
	toneCorrection := toneCorrectionFunc(5)

	m := image.NewRGBA(image.Rectangle{Max: size})

	for x, fs := range fss {
		for y, v := range fs {

			v = clamp(v, 0, 1)
			v = toneCorrection(v)
			v = 1 - v

			Y := uint16(math.Round(v * math.MaxUint16))

			m.Set(x, size.Y-1-y, color.Gray16{Y: Y})
		}
	}

	err := saveImagePNG("result.png", m)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Logistic map:
// https://en.wikipedia.org/wiki/Logistic_map

func LogisticMapFunc(r float64) func(float64) float64 {
	return func(x float64) float64 {
		return r * x * (1 - x)
	}
}

type Interval struct {
	Min, Max float64
}

func newRandNow() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randValueIn(r *rand.Rand, l Interval) float64 {
	return LerpIvl(l, r.Float64())
}

func LerpIvl(l Interval, t float64) float64 {
	return lerp(l.Min, l.Max, t)
}

func lerp(v0, v1 float64, t float64) float64 {
	return v0*(1-t) + v1*t
}

func NormalizeIvl(l Interval, x float64) float64 {
	return normalize(l.Min, l.Max, x)
}

func normalize(min, max float64, x float64) float64 {
	return (x - min) / (max - min)
}

func floorCeil(x float64) (floor, ceil float64) {
	floor = math.Floor(x)
	ceil = floor + 1
	return
}

func maxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func clamp(a float64, min, max float64) float64 {
	if a < min {
		a = min
	}
	if a > max {
		a = max
	}
	return a
}

func saveImagePNG(filename string, m image.Image) error {
	var b bytes.Buffer
	err := png.Encode(&b, m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b.Bytes(), 0644)
}

func gammaCorrectionFunc(gamma float64) func(float64) float64 {
	return func(inputValue float64) (outputValue float64) {
		outputValue = math.Pow(inputValue, gamma)
		return
	}
}

func toneCorrectionFunc(toneFactor float64) func(float64) float64 {
	factor := toneFactor
	return func(x float64) float64 {
		if x < 0 {
			return 0
		}
		if x < 1 {
			return 1 - math.Exp(factor*math.Log(1-x))
		}
		return 1
	}
}
