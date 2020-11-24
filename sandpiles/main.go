package main

import (
	"flag"
	"image"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

// https://en.wikipedia.org/wiki/Abelian_sandpile_model
// https://www.youtube.com/watch?v=1MtEUErz7Gg

var (
	randSource = rand.NewSource(time.Now().UnixNano())
	rnd        = rand.New(randSource)
)

func main() {

	var (
		fWidth  = flag.Int("width", 256, "")
		fHeight = flag.Int("height", 256, "")
		fState  = flag.Int("state", 4, "initial state for cells")
		fScale  = flag.Int("scale", 2, "scale factor")
	)

	flag.Parse()

	var size = image.Point{
		X: *fWidth,
		Y: *fHeight,
	}

	//population := int((size.X * size.Y) / 2)
	scale := *fScale

	w := NewWorld(size)
	w.fillValue(int8(*fState))

	go func() {
		const stepsPerSecond = 50
		d := time.Second / stepsPerSecond
		sl := newSleeper(d)
		for {
			w.Progress()
			sl.Sleep()
		}
	}()

	//w.FillTest(20)
	//w.RandomSeed(population)

	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowSize(scale*size.X, scale*size.Y)
	ebiten.SetWindowTitle("sandpiles")
	//ebiten.SetWindowResizable(true)
	//ebiten.SetWindowDecorated(true)
	//ebiten.SetCursorMode(ebiten.CursorModeHidden)

	err := ebiten.RunGame(w)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
