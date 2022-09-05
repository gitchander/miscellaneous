package main

import (
	"log"

	"github.com/gitchander/miscellaneous/attractor"
	"github.com/gitchander/miscellaneous/attractor/utils"
	opt "github.com/gitchander/miscellaneous/attractor/utils/optional"
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
		Value:   3109666676878827058,
	}

	t := randTrigOptSeed(optSeed)
	nr := newAttrTrig(t, randFirstPoint())

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
