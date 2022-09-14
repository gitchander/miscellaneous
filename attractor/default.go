package attractor

import (
	"github.com/gitchander/miscellaneous/attractor/utils"
)

var DefaultRenderConfig = RenderConfig{
	ImageFilename: "attractor.png",
	ImageSize:     "512x512",
	TotalPoints:   10_000_000,
	RadiusFactor:  0.9,
	FillBG:        true,
	ColorBG:       "#fff",
	ColorFG:       "#000",
	Smooth:        true,
	ToneFactor:    15,
}

const DefaultRenderConfigFilename = "render_config.json"

func MakeDefaultRenderConfigFile(p *RenderConfig) error {
	rc := DefaultRenderConfig
	err := utils.WriteConfigJSON(DefaultRenderConfigFilename, rc)
	if err != nil {
		return err
	}
	*p = rc
	return nil
}
