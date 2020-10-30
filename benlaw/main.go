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
	"math/rand"
	"time"
)

// Benford's law
// https://en.wikipedia.org/wiki/Benford%27s_law

func main() {
	test1()
	//test2()
}

var colorTable = []color.Color{
	rgbf(hsv(0, 1, 1)),
	rgbf(hsv(35, 1, 1)),
	rgbf(hsv(70, 1, 1)),
	rgbf(hsv(105, 1, 1)),
	rgbf(hsv(140, 1, 1)),
	rgbf(hsv(175, 1, 1)),
	rgbf(hsv(210, 1, 1)),
	rgbf(hsv(245, 1, 1)),
	rgbf(hsv(280, 1, 1)),
}

func hsv(h, s, v float64) (r, g, b float64) {

	var (
		Hp = h / 60.0
		C  = v * s
		X  = C * (1 - math.Abs(math.Mod(Hp, 2)-1))
	)

	switch {
	case (0 <= Hp) && (Hp < 1):
		r = C
		g = X
	case (1 <= Hp) && (Hp < 2):
		r = X
		g = C
	case (2 <= Hp) && (Hp < 3):
		g = C
		b = X
	case (3 <= Hp) && (Hp < 4):
		g = X
		b = C
	case (4 <= Hp) && (Hp < 5):
		r = X
		b = C
	case (5 <= Hp) && (Hp < 6):
		r = C
		b = X
	}

	m := v - C

	r += m
	g += m
	b += m

	return
}

func rgbf(r, g, b float64) color.Color {
	return color.RGBA{
		R: colorChanPrepare(r),
		G: colorChanPrepare(g),
		B: colorChanPrepare(b),
		A: 255,
	}
}

func colorChanPrepare(x float64) uint8 {
	if x > 1 {
		x = 1
	}
	if x < 0 {
		x = 0
	}
	return uint8(math.Floor(x * 255))
}

func test2() {
	var (
		Dx = 800
		Dy = 400
	)
	m := image.NewRGBA(image.Rect(0, 0, Dx, Dy))

	draw.Draw(m, m.Bounds(), &image.Uniform{C: color.White}, image.ZP, draw.Over)

	var ds [9]int
	for i := 1; i <= 1000000; i++ {

		number := i

		dig := lastDigit(number)
		for j := range ds {
			if j+1 == dig {
				ds[j]++
			}
		}

		x := int(math.Floor(130 * math.Log10(float64(number))))

		for j, d := range ds {
			y := d * Dy / i
			m.Set(x, Dy-y, colorTable[j])
		}
	}

	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("test.png", buf.Bytes(), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func test1() {
	var stat [10]int
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	for i := 1; i < 10000000; i++ {
		stat[lastDigit(randNumber(r))]++
	}
	var sum int
	for _, d := range stat {
		sum += d
	}
	for i, d := range stat {
		if i > 0 {
			fmt.Printf("%d: %.2f%%\n", i, float64(d)*100/float64(sum))
		}
	}
}

func lastDigit(x int) (dig int) {
	for x > 0 {
		x, dig = quoRem(x, 10)
	}
	return
}

func quoRem(x, y int) (quo, rem int) {
	quo = x / y
	rem = x - quo*y
	return
}

func randNumber(r *rand.Rand) int {
	for {
		d := int(r.Uint32() >> uint(r.Intn(31)))
		if d != 0 {
			return d
		}
	}
}

//func randNumber2(r *rand.Rand) int {
//	n := 1 + r.Intn(10)
//	d := 1
//	for i := 0; i < n; i++ {
//		d *= 10
//	}
//	return r.Intn(d)
//}
