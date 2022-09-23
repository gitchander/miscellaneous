package attractor

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/gitchander/miscellaneous/attractor/utils"
	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
)

func Render(rc RenderConfig, nr Nexter) (image.Image, error) {

	pas, err := configToParams(rc)
	if err != nil {
		return nil, err
	}

	m := utils.NewImageBySize(pas.imageSize)

	if pas.fillBG {
		utils.FillImage(m, pas.colorBG)
	}

	if not(pas.smooth) {
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

	var correctFunc = pas.correctFunc

	fg := color.NRGBA64Model.Convert(pas.colorFG).(color.NRGBA64)

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
			t := correctFunc(v)

			var (
				bg  = g.At(x, y)
				fgt = colorLerpAlpha(fg, t)
			)

			c := utils.ColorOver(bg, fgt)

			g.Set(x, y, c)
		}
	}
}

func floorCeil(x float64) (floor float64, ceil float64) {
	floor = math.Floor(x)
	//ceil = math.Ceil(x)
	ceil = floor + 1
	return
}

// t: [0..1]
func colorLerpAlpha(c color.NRGBA64, t float64) color.NRGBA64 {
	c.A = uint16(math.Round(t * float64(c.A)))
	return c
}
