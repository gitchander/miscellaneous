package main

import (
	"log"
	"math"

	"github.com/gitchander/miscellaneous/attractor"
	"github.com/gitchander/miscellaneous/attractor/utils"
	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
)

func main() {
	checkError(run())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {

	var (
		// n = 2
		// t = 0.7

		// n = 3
		// t = 0.5

		n = 6
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
	)

	ps := makePoints(n)
	nr := attractor.NewPtNext(ps, t)
	nr.Randomize(Pt2f(-1, -1), Pt2f(1, 1))

	//--------------------------------------------------------------------------
	rc := attractor.RenderConfig{
		ImageSize:    "700x700",
		TotalPoints:  10_000_000,
		RadiusFactor: 0.9,
		FillBG:       true,
		ColorBG:      "#fff",
		ColorFG:      "#000",
		Smooth: attractor.SmoothConfig{
			Present: true,
			Range:   1,
			Factor:  15,
		},
	}

	m, err := attractor.Render(rc, nr)
	if err != nil {
		return err
	}

	return utils.SaveImagePNG("attractor.png", m)
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
