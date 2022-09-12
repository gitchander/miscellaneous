package main

import (
	"flag"
	"log"

	attr "github.com/gitchander/miscellaneous/attractor"
	"github.com/gitchander/miscellaneous/attractor/utils"
	"github.com/gitchander/miscellaneous/attractor/utils/random"
)

func main() {
	var c Config
	flag.IntVar(&(c.RegularAttractor.Points), "points", 3, "number of points")
	flag.Float64Var(&(c.RegularAttractor.Koef), "koef", 0.5, "attractor koef")
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

type RegularAttractor struct {
	Points int
	Koef   float64
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
		t = c.RegularAttractor.Koef
	)

	r := random.NewRandNow()

	ps := attr.MakeRegularPoints(n)
	fr := attr.NewPsFeeder(ps, t)
	p := attr.RandPointInRadius(r, 1)
	nr := attr.MakeNexter(fr, p)

	//--------------------------------------------------------------------------
	m, err := attr.Render(rc, nr)
	if err != nil {
		return err
	}

	return utils.SaveImage(rc.ImageFilename, m)
}
