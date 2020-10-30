package main

import (
	"fmt"
	"image"
	"math"
	"sync"
	"time"

	"github.com/gotk3/gotk3/cairo"
)

type Palette struct {
	clBackground ColorRGBf
	clForeground ColorRGBf
	clPath       ColorRGBf // Trajectory

	clMassFill   ColorRGBf
	clMassStroke ColorRGBf
}

func makeDP(dp DoublePendulum) Drawer {

	dt := 0.2
	core := NewCore(dp, dt)

	pal := Palette{
		clBackground: RGBf(1, 1, 1),
		clForeground: RGBf(0, 0, 0),
		clPath:       RGBf(0.3, 0.3, 1),

		clMassFill:   RGBf(0, 0.5, 0),
		clMassStroke: RGBf(0, 0.2, 0),
	}

	return NewDPDrawer(core, pal)
}

type dpDrawer struct {
	guard sync.Mutex

	core *Core
	pal  Palette

	allocSize image.Point
	size      image.Point

	surfaceTrace     *cairo.Surface
	surfacePendulums *cairo.Surface

	contextTrace     *cairo.Context
	contextPendulums *cairo.Context

	prev OptPoint2f
}

var _ Drawer = &dpDrawer{}

func NewDPDrawer(core *Core, pal Palette) *dpDrawer {
	return &dpDrawer{
		core: core,
		pal:  pal,
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
		p.core.Next()
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

	x1, y1, x2, y2 := p.core.Coords()

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

type Pendulum struct {
	Mass     float64
	Length   float64
	Theta    float64
	Velocity float64
}

type DoublePendulum [2]Pendulum

type Core struct {
	dp DoublePendulum

	// m1 float64
	// m2 float64

	// r1 float64 // length 1
	// r2 float64 // length 2

	a1 float64 // theta 1
	a2 float64 // theta 2

	// velocities
	a1_v float64 // a1 velocity
	a2_v float64 // a2 velocity

	//thetaVelocity [2]float64

	//a1_a, a2_a float64 // accelerations

	dt float64 // deltaTime
}

func NewCore(dp DoublePendulum, dt float64) *Core {

	// timesPerSecond := 30
	// d := time.Second / time.Duration(timesPerSecond)

	return &Core{

		dp: dp,

		// m1: dp[0].Mass,
		// m2: dp[1].Mass,

		// r1: dp[0].Length,
		// r2: dp[1].Length,

		a1: dp[0].Theta,
		a2: dp[1].Theta,

		a1_v: dp[0].Velocity,
		a2_v: dp[1].Velocity,

		dt: dt,
	}
}

func (c *Core) Coords() (x1, y1, x2, y2 float64) {

	var (
		p1 = c.dp[0]
		p2 = c.dp[1]
	)

	var (
		r1 = p1.Length
		r2 = p2.Length

		a1 = c.a1
		a2 = c.a2
	)

	sin1, cos1 := math.Sincos(a1)

	x1 = r1 * sin1
	y1 = r1 * cos1

	sin2, cos2 := math.Sincos(a2)

	x2 = x1 + r2*sin2
	y2 = y1 + r2*cos2

	return
}

func (c *Core) Next() {

	var (
		p1 = c.dp[0]
		p2 = c.dp[1]
	)

	var (
		m1 = p1.Mass
		m2 = p2.Mass

		r1 = p1.Length
		r2 = p2.Length

		a1 = c.a1
		a2 = c.a2

		a1_v = c.a1_v
		a2_v = c.a2_v

		dt = c.dt
	)

	//---------------------------------------------------

	a1_a, a2_a := calc(m1, m2, r1, r2, a1, a2, a1_v, a2_v)

	a1_v += a1_a * dt
	a2_v += a2_a * dt

	a1 += a1_v * dt
	a2 += a2_v * dt

	// a1_v *= 0.99
	// a2_v *= 0.99

	//---------------------------------------------------

	c.a1 = a1
	c.a2 = a2

	c.a1_v = a1_v
	c.a2_v = a2_v
}

// type Thetas struct {
// 	theta    float64 // theta
// 	d_theta  float64 // theta'
// 	dd_theta float64 // theta"
// }

// https://github.com/myphysicslab/myphysicslab/blob/master/src/sims/pendulum/DoublePendulumSim.js

// accelerations: a1_a, a2_a
func calc(m1, m2 float64, r1, r2 float64, a1, a2 float64, a1_v, a2_v float64) (a1_a, a2_a float64) {

	const g = 9.81

	den := (2*m1 + m2 - m2*math.Cos(2*(a1-a2)))

	tmp1 := -g * (2*m1 + m2) * math.Sin(a1)
	tmp2 := -g * m2 * math.Sin(a1-2*a2)
	tmp3 := -2 * math.Sin(a1-a2) * m2
	tmp4 := a2_v*a2_v*r2 + a1_v*a1_v*r1*math.Cos(a1-a2)

	a1_a = (tmp1 + tmp2 + tmp3*tmp4) / (r1 * den)

	tmp1 = 2 * math.Sin(a1-a2)
	tmp2 = a1_v * a1_v * r1 * (m1 + m2)
	tmp3 = g * (m1 + m2) * math.Cos(a1)
	tmp4 = a2_v * a2_v * r2 * m2 * math.Cos(a1-a2)

	a2_a = (tmp1 * (tmp2 + tmp3 + tmp4)) / (r2 * den)

	return
}

func makeImages() {

	// ps := [2]Pendulum{
	// 	0: Pendulum{
	// 		Mass:   50,
	// 		Length: 120,
	// 		Theta:  math.Pi * 0.3,
	// 	},
	// 	1: Pendulum{
	// 		Mass:   50,
	// 		Length: 120,
	// 		Theta:  -math.Pi * 0.5,
	// 	},
	// }
	// const t_max = 50.0

	ps := [2]Pendulum{
		0: Pendulum{
			Mass:   50,
			Length: 120,
			Theta:  math.Pi / 2,
		},
		1: Pendulum{
			Mass:   50,
			Length: 120,
			Theta:  math.Pi * 0.5,
		},
	}
	const t_max = 106.0

	dts := []float64{0.1, 0.05, 0.01, 0.001, 0.0001}

	type node struct {
		core *Core
		dt   float64
		prev Point2f
		cl   ColorRGBf
	}

	ns := make([]*node, len(dts))
	for i, dt := range dts {

		core := NewCore(ps, dt)
		_, _, x2, y2 := core.Coords()
		prev := Point2f{X: x2, Y: y2}

		t := float64(i) / float64(len(dts)-1)
		cl := clerp(RGBf(1, 0, 0), RGBf(0, 0, 0), t)

		//fmt.Println(t, cl)

		ns[i] = &node{
			core: core,
			dt:   dt,
			prev: prev,
			cl:   cl,
		}
	}

	size := image.Point{X: 512, Y: 512}

	surface := cairo.CreateImageSurface(cairo.FORMAT_ARGB32, size.X, size.Y)
	context := cairo.Create(surface)

	context.SetSourceRGB(1, 1, 1)
	context.Paint()

	x0 := float64(size.X) / 2
	y0 := float64(size.Y) / 2
	context.Translate(x0, y0)

	// setColor(context, RGBf(0, 0, 0))
	// //context.SetSourceRGB(0, 0, 0)
	// context.MoveTo(0, 0)
	// context.LineTo(100, 100)
	// context.SetLineWidth(10)
	// context.Stroke()

	// setColor(context, RGBf(1, 0, 0))
	// //context.SetSourceRGB(1, 0, 0)
	// context.MoveTo(100, 0)
	// context.LineTo(0, 100)
	// context.SetLineWidth(10)
	// context.Stroke()

	if true {

		for _, n := range ns {

			fmt.Println(n.cl)

			context.MoveTo(n.prev.X, n.prev.Y)
			t := 0.0
			for t < t_max {

				n.core.Next()

				_, _, x2, y2 := n.core.Coords()
				n.prev = Point2f{X: x2, Y: y2}

				context.LineTo(x2, y2)

				t += n.dt
			}

			context.SetLineWidth(1)
			setColor(context, n.cl)
			context.Stroke()
		}

	}

	err := surface.WriteToPNG("test.png")
	checkError(err)
}
