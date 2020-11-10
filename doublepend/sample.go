package main

import (
	"container/ring"

	"github.com/gotk3/gotk3/cairo"
)

const lengthScale = 100.0

type Sample struct {
	dp      *DoublePendulum
	palette Palette
	tail    *ring.Ring
}

func newSample(dp *DoublePendulum, palette Palette) *Sample {
	return &Sample{
		dp:      dp,
		palette: palette,
		tail:    ring.New(40),
	}
}

func (sample *Sample) renderTail(c *cairo.Context, x0, y0 float64) {
	//sample.renderTailSolid(c, x0, y0)
	sample.renderTailAlpha(c, x0, y0)
}

func (sample *Sample) renderTailSolid(c *cairo.Context, x0, y0 float64) {

	c.Save()

	c.Translate(x0, y0)
	c.SetLineWidth(1)

	fg := sample.palette.MassFill

	first := true
	sample.tail.Do(
		func(v interface{}) {
			if p, ok := v.(Point2f); ok {
				if first {
					c.MoveTo(p.X, p.Y)
					first = false
				} else {
					c.LineTo(p.X, p.Y)
				}
			}
		})

	setColor(c, fg)
	c.Stroke()

	c.Restore()
}

func (sample *Sample) renderTailAlpha(c *cairo.Context, x0, y0 float64) {

	c.Save()

	c.Translate(x0, y0)
	c.SetLineWidth(1)

	fg := sample.palette.MassFill

	i := 0
	n := sample.tail.Len()
	var prev Point2f
	first := true

	sample.tail.Do(
		func(v interface{}) {
			if p, ok := v.(Point2f); ok {
				if first {
					prev = p
					first = false
				} else {
					c.MoveTo(prev.X, prev.Y)
					c.LineTo(p.X, p.Y)
					prev = p

					alpha := float64(i) / float64(n-1)
					setColorAlpha(c, fg, alpha)
					c.Stroke()
				}
			}
			i++
		})

	c.Restore()
}

func (sample *Sample) renderSample(c *cairo.Context, x0, y0 float64) {

	x1, y1, x2, y2 := getDPCoords(sample.dp, lengthScale)

	c.Save()

	c.Translate(x0, y0)
	c.SetLineWidth(5)
	c.SetLineJoin(cairo.LINE_JOIN_ROUND)

	radius := 5.0

	c.MoveTo(0, 0)
	c.LineTo(x1, y1)
	c.LineTo(x2, y2)
	//setColor(c, sample.palette.Foreground)
	setColorAlpha(c, sample.palette.Foreground, 0.5)
	c.Stroke()

	c.SetLineWidth(1)

	c.Arc(x1, y1, radius, 0, Tau)
	setColor(c, sample.palette.MassFill)
	c.FillPreserve()
	setColor(c, sample.palette.MassStroke)
	c.Stroke()

	c.Arc(x2, y2, radius, 0, Tau)
	setColor(c, sample.palette.MassFill)
	c.FillPreserve()
	setColor(c, sample.palette.MassStroke)
	c.Stroke()

	c.Restore()
}

func (sample *Sample) calcNextStep(deltaTime float64) {
	nextStep(sample.dp, deltaTime)

	_, _, x2, y2 := getDPCoords(sample.dp, lengthScale)

	sample.tail.Value = Pt2f(x2, y2)
	sample.tail = sample.tail.Next()
}
