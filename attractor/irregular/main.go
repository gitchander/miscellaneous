package main

import (
	"fmt"
	"log"

	"github.com/gitchander/miscellaneous/attractor"
	"github.com/gitchander/miscellaneous/attractor/utils"
	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
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

	var seed int64

	r := random.NewRandNow()

	seed = int64(r.Intn(10000))
	fmt.Println("seed:", seed)
	r.Seed(seed)
	//r.Seed(1286)

	ps := make([]Point2f, 5)
	for i := range ps {
		ps[i] = Point2f{
			X: random.RandInterval(r, -1, 1),
			Y: random.RandInterval(r, -1, 1),
		}
	}

	t := 0.5

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
