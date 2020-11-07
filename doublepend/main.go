package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"time"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	//makeImages()
	//Main()
	runRandom()
}

func Main() {

	var width, height int

	var (
		m1, m2 float64
		r1, r2 float64
		a1, a2 float64
	)

	flag.IntVar(&width, "width", 800, "width")
	flag.IntVar(&height, "height", 800, "height")

	flag.Float64Var(&m1, "m1", 10, "mass of pendulum 1")
	flag.Float64Var(&m2, "m2", 10, "mass of pendulum 2")

	flag.Float64Var(&r1, "r1", 150, "length of pendulum 1")
	flag.Float64Var(&r2, "r2", 150, "length of pendulum 2")

	flag.Float64Var(&a1, "a1", 90, "theta of pendulum 1")
	flag.Float64Var(&a2, "a2", 90, "theta of pendulum 2")

	flag.Parse()

	c := Config{
		Size: image.Point{X: width, Y: height},
		DP: DoublePendulum{
			Pendulum{
				Mass:   m1,
				Length: r1,
				Theta:  DegToRad(a1),
			},
			Pendulum{
				Mass:   m2,
				Length: r2,
				Theta:  DegToRad(a2),
			},
		},
	}

	samples := []*Sample{
		{
			dp:      &(c.DP),
			palette: GetPalette(0),
		},
	}

	err := Run(c.Size, samples)
	checkError(err)
}

type Config struct {
	Size image.Point
	DP   DoublePendulum
}

func runRandom() {

	r := newRandNow()

	n := 4

	dps := make([]*DoublePendulum, n)
	dp0 := randDoublePendulum(r)
	for i := 0; i < n; i++ {
		if i == 0 {
			dps[i] = dp0
		} else {
			clone := dp0.Clone()
			randChangeDoublePendulum(r, clone)
			dps[i] = clone
		}
	}

	samples := make([]*Sample, n)
	for i, dp := range dps {
		samples[i] = &Sample{
			dp:      dp,
			palette: GetPalette(i),
		}
	}

	size := image.Point{X: 800, Y: 800}

	err := Run(size, samples)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Run(size image.Point, samples []*Sample) error {

	gtk.Init(nil)

	w, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return err
	}

	w.Connect("destroy", func() {
		gtk.MainQuit()
	})

	w.SetTitle("Double Pendulum")
	w.SetSizeRequest(size.X, size.Y)
	w.SetPosition(gtk.WIN_POS_CENTER)

	const timesPerSecond = 30
	deltaTime := time.Second / time.Duration(timesPerSecond)

	fmt.Println("deltaTime:", deltaTime)

	background := RGBf(1, 1, 1)
	engine := NewEngine(samples, deltaTime, background)

	da, err := makeDrawingArea(engine)
	if err != nil {
		return err
	}

	process := func() bool {
		engine.CalcNextStep()
		glib.IdleAdd(da.QueueDraw)
		return true
	}
	go runPeriodic(deltaTime, process)

	w.Add(da)

	w.ShowAll()

	gtk.Main()

	return nil
}

func makeDrawingArea(d Drawer) (*gtk.DrawingArea, error) {

	da, err := gtk.DrawingAreaNew()
	if err != nil {
		return nil, err
	}

	da.Connect("configure-event",
		func(da *gtk.DrawingArea, event *gdk.Event) {
			var (
				w = da.GetAllocatedWidth()
				h = da.GetAllocatedHeight()
			)
			d.Resize(w, h)
		})

	da.Connect("draw",
		func(da *gtk.DrawingArea, c *cairo.Context) {
			d.Draw(c)
		})

	return da, nil
}

func runPeriodic(period time.Duration, f func() bool) {
	t := time.Now()
	for {
		if !f() {
			return
		}
		// calc sleep
		t = t.Add(period)
		d := t.Sub(time.Now())
		if d > 0 {
			time.Sleep(d)
		}
	}
}
