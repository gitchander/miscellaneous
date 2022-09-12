package main

import (
	"flag"
	"log"

	attr "github.com/gitchander/miscellaneous/attractor"
	"github.com/gitchander/miscellaneous/attractor/utils"
	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
	"github.com/gitchander/miscellaneous/attractor/utils/random"
)

func main() {
	var c Config
	flag.StringVar(&(c.Seed), "seed", "", "seed value")
	flag.IntVar(&(c.IrregularAttractor.Points), "points", 3, "number of points")
	flag.Float64Var(&(c.IrregularAttractor.Factor), "factor", 0.5, "distance factor")
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
	Seed               string
	IrregularAttractor IrregularAttractor
	RenderConfig       string
}

type IrregularAttractor struct {
	Points int
	Factor float64 // Distance factor
}

func run(c Config) error {

	seed, err := random.ParseSeedOrMake(c.Seed)
	if err != nil {
		return err
	}

	var rc attr.RenderConfig
	err = utils.ReadConfigJSON(c.RenderConfig, &rc)
	if err != nil {
		err := attr.MakeDefaultRenderConfigFile(&rc)
		if err != nil {
			return err
		}
	}

	var (
		n = c.IrregularAttractor.Points
		t = c.IrregularAttractor.Factor
	)

	r := random.NewRandSeed(seed)

	ps := make([]Point2f, n)
	for i := range ps {
		ps[i] = Point2f{
			X: random.RandInterval(r, -1, 1),
			Y: random.RandInterval(r, -1, 1),
		}
	}

	cps := attr.CornerpointRandom

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
