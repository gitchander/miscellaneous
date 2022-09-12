package attractor

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/gitchander/miscellaneous/attractor/utils"
	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
)

type RenderConfig struct {
	ImageFilename string       `json:"image_filename"`
	ImageSize     string       `json:"image_size"`
	TotalPoints   int          `json:"total_points"`
	RadiusFactor  float64      `json:"radius_factor"`
	FillBG        bool         `json:"fill_bg"`
	ColorBG       string       `json:"color_bg"`
	ColorFG       string       `json:"color_fg"`
	Smooth        SmoothConfig `json:"smooth"`
}

type SmoothConfig struct {
	Present bool    `json:"present"`
	Range   int     `json:"range"`
	Factor  float64 `json:"factor"`
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
		//renderSmooth1(pas, nr, m)
		renderSmooth2(pas, nr, m)
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

func renderSmooth1(pas *params, nr Nexter, g draw.Image) {

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

func renderSmooth2(pas *params, nr Nexter, g draw.Image) {

	var smoothFactor = pas.smooth.Factor

	size := pas.imageSize

	//r:= image.Rectangle{Max:pas.imageSize }

	ssv := make([][]float64, size.X)
	for i := range ssv {
		ssv[i] = make([]float64, size.Y)
	}

	radius := pas.radiusFactor * float64(minInt(size.X, size.Y)) / 2
	center := Point2f{
		X: float64(size.X),
		Y: float64(size.Y),
	}.DivScalar(2)

	cellAdd := func(x, y int, v float64) {

		if not((0 <= x) && (x < size.X)) {
			return
		}
		if not((0 <= y) && (y < size.Y)) {
			return
		}

		ssv[x][y] += v
	}

	for i := 0; i < pas.totalPoints; i++ {

		pn := nr.Next()
		p := center.Add(pn.MulScalar(radius))

		var (
			x0, x1 = floorCeil(p.X)
			y0, y1 = floorCeil(p.Y)
		)

		var (
			t00 = (1 - (p.X - x0)) * (1 - (p.Y - y0))
			t01 = (1 - (p.X - x0)) * (1 - (y1 - p.Y))
			t10 = (1 - (x1 - p.X)) * (1 - (p.Y - y0))
			t11 = (1 - (x1 - p.X)) * (1 - (y1 - p.Y))
		)

		cellAdd(int(x0), int(y0), t00)
		cellAdd(int(x0), int(y1), t01)
		cellAdd(int(x1), int(y0), t10)
		cellAdd(int(x1), int(y1), t11)
	}

	var maxVal float64
	for _, sv := range ssv {
		for _, v := range sv {
			maxVal = maxFloat64(maxVal, v)
		}
	}

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {

			v := ssv[x][y] / maxVal
			t := applyFactor(v, smoothFactor)

			c0 := g.At(x, y)
			c := utils.LerpColor(c0, pas.colorFG, t)
			g.Set(x, y, c)
		}
	}
}

func not(b bool) bool {
	return !b
}

func floorCeil(x float64) (floor float64, ceil float64) {
	floor = math.Floor(x)
	//ceil = math.Ceil(x)
	ceil = floor + 1
	return
}
