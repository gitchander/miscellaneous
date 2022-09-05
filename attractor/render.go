package attractor

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/gitchander/miscellaneous/attractor/utils"
	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
)

type RenderConfig struct {
	ImageSize    string       `json:"image_size"`
	TotalPoints  int          `json:"total_points"`
	RadiusFactor float64      `json:"radius_factor"`
	FillBG       bool         `json:"fill_bg"`
	ColorBG      string       `json:"color_bg"`
	ColorFG      string       `json:"color_fg"`
	Smooth       SmoothConfig `json:"smooth"`
}

type SmoothConfig struct {
	Present bool    `json:"present"`
	Range   int     `json:"range"`
	Factor  float64 `json:"factor"`
}

type Nexter interface {
	Next() Point2f
}

type params struct {
	imageSize    image.Point
	totalPoints  int
	radiusFactor float64

	fillBG  bool
	colorBG color.Color
	colorFG color.Color

	smooth SmoothConfig
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
	}

	return p, nil
}

func Render(rc RenderConfig, nr Nexter) (image.Image, error) {

	pas, err := configToParams(rc)
	if err != nil {
		return nil, err
	}

	m := utils.NewImageBySize(pas.imageSize)

	if pas.fillBG {
		utils.FillImage(m, pas.colorBG)
	}

	if !(pas.smooth.Present) {
		renderBrute(pas, nr, m)
	} else {
		renderSmooth(pas, nr, m)
	}

	return m, nil
}

func renderBrute(pas *params, nr Nexter, g draw.Image) {

	radius := pas.radiusFactor * float64(minInt(pas.imageSize.X, pas.imageSize.Y)) / 2
	center := Point2f{
		X: float64(pas.imageSize.X),
		Y: float64(pas.imageSize.Y),
	}.DivScalar(2)

	for i := 0; i < pas.totalPoints; i++ {
		pn := nr.Next()
		p := center.Add(pn.MulScalar(radius)).Point()
		g.Set(p.X, p.Y, pas.colorFG)
	}
}

func renderSmooth(pas *params, nr Nexter, g draw.Image) {

	var (
		m            = pas.smooth.Range
		smoothFactor = pas.smooth.Factor
	)

	ssv := make([][]int, (pas.imageSize.X * m))
	for i := range ssv {
		ssv[i] = make([]int, (pas.imageSize.Y * m))
	}

	radius := float64(m) * (pas.radiusFactor * float64(minInt(pas.imageSize.X, pas.imageSize.Y)) / 2)
	center := Point2f{
		X: float64(pas.imageSize.X),
		Y: float64(pas.imageSize.Y),
	}.MulScalar(float64(m) / 2)

	var maxVal int
	for i := 0; i < pas.totalPoints; i++ {
		pn := nr.Next()
		p := center.Add(pn.MulScalar(radius)).Point()
		if (0 <= p.X) && (p.X < (pas.imageSize.X * m)) {
			if (0 <= p.Y) && (p.Y < (pas.imageSize.Y * m)) {
				ssv[p.X][p.Y]++
				maxVal = maxInt(maxVal, ssv[p.X][p.Y])
			}
		}
	}

	mm := m * m
	for x := 0; x < pas.imageSize.X; x++ {
		mX := m * x
		for y := 0; y < pas.imageSize.Y; y++ {
			mY := m * y
			var sum int
			for dX := 0; dX < m; dX++ {
				iX := mX + dX
				for dY := 0; dY < m; dY++ {
					iY := mY + dY
					sum += ssv[iX][iY]
				}
			}
			v := float64(sum) / float64(mm*maxVal)
			t := applyFactor(v, smoothFactor)

			c0 := g.At(x, y)
			c := utils.LerpColor(c0, pas.colorFG, t)
			g.Set(x, y, c)
		}
	}
}
