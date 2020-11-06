package main

import (
	"image"
	"math"
	"sync"
	"time"

	"github.com/gotk3/gotk3/cairo"
)

// timesPerSecond := 30
// d := time.Second / time.Duration(timesPerSecond)

const dpDeltaTime = 0.2

type dpDrawer struct {
	guard sync.Mutex

	dp  *DoublePendulum
	pal Palette

	allocSize image.Point
	size      image.Point

	surfaceTrace     *cairo.Surface
	surfacePendulums *cairo.Surface

	contextTrace     *cairo.Context
	contextPendulums *cairo.Context

	prev OptPoint2f
}

var _ Drawer = &dpDrawer{}

func NewDPDrawer(dp *DoublePendulum, pal Palette) *dpDrawer {
	return &dpDrawer{
		dp:  dp,
		pal: pal,
	}
}

func (p *dpDrawer) Resize(width, height int) {
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

		setColor(p.contextTrace, p.pal.clBackground)
		p.contextTrace.Paint()
	}
	p.guard.Unlock()
}

func (p *dpDrawer) Draw(context *cairo.Context) {
	p.guard.Lock()
	{
		if p.surfacePendulums != nil {
			context.SetSourceSurface(p.surfacePendulums, 0, 0)
			context.Paint()
		}
	}
	p.guard.Unlock()
}

func (p *dpDrawer) Render(d time.Duration) {
	p.guard.Lock()
	{
		nextStep(p.dp, dpDeltaTime)
		p.renderNext()
	}
	p.guard.Unlock()
}

func (p *dpDrawer) renderNext() {

	if p.surfacePendulums == nil {
		return
	}

	x0 := float64(p.size.X) / 2
	y0 := float64(p.size.Y) / 2

	x1, y1, x2, y2 := getDPCoords(p.dp)

	//-------------------------------------------------------
	if true {

		c := p.contextTrace

		if p.prev.Present {
			c.Save()
			c.Translate(x0, y0)
			c.MoveTo(p.prev.Value.X, p.prev.Value.Y)
			c.LineTo(x2, y2)
			c.SetLineWidth(1)
			setColor(c, p.pal.clPath)
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

	c.SetSourceSurface(p.surfaceTrace, 0, 0)
	c.Paint()

	c.Translate(x0, y0)

	c.MoveTo(0, 0)
	c.LineTo(x1, y1)
	c.LineTo(x2, y2)
	setColor(c, p.pal.clForeground)
	c.Stroke()

	c.Arc(x1, y1, 10, 0, 2*math.Pi)
	setColor(c, p.pal.clMassFill)
	c.FillPreserve()
	setColor(c, p.pal.clMassStroke)
	c.Stroke()

	c.Arc(x2, y2, 10, 0, 2*math.Pi)
	setColor(c, p.pal.clMassFill)
	c.FillPreserve()
	setColor(c, p.pal.clMassStroke)
	c.Stroke()

	c.Restore()
}
