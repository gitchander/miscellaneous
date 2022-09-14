package attractor

import (
	"image"
	"image/color"

	"github.com/gitchander/miscellaneous/attractor/utils"
)

type RenderConfig struct {
	ImageFilename string  `json:"image_filename"`
	ImageSize     string  `json:"image_size"`
	TotalPoints   int     `json:"total_points"`
	RadiusFactor  float64 `json:"radius_factor"`
	FillBG        bool    `json:"fill_bg"`
	ColorBG       string  `json:"color_bg"`
	ColorFG       string  `json:"color_fg"`
	Smooth        bool    `json:"smooth"`
	ToneFactor    float64 `json:"tone_factor"`
}

type params struct {
	imageSize    image.Point
	totalPoints  int
	radiusFactor float64

	fillBG  bool
	colorBG color.Color
	colorFG color.Color

	smooth      bool
	correctFunc CorrectFunc
}

func configToParams(rc RenderConfig) (*params, error) {

	imageSize, err := utils.ParseSize(rc.ImageSize)
	if err != nil {
		return nil, err
	}

	colorBG, err := utils.ParseColor(rc.ColorBG)
	if err != nil {
		return nil, err
	}
	colorFG, err := utils.ParseColor(rc.ColorFG)
	if err != nil {
		return nil, err
	}

	p := &params{
		imageSize:    imageSize,
		totalPoints:  rc.TotalPoints,
		radiusFactor: rc.RadiusFactor,
		fillBG:       rc.FillBG,
		colorBG:      colorBG,
		colorFG:      colorFG,
		smooth:       rc.Smooth,
		correctFunc:  ToneCorrectionFunc(rc.ToneFactor),
	}

	return p, nil
}
