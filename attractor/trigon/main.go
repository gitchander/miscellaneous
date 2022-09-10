package main

import (
	"fmt"
	"log"

	"github.com/gitchander/miscellaneous/attractor"
	"github.com/gitchander/miscellaneous/attractor/utils"
	opt "github.com/gitchander/miscellaneous/attractor/utils/optional"
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

	optSeed := opt.OptInt64{
		Present: false,
		Value:   6579706586152203246,
		// 6579706586152203246
	}

	var seed int64
	if optSeed.Present {
		seed = optSeed.Value
	} else {
		seed = random.NewRandNow().Int63()
	}
	fmt.Println("seed", seed)

	r := random.NewRandSeed(seed)

	p := attractor.RandPointInRadius(r, 2)

	rps := attractor.RegularPoints(2, 2, -0.25)
	fr2 := attractor.NewPsFeeder(rps, 0.5)

	mf := attractor.MultiFeeder(randTrig(r), fr2)
	nr := attractor.MakeNexter(mf, p)

	//--------------------------------------------------------------------------
	rc := attractor.RenderConfig{
		ImageSize:    "700x700",
		TotalPoints:  10_000_000,
		RadiusFactor: 0.9 * 0.5, // interval: [-2..+2]
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
