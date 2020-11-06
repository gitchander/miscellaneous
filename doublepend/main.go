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

// https://www.youtube.com/watch?v=uWzPe_S-RVE
// https://www.myphysicslab.com/pendulum/double-pendulum-en.html
// http://bestofallpossibleurls.com/double-pendulum.html

func main() {
	//makeImages()
	Main()
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

	err := Run(c)
	checkError_(err)
}

func checkError_(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	Size image.Point
	DP   DoublePendulum
}

func Run(c Config) error {

	gtk.Init(nil)

	w, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return err
	}

	w.Connect("destroy", func() {
		gtk.MainQuit()
	})

	w.SetTitle("Double Pendulum")
	w.SetSizeRequest(c.Size.X, c.Size.Y)
	w.SetPosition(gtk.WIN_POS_CENTER)

	pal := makePalette1()

	var (
		//drawer = dummyDrawer{}
		drawer = NewDPDrawer(&(c.DP), pal)
	)

	da, err := makeDrawingArea(drawer)
	if err != nil {
		return err
	}

	process := func(d time.Duration) {
		drawer.Render(d)
		glib.IdleAdd(da.QueueDraw)
	}

	go playCore(process)

	w.Add(da)

	w.ShowAll()

	gtk.Main()

	return nil
}

type Drawer interface {
	Resize(width, height int)
	Draw(c *cairo.Context)
	Render(time.Duration)
}

type dummyDrawer struct{}

var _ Drawer = dummyDrawer{}

func (dummyDrawer) Resize(width, height int) {
	fmt.Printf("size: (%d,%d)\n", width, height)
}

func (dummyDrawer) Draw(*cairo.Context) {

}

func (dummyDrawer) Render(time.Duration) {

}

func makeDrawingArea(d Drawer) (*gtk.DrawingArea, error) {

	da, err := gtk.DrawingAreaNew()
	if err != nil {
		return nil, err
	}

	da.Connect("configure-event", func(da *gtk.DrawingArea, event *gdk.Event) {
		var (
			w = da.GetAllocatedWidth()
			h = da.GetAllocatedHeight()
		)
		d.Resize(w, h)
	})

	da.Connect("draw", func(da *gtk.DrawingArea, c *cairo.Context) {
		d.Draw(c)
	})

	return da, nil
}

func start(quit <-chan struct{}, ticks chan<- time.Duration, timesPerSecond int) {

	t0 := time.Now()
	dt := time.Second / time.Duration(timesPerSecond)
	t := t0

	var d time.Duration

	for {
		select {
		case <-quit:
			return
		case ticks <- d:
		}
		now := time.Now()
		d = now.Sub(t0)
		t = t.Add(dt)
		dSleep := t.Sub(now)
		if dSleep > 0 {
			time.Sleep(dSleep)
		}
	}
}

func playCore(process func(time.Duration)) {

	quit := make(chan struct{})
	ticks := make(chan time.Duration)

	go start(quit, ticks, 30)

	for d := range ticks {
		process(d)
	}
}
