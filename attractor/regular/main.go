package main

import (
	"log"

	"github.com/gitchander/miscellaneous/attractor"
	"github.com/gitchander/miscellaneous/attractor/utils"
	"github.com/gitchander/miscellaneous/attractor/utils/random"
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

	r := random.NewRandNow()

	ps := attractor.MakeRegularPoints(n)
	fr := attractor.NewPsFeeder(ps, t)
	p := attractor.RandPointInRadius(r, 1)
	nr := attractor.MakeNexter(fr, p)

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
