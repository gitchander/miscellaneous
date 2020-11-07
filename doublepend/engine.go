package main

import (
	"image"
	"sync"
	"time"

	"github.com/gotk3/gotk3/cairo"
)

type Sample struct {
	dp      *DoublePendulum
	palette Palette
}

type Engine struct {
	guard sync.Mutex

	samples      []*Sample
	deltaTimeSec float64
	background   ColorRGBf

	allocSize image.Point
	size      image.Point

	surfaceTrace     *cairo.Surface
	surfacePendulums *cairo.Surface

	contextTrace     *cairo.Context
	contextPendulums *cairo.Context

	prev OptPoint2f
}

var _ Drawer = &Engine{}

func NewEngine(samples []*Sample, deltaTime time.Duration, background ColorRGBf) *Engine {
	return &Engine{
		samples:      samples,
		deltaTimeSec: deltaTime.Seconds(),
		background:   background,
	}
}

func (p *Engine) Resize(width, height int) {
	p.guard.Lock()
	{
		size := image.Point{
			X: width,
			Y: height,
		}

		if (size.X > p.allocSize.X) || (size.Y > p.allocSize.Y) {

			allocSize := image.Point{
				X: ceilPowerOfTwo(size.X),
				Y: ceilPowerOfTwo(size.Y),
			}

			p.surfaceTrace = cairo.CreateImageSurface(cairo.FORMAT_ARGB32,
				allocSize.X, allocSize.Y)

			p.surfacePendulums = cairo.CreateImageSurface(cairo.FORMAT_ARGB32,
				allocSize.X, allocSize.Y)

			p.contextTrace = cairo.Create(p.surfaceTrace)
			p.contextPendulums = cairo.Create(p.surfacePendulums)

			p.allocSize = allocSize
		}

		p.size = size

		setColor(p.contextTrace, p.background)
		p.contextTrace.Paint()
	}
	p.guard.Unlock()
}

func (p *Engine) Draw(context *cairo.Context) {
	p.guard.Lock()
	{
		if p.surfacePendulums != nil {
			context.SetSourceSurface(p.surfacePendulums, 0, 0)
			context.Paint()
		}
	}
	p.guard.Unlock()
}

// UpdateTime
func (p *Engine) CalcNextStep() {
	p.guard.Lock()
	{
		p.renderNext()
	}
	p.guard.Unlock()
}

func (p *Engine) renderNext() {

	if p.surfacePendulums == nil {
		return
	}

	c := p.contextPendulums
	c.SetSourceSurface(p.surfaceTrace, 0, 0)
	c.Paint()

	var (
		x0 = float64(p.size.X) / 2
		y0 = float64(p.size.Y) / 2
	)

	for _, sample := range p.samples {
		nextStep(sample.dp, p.deltaTimeSec)
		p.renderSample(x0, y0, sample)
	}

	// draw anchor circle
	{
		c.Arc(x0, y0, 3, 0, Tau)
		setColor(c, Gray50)
		c.FillPreserve()
		setColor(c, Black)
		c.Stroke()
	}
}

func (p *Engine) renderSample(x0, y0 float64, sample *Sample) {

	x1, y1, x2, y2 := getDPCoords(sample.dp)

	//-------------------------------------------------------
	if false {

		c := p.contextTrace

		if p.prev.Present {
			c.Save()
			c.Translate(x0, y0)
			c.MoveTo(p.prev.Value.X, p.prev.Value.Y)
			c.LineTo(x2, y2)
			c.SetLineWidth(1)
			setColor(c, sample.palette.Path)
			c.Stroke()
			c.Restore()
		}

		p.prev.Value.X = x2
		p.prev.Value.Y = y2
		p.prev.Present = true
	}

	//-------------------------------------------------------
	c := p.contextPendulums

	c.Save()

	//c.SetSourceSurface(p.surfaceTrace, 0, 0)

	c.Translate(x0, y0)

	radius := 7.0

	c.MoveTo(0, 0)
	c.LineTo(x1, y1)
	c.LineTo(x2, y2)
	setColor(c, sample.palette.Foreground)
	c.Stroke()

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
