package main

import (
	"image"
	"sync"

	"github.com/gotk3/gotk3/cairo"
)

type Engine struct {
	guard sync.Mutex

	samples    []*Sample
	background ColorRGBf
	pause      bool

	allocSize image.Point
	size      image.Point

	surfaceTrace     *cairo.Surface
	surfacePendulums *cairo.Surface

	contextTrace     *cairo.Context
	contextPendulums *cairo.Context
}

func NewEngine(samples []*Sample, background ColorRGBf) *Engine {
	return &Engine{
		samples:    samples,
		background: background,
	}
}

func (p *Engine) SetSamples(samples []*Sample) {
	p.guard.Lock()
	{
		p.samples = samples
	}
	p.guard.Unlock()
}

func (p *Engine) _Resize(width, height int) {
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

func (p *Engine) _Draw(context *cairo.Context) {
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
func (p *Engine) CalcNextStep(deltaTime float64) {
	p.guard.Lock()
	{
		if !p.pause {
			p.calcNextStep(deltaTime)
		}
	}
	p.guard.Unlock()
}

func (p *Engine) calcNextStep(deltaTime float64) {

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
		sample.calcNextStep(deltaTime)
		sample.renderTail(p.contextPendulums, x0, y0)
		sample.renderSample(p.contextPendulums, x0, y0)
	}

	// draw anchor circle
	{
		c.Arc(x0, y0, 3, 0, Tau)
		setColor(c, Gray75)
		//setColor(c, White)
		c.FillPreserve()
		setColor(c, Gray25)
		c.Stroke()
	}
}

func (p *Engine) Pause() {
	p.guard.Lock()
	{
		p.pause = !(p.pause)
	}
	p.guard.Unlock()
}
