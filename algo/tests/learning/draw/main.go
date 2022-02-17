package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"

	"github.com/fogleman/gg"

	lrn "github.com/gitchander/miscellaneous/algo/tests/learning"
	"github.com/gitchander/miscellaneous/algo/tests/utils"
)

const tau = 2 * math.Pi

func main() {

	size := lrn.Pt(512, 512)

	r := utils.NewRandNow()

	//----------------------------------------------------------
	// rt := lrn.Rect{
	// 	Min: lrn.Pt(size.X/4, size.Y/4),
	// 	Max: lrn.Pt(size.X*3/4, size.Y*3/4),
	// }
	// ps := randPointsRect(r, rt, 100)
	//----------------------------------------------------------

	c := Circle{
		Center: lrn.Pt(256, 256),
		Radius: 200,
	}
	ps := randPointsCircle(r, c, 100)
	//----------------------------------------------------------

	// ps = []Point{
	// 	Pt(125, 125),
	// 	Pt(100, 100),
	// 	Pt(110, 110),
	// 	Pt(100, 150),
	// 	Pt(100, 200),
	// 	Pt(150, 175),
	// 	Pt(200, 150),
	// 	Pt(250, 125),
	// 	Pt(300, 100),
	// 	Pt(150, 100),
	// }
	// fmt.Println(ps)

	// cs := convexPolygon(ps)
	// fmt.Println(cs)

	err := drawPoints(size, ps)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Circle struct {
	Center lrn.Point
	Radius int
}

func randPointsRect(r *rand.Rand, rt lrn.Rect, n int) []lrn.Point {
	ps := make([]lrn.Point, n)
	for i := range ps {
		ps[i] = lrn.Point{
			X: utils.RandIntMinMax(r, rt.Min.X, rt.Max.X),
			Y: utils.RandIntMinMax(r, rt.Min.Y, rt.Max.Y),
		}
	}
	return ps
}

func randPointsCircle(r *rand.Rand, c Circle, n int) []lrn.Point {

	var min, max lrn.Point

	min.X = c.Center.X - c.Radius
	max.X = c.Center.X + c.Radius
	if min.X > max.X {
		min.X, max.X = max.X, min.X
	}

	min.Y = c.Center.Y - c.Radius
	max.Y = c.Center.Y + c.Radius
	if min.Y > max.Y {
		min.Y, max.Y = max.Y, min.Y
	}

	rr := c.Radius * c.Radius

	ps := make([]lrn.Point, n)
	for i := range ps {

		for {
			p := lrn.Point{
				X: utils.RandIntMinMax(r, min.X, max.X),
				Y: utils.RandIntMinMax(r, min.Y, max.Y),
			}

			dx := p.X - c.Center.X
			dy := p.Y - c.Center.Y

			if ((dx * dx) + (dy * dy)) <= rr {
				ps[i] = p
				break
			}
		}
	}
	return ps
}

func rasterLine(r lrn.Rect, m1, m2 lrn.Point) string {
	var b strings.Builder
	for y := r.Max.Y - 1; y >= r.Min.Y; y-- {
		for x := r.Min.X; x < r.Max.X; x++ {
			var char byte
			p := lrn.Pt(x, y)
			d := lrn.LineTest(m1, m2, p)
			switch {
			case d < 0:
				char = '-'
			case d > 0:
				char = '+'
			case d == 0:
				char = '0'
			}
			b.WriteByte(char)
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func drawPoints(size lrn.Point, ps []lrn.Point) error {

	var (
		drawPoints        = true
		drawPointsIndexes = false

		drawConvexPolygon       = true
		drawConvexPolygonPoints = true

		drawArrows = true
	)

	cp := lrn.ConvexPolygon(ps)

	dc := gg.NewContext(size.X, size.Y)

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	if drawConvexPolygon {
		dc.SetRGB(0.6, 0.0, 0.0)
		dc.SetLineWidth(1)

		// draw convex polygon.
		dc.NewSubPath()
		for _, p := range cp {
			var (
				x = float64(p.X)
				y = float64(p.Y)
			)
			dc.LineTo(x, y)
		}
		dc.ClosePath()
		dc.Stroke()
	}
	if drawArrows {
		//dc.SetRGB(0.5, 0.0, 0.0)
		dc.SetRGB(0.0, 0.4, 0.0)
		i := len(cp) - 1
		for j := 0; j < len(cp); j++ {
			var (
				x1 = float64(cp[i].X)
				y1 = float64(cp[i].Y)

				x2 = float64(cp[j].X)
				y2 = float64(cp[j].Y)
			)
			drawArrow(dc, x1, y1, x2, y2)
			i = j
		}
		dc.Fill()
		//dc.Stroke()
	}

	if drawPoints {

		//dc.SetRGB(0.2, 0.3, 0.8)
		dc.SetLineWidth(1)

		pointRadius := 2.0

		for i, p := range ps {

			var (
				x = float64(p.X)
				y = float64(p.Y)
			)
			dc.DrawCircle(x, y, pointRadius)

			dc.SetColor(color.White)
			dc.FillPreserve()

			dc.SetRGB(0.2, 0.3, 0.8)
			dc.Stroke()

			if drawPointsIndexes {
				dc.SetRGB(0, 0, 0)
				dc.DrawString(strconv.Itoa(i), x+5, y)
			}
		}
	}

	// draw convex points.
	if drawConvexPolygonPoints {
		//dc.SetRGB(0.0, 0.5, 0.0)
		dc.SetLineWidth(1)
		for _, p := range cp {
			var (
				x = float64(p.X)
				y = float64(p.Y)
			)
			dc.DrawCircle(x, y, 5)

			// dc.SetRGB(1.0, 1.0, 1.0)
			// dc.FillPreserve()

			dc.SetRGB(0.0, 0.5, 0.0)
			dc.Stroke()
		}
	}

	return dc.SavePNG("image.png")
}

func drawCircle(dc *gg.Context, x, y float64, radius float64) {
	dc.DrawArc(x, y, radius, 0, tau)
}

// (x1, y1) -> (x2, y2)
func drawArrow(dc *gg.Context, x1, y1 float64, x2, y2 float64) {

	dx := (x2 - x1)
	dy := (y2 - y1)

	var (
		angle    = math.Atan2(dy, dx)
		distance = math.Hypot(dy, dx)
	)

	dangle := tau / 8
	radius := 8.0
	if radius > distance/4 {
		radius = distance / 4
	}

	w := 20.0
	if w > distance/2 {
		w = distance / 2
	}

	// center of circle
	cx := x2 - (w+radius)*math.Cos(angle)
	cy := y2 - (w+radius)*math.Sin(angle)

	a1 := angle - dangle
	ax1 := cx + radius*math.Cos(a1)
	ay1 := cy + radius*math.Sin(a1)

	a2 := angle + dangle
	ax2 := cx + radius*math.Cos(a2)
	ay2 := cy + radius*math.Sin(a2)

	dc.NewSubPath()
	dc.MoveTo(ax1, ay1)
	dc.LineTo(x2, y2)
	dc.LineTo(ax2, ay2)
	dc.DrawArc(cx, cy, radius, a2, a1)
	dc.ClosePath()
}
