package main

import (
	"flag"
	"log"

	attr "github.com/gitchander/miscellaneous/attractor"
	"github.com/gitchander/miscellaneous/attractor/utils"
	"github.com/gitchander/miscellaneous/attractor/utils/random"
)

// The Chaos Game

func main() {
	var c Config
	flag.IntVar(&(c.RegularAttractor.Points), "points", 3, "number of points")
	flag.Float64Var(&(c.RegularAttractor.Factor), "factor", 0.5, "distance factor")
	flag.StringVar(&(c.RegularAttractor.CPS), "cps", "random", "cornerpoint selector")
	flag.StringVar(&(c.RenderConfig), "render", attr.DefaultRenderConfigFilename, "render config filename")
	flag.Parse()

	checkError(run(c))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	RegularAttractor RegularAttractor
	RenderConfig     string
}

// https://beltoforion.de/en/recreational_mathematics/chaos_game.php
func FavorableFactor(n int) float64 {
	fn := float64(n)
	return fn / (fn + 3)
}

type RegularAttractor struct {
	Points int
	Factor float64 // Distance factor
	CPS    string  // CornerpointSelector
}

func run(c Config) error {

	var rc attr.RenderConfig
	err := utils.ReadConfigJSON(c.RenderConfig, &rc)
	if err != nil {
		err := attr.MakeDefaultRenderConfigFile(&rc)
		if err != nil {
			return err
		}
	}

	var (
		n = c.RegularAttractor.Points
		t = c.RegularAttractor.Factor
	)

	cps, err := attr.ParseCornerpointSelector(c.RegularAttractor.CPS)
	if err != nil {
		return err
	}

	r := random.NewRandNow()

	ps := attr.MakeRegularPoints(n)
	fr := attr.NewPsFeeder(ps, t, cps)
	p := attr.RandPointInRadius(r, 1)
	nr := attr.MakeNexter(fr, p)

	//--------------------------------------------------------------------------
	m, err := attr.Render(rc, nr)
	if err != nil {
		return err
	}

	return utils.SaveImage(rc.ImageFilename, m)
}
