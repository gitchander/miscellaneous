package main

import (
	"fmt"
	"image"
	"math"
	"time"

	"github.com/gotk3/gotk3/cairo"
)

func makeImages() error {

	const timesPerSecond = 30
	dur := time.Second / time.Duration(timesPerSecond)
	deltaTime := dur.Seconds()

	// ps := [2]Pendulum{
	// 	0: Pendulum{
	// 		Mass:   50,
	// 		Length: 120,
	// 		Theta:  math.Pi * 0.3,
	// 	},
	// 	1: Pendulum{
	// 		Mass:   50,
	// 		Length: 120,
	// 		Theta:  -math.Pi * 0.5,
	// 	},
	// }
	// const t_max = 50.0

	dp := DoublePendulum{
		0: Pendulum{
			Mass:   50,
			Length: 120,
			Theta:  math.Pi / 2,
		},
		1: Pendulum{
			Mass:   50,
			Length: 120,
			Theta:  math.Pi * 0.5,
		},
	}
	const t_max = 106.0

	dts := []float64{0.1, 0.05, 0.01, 0.001, 0.0001}

	type node struct {
		dp   *DoublePendulum
		dt   float64
		prev Point2f
		cl   ColorRGBf
	}

	ns := make([]*node, len(dts))
	for i, dt := range dts {

		dp1 := dp
		_, _, x2, y2 := getDPCoords(&dp1)
		prev := Point2f{X: x2, Y: y2}

		t := float64(i) / float64(len(dts)-1)
		cl := clerp(RGBf(1, 0, 0), RGBf(0, 0, 0), t)

		//fmt.Println(t, cl)

		ns[i] = &node{
			dp:   &dp1,
			dt:   dt,
			prev: prev,
			cl:   cl,
		}
	}

	size := image.Point{X: 512, Y: 512}

	surface := cairo.CreateImageSurface(cairo.FORMAT_ARGB32, size.X, size.Y)
	context := cairo.Create(surface)

	context.SetSourceRGB(1, 1, 1)
	context.Paint()

	x0 := float64(size.X) / 2
	y0 := float64(size.Y) / 2
	context.Translate(x0, y0)

	// setColor(context, RGBf(0, 0, 0))
	// //context.SetSourceRGB(0, 0, 0)
	// context.MoveTo(0, 0)
	// context.LineTo(100, 100)
	// context.SetLineWidth(10)
	// context.Stroke()

	// setColor(context, RGBf(1, 0, 0))
	// //context.SetSourceRGB(1, 0, 0)
	// context.MoveTo(100, 0)
	// context.LineTo(0, 100)
	// context.SetLineWidth(10)
	// context.Stroke()

	if true {
		for _, n := range ns {

			fmt.Println(n.cl)

			context.MoveTo(n.prev.X, n.prev.Y)
			t := 0.0
			for t < t_max {

				nextStep(n.dp, deltaTime)

				_, _, x2, y2 := getDPCoords(n.dp)
				n.prev = Point2f{X: x2, Y: y2}

				context.LineTo(x2, y2)

				t += n.dt
			}

			context.SetLineWidth(1)
			setColor(context, n.cl)
			context.Stroke()
		}
	}

	return surface.WriteToPNG("test.png")
}
