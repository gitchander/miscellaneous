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

func main() {
	size := image.Pt(800, 800)
	r := image.Rectangle{Max: size}
	g := image.NewGray(r)
	draw.Draw(g, r, image.NewUniform(color.Gray{Y: 10}), image.ZP, draw.Src)

	render1(g, size)
	//render2(g, size)
	//render3(g, size)

	var buf bytes.Buffer
	err := png.Encode(&buf, g)
	checkError(err)
	err = ioutil.WriteFile("test.png", buf.Bytes(), 0666)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}

func newRandSeed(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func makePoints(n int) []Point2f {
	ps := make([]Point2f, n)
	var (
		radius = 1.0
		u0     = -math.Pi / 2
		du     = 2 * math.Pi / float64(n)
	)
	for i := range ps {
		u := u0 + float64(i)*du
		ps[i] = PolarToPoint2f(u, radius)
	}
	return ps
}

func render1(g *image.Gray, size image.Point) {

	var (
		//		n = 2
		//		t = 0.7

		n = 3
		t = 0.5

		// n = 4
		// t = 0.52

		// n = 5
		// t = 0.62

		// n = 6
		// t = 0.668

		// n = 7
		// t = 0.693

		// n = 12
		// t = 0.79

		// n = 3
		// t = 0.5
	)

	var nr Nexter = PtNext{
		t:  t,
		r:  newRand(),
		Ps: makePoints(n),
	}

	//radiusFactor := 1.0
	radiusFactor := 0.85

	var smooth = true

	if !smooth {
		render(g, size, nr, 10000, radiusFactor)
	} else {
		sc := SmoothCore{
			Range:  3,
			Factor: 5.0,
		}
		renderSmooth(g, size, nr, 1000000, radiusFactor, sc)
	}
}

func render2(g *image.Gray, size image.Point) {

	var seed int64

	r := newRand()

	seed = int64(r.Intn(10000))
	fmt.Println("seed:", seed)
	r.Seed(seed)
	//r.Seed(1286)
	// r := newRandSeed(seed)

	ps := make([]Point2f, 3)
	for i := range ps {
		ps[i] = Point2f{
			X: randRange(r, -1, 1),
			Y: randRange(r, -1, 1),
		}
	}

	radiusFactor := 0.85

	t := 0.5

	var nr Nexter = PtNext{
		t:  t,
		r:  newRand(),
		Ps: ps,
	}

	var smooth = true

	if !smooth {
		render(g, size, nr, 10000, radiusFactor)
	} else {
		sc := SmoothCore{
			Range:  3,
			Factor: 5.0,
		}
		renderSmooth(g, size, nr, 1000000, radiusFactor, sc)
	}
}

type PtNext struct {
	t  float64 // [0..1]
	r  *rand.Rand
	Ps []Point2f
}

var _ Nexter = PtNext{}

func (t PtNext) Next(p Point2f) Point2f {
	as := t.Ps
	k := t.r.Intn(len(as))
	return pointLerp(p, as[k], t.t)
}

type Params struct {
	Size image.Point

	n Nexter
}

type SmoothCore struct {
	Range  int
	Factor float64
}

func render3(g *image.Gray, size image.Point) {

	var seed int64

	// [-1*Pi ... +1*Pi]: 3, 7, 16, 22, 29, 39, 54, 57, 59, 67, 75, 80, 85, 86, 102, 109, 117, 2744, 9962, 9403, 4146, 9938, 8806
	// [-2*Pi ... +2*Pi]: 8809, 2980, 3040, 4683, 4612, 4465, 1286, 6806, 3831, 2353, 8816, 4324, 6782, 6461, 2102, 9389, 5757, 7955

	r := newRand()

	seed = int64(r.Intn(10000))
	fmt.Println("seed:", seed)
	r.Seed(seed)
	//r.Seed(8816)
	// r := newRandSeed(seed)

	d := 2.0
	var nr Nexter = Trig{
		A: randRange(r, -d, d) * math.Pi,
		B: randRange(r, -d, d) * math.Pi,
		C: randRange(r, -d, d) * math.Pi,
		D: randRange(r, -d, d) * math.Pi,
	}

	radiusFactor := 0.5

	var smooth = true

	if !smooth {
		render(g, size, nr, 1000000, radiusFactor)
	} else {
		sc := SmoothCore{
			Range:  3,
			Factor: 10.0,
		}
		renderSmooth(g, size, nr, 1000000, radiusFactor, sc)
	}
}

func randRange(r *rand.Rand, min, max float64) float64 {
	return min + (max-min)*r.Float64()
}

func applyFactor(x, factor float64) float64 {
	if x < 0 {
		return 0
	}
	if x < 1 {
		return 1 - math.Exp(factor*math.Log(1-x))
	}
	return 1
}

type Nexter interface {
	Next(p Point2f) Point2f
}

type Trig struct {
	A, B, C, D float64
}

var _ Nexter = Trig{}

func (t Trig) Next(p Point2f) Point2f {
	return Point2f{
		X: math.Sin(t.A*p.Y) + math.Cos(t.B*p.X),
		Y: math.Sin(t.C*p.X) + math.Cos(t.D*p.Y),
	}
}

func render(g *image.Gray, size image.Point, nr Nexter, n int, radiusFactor float64) {
	radius := radiusFactor * float64(minInt(size.X, size.Y)) / 2
	center := Point2f{
		X: float64(size.X),
		Y: float64(size.Y),
	}.DivScalar(2)

	r := newRand()

	p := Point2f{
		X: randRange(r, -1, 1),
		Y: randRange(r, -1, 1),
	}

	color1 := color.Gray{Y: 255}

	for i := 0; i < n; i++ {
		p = nr.Next(p)
		iX, iY := center.Add(p.MulScalar(radius)).IntXY()
		g.SetGray(iX, iY, color1)
	}
}

func renderSmooth(g *image.Gray, size image.Point, nr Nexter, n int, radiusFactor float64, sc SmoothCore) {

	m := sc.Range

	ssv := make([][]int, size.X*m)
	for i := range ssv {
		ssv[i] = make([]int, size.Y*m)
	}

	radius := float64(m) * (radiusFactor * float64(minInt(size.X, size.Y)) / 2)
	center := Point2f{
		X: float64(size.X),
		Y: float64(size.Y),
	}.MulScalar(float64(m) / 2)

	r := newRand()

	p := Point2f{
		X: randRange(r, -1, 1),
		Y: randRange(r, -1, 1),
	}

	var maxVal int
	for i := 0; i < n; i++ {
		p = nr.Next(p)
		iX, iY := center.Add(p.MulScalar(radius)).IntXY()
		if (0 <= iX) && (iX < size.X*m) {
			if (0 <= iY) && (iY < size.Y*m) {
				ssv[iX][iY]++
				maxVal = maxInt(maxVal, ssv[iX][iY])
			}
		}
	}

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
					sum += ssv[iX][iY]
				}
			}
			v := float64(sum) / float64(mm*maxVal)
			v = applyFactor(v, sc.Factor)
			gY := round(v * 255)
			g.SetGray(x, y, color.Gray{Y: uint8(gY)})
		}
	}
}
