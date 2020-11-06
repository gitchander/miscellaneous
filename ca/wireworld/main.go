package main

import (
	"flag"
	"log"
	"time"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"

	"github.com/mattn/go-runewidth"
)

var (
	cellSize = Point{
		X: 2,
		Y: 1,
	}

	// cellSize = Point{
	// 	X: 4,
	// 	Y: 2,
	// }

	// cellSize = Point{
	// 	X: 6,
	// 	Y: 3,
	// }
)

func main() {
	//printMooreNeighborhood(3)
	//printVonNeumannNeighborhood(3)

	Main()
}

func Main() {

	var filename string

	flag.StringVar(&filename, "sample", "", "sample filename")

	flag.Parse()

	w := NewWorld(Pt(50, 40))

	err := ParseFile(w, Pt(2, 2), filename)
	checkError(err)

	encoding.Register()

	s, err := tcell.NewScreen()
	checkError(err)
	defer s.Fini()

	err = s.Init()
	checkError(err)

	bg := styleByHexColor(0x222222) // gray

	s.SetStyle(bg)

	commands := make(chan int)

	go func() {
		for {
			event := s.PollEvent()
			switch t := event.(type) {
			case *tcell.EventResize:
				{
					s.Sync()
				}
			case *tcell.EventKey:
				{
					key := t.Key()
					if key == tcell.KeyEscape {
						commands <- commandQuit
						return
					}

					switch r := t.Rune(); r {
					case ' ':
						commands <- commandPause
					case 'n':
						commands <- commandNext
					case 'f':
						commands <- commandFaster
					case 's':
						commands <- commandSlower
					}
				}
			}
		}
	}()

	stepDur := &stepDuration{
		d:    150 * time.Millisecond,
		min:  10 * time.Millisecond,
		max:  1000 * time.Millisecond,
		gain: 25 * time.Millisecond,
	}

	ticker := time.NewTicker(stepDur.Duration())
	defer ticker.Stop()

	var pause bool

start:
	for {
		select {
		case <-(ticker.C):
			if !pause {
				w.Next()
				renderWorld(s, w)
			}
		case command := <-commands:
			{
				switch command {
				case commandQuit:
					break start
				case commandPause:
					pause = !pause
				case commandNext:
					{
						w.Next()
						renderWorld(s, w)
						pause = true
					}
				case commandFaster:
					{
						stepDur.Faster()
						ticker.Reset(stepDur.Duration())
					}
				case commandSlower:
					{
						stepDur.Slower()
						ticker.Reset(stepDur.Duration())
					}
				}
			}
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	commandUndefined = iota
	commandQuit
	commandPause
	commandNext
	commandFaster
	commandSlower
)

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func getScreenSize(s tcell.Screen) Point {
	width, height := s.Size()
	return Point{
		X: width,
		Y: height,
	}
}

func fillRect(s tcell.Screen, x0, y0, x1, y1 int, r rune, style tcell.Style) {
	for y := y0; y < y1; y++ {
		for x := x0; x < x1; x++ {
			s.SetContent(x, y, r, nil, style)
		}
	}
}

func renderWorld(s tcell.Screen, w *World) {

	s.Clear()

	f := func(p Point, state int8) bool {

		style := stylesPalette[state]

		var (
			x0 = (p.X * cellSize.X)
			y0 = (p.Y * cellSize.Y)

			x1 = x0 + cellSize.X
			y1 = y0 + cellSize.Y
		)

		fillRect(s, x0, y0, x1, y1, ' ', style)

		return true
	}
	w.Range(f)

	s.Show()
}
