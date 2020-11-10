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
		newSample(&(c.DP), GetPalette(0)),
	}

	err := Run(c.Size, samples, 1)
	checkError(err)
}

type Config struct {
	Size image.Point
	DP   DoublePendulum
}

func runRandom() {

	var number int

	flag.IntVar(&number, "number", 1, "number of double pendulums")

	flag.Parse()

	r := newRandNow()

	samples := randSamples(r, number)

	size := image.Point{X: 800, Y: 800}

	err := Run(size, samples, number)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Run(size image.Point, samples []*Sample, number int) error {

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
	var (
		deltaTime    = time.Second / time.Duration(timesPerSecond)
		deltaTimeSec = deltaTime.Seconds()
	)

	fmt.Println("deltaTime:", deltaTime)

	var (
		background = White
		// background = Black
	)

	engine := NewEngine(samples, background)

	eh := &engineEventHandler{
		number: number,
		engine: engine,
	}

	da, err := makeDrawingArea(eh)
	if err != nil {
		return err
	}

	process := func() bool {
		engine.CalcNextStep(deltaTimeSec)
		glib.IdleAdd(da.QueueDraw)
		return true
	}
	go runPeriodic(deltaTime, process)

	w.Add(da)

	da.SetCanFocus(true)

	w.ShowAll()

	gtk.Main()

	return nil
}

func makeDrawingArea(eh EventHandler) (*gtk.DrawingArea, error) {

	da, err := gtk.DrawingAreaNew()
	if err != nil {
		return nil, err
	}

	da.Connect("configure-event",
		func(da *gtk.DrawingArea, event *gdk.Event) {
			var (
				width  = da.GetAllocatedWidth()
				height = da.GetAllocatedHeight()
			)
			eh.Resize(width, height)
		})

	da.Connect("draw",
		func(da *gtk.DrawingArea, c *cairo.Context) {
			eh.Draw(c)
		})

	da.Connect("key-press-event",
		func(da *gtk.DrawingArea, event *gdk.Event) {

			eventKey := &gdk.EventKey{event}
			keyCode := eventKey.HardwareKeyCode()

			eh.EventKey(keyCode)
			//fmt.Println("key code:", keyCode)

			// t := eventKey.Type()
			// switch t {
			// case gdk.EVENT_KEY_PRESS:
			// 	fmt.Println("key press")
			// case gdk.EVENT_KEY_RELEASE:
			// 	fmt.Println("key release")
			// }
		})

	return da, nil
}
