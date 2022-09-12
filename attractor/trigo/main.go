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
	flag.StringVar(&(c.Seed), "seed", "", "random seed")
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
	Seed         string
	RenderConfig string
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

	r := random.NewRandSeed(seed)

	fr := randTrigo(r)
	p := attr.RandPointInRadius(r, 1)

	nr := attr.MakeNexter(fr, p)

	//--------------------------------------------------------------------------
	m, err := attr.Render(rc, nr)
	if err != nil {
		return err
	}

	return utils.SaveImage(rc.ImageFilename, m)
}
