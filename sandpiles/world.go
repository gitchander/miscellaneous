package main

import (
	"image"
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten"
)

var palette = []color.Color{
	color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xFF},
	color.RGBA{R: 0xFF, G: 0xFF, B: 0x00, A: 0xFF},
	color.RGBA{R: 0x33, G: 0xFF, B: 0xFF, A: 0xFF},
	color.RGBA{R: 0x99, G: 0x4C, B: 0x00, A: 0xFF},
	color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
}

// World represents the game state
type World struct {
	guard sync.Mutex
	index int
	areas [2][][]int8
	size  image.Point

	noiseImage *image.RGBA

	pause bool
	keys  map[ebiten.Key]*KeyState
}

// NewWorld creates a new world
func NewWorld(size image.Point) *World {
	return &World{
		areas: [2][][]int8{
			0: makeArea(size),
			1: makeArea(size),
		},
		size:       size,
		noiseImage: image.NewRGBA(image.Rect(0, 0, size.X, size.Y)),
		keys: map[ebiten.Key]*KeyState{
			ebiten.KeySpace: NewKeyState(ebiten.KeySpace),
			ebiten.Key0:     NewKeyState(ebiten.Key0),
			ebiten.Key1:     NewKeyState(ebiten.Key1),
			ebiten.Key2:     NewKeyState(ebiten.Key2),
			ebiten.Key3:     NewKeyState(ebiten.Key3),
			ebiten.Key4:     NewKeyState(ebiten.Key4),
			ebiten.Key5:     NewKeyState(ebiten.Key5),
			ebiten.Key6:     NewKeyState(ebiten.Key6),
			ebiten.Key7:     NewKeyState(ebiten.Key7),
			ebiten.KeyR:     NewKeyState(ebiten.KeyR),
		},
	}
}

func makeArea(size image.Point) [][]int8 {
	area := make([][]int8, size.X)
	for i := range area {
		area[i] = make([]int8, size.Y)
	}
	return area
}

// RandomSeed inits world with a random state
func (w *World) randomSeed(limit int) {

	var (
		yN = w.size.Y
		xN = w.size.X
	)
	curr := w.areas[w.index]
	for i := 0; i < limit; i++ {
		var (
			x = rnd.Intn(xN)
			y = rnd.Intn(yN)
		)
		curr[x][y] = int8(rnd.Intn(6))
	}
}

func (w *World) fillValue(val int8) {

	var (
		yN = w.size.Y
		xN = w.size.X
	)
	curr := w.areas[w.index]
	//next := w.areas[w.nextIndex()]

	for x := 0; x < xN; x++ {
		for y := 0; y < yN; y++ {
			curr[x][y] = val
			//next[x][y] = val
		}
	}
}

func (w *World) FillTest(val int8) {

	w.guard.Lock()
	defer w.guard.Unlock()

	var (
		yN = w.size.Y
		xN = w.size.X
	)
	curr := w.areas[w.index]

	xc := xN / 2
	yc := yN / 2
	d := xN / 4

	x0 := xc - d
	x1 := xc + d

	y0 := yc - d
	y1 := yc + d

	for y := y0; y < y1; y++ {
		for x := x0; x < x1; x++ {
			curr[x][y] = val
		}
	}
}

func (w *World) nextIndex() int {
	var next int
	if w.index == 0 {
		next = 1
	}
	return next
}

// Progress game state by one tick
func (w *World) Progress() {

	w.guard.Lock()
	defer w.guard.Unlock()

	if w.pause {
		return
	}

	var (
		yN = w.size.Y
		xN = w.size.X
	)

	curr := w.areas[w.index]

	var nextIndex = w.nextIndex()
	next := w.areas[nextIndex]

	for y := 0; y < yN; y++ {
		for x := 0; x < xN; x++ {
			next[x][y] = curr[x][y]
		}
	}

	for y := 0; y < yN; y++ {
		for x := 0; x < xN; x++ {

			if curr[x][y] >= 4 {
				next[x][y] -= 4

				if y > 0 {
					next[x][y-1]++
				}
				if y < (yN - 1) {
					next[x][y+1]++
				}
				if x > 0 {
					next[x-1][y]++
				}
				if x < (xN - 1) {
					next[x+1][y]++
				}
			}
		}
	}

	w.index = nextIndex
}

// DrawImage paints current game state
func (w *World) drawImage(img *image.RGBA) {

	curr := w.areas[w.index]

	var (
		yN = w.size.Y
		xN = w.size.X
	)

	for y := 0; y < yN; y++ {
		for x := 0; x < xN; x++ {
			pos := 4*y*xN + 4*x

			val := int(curr[x][y])

			val = cropMax(val, len(palette)-1)
			setColor(img.Pix[pos:], palette[val])
		}
	}
}

func setColor(Pix []uint8, c color.Color) {
	x := color.RGBAModel.Convert(c).(color.RGBA)
	Pix[0] = x.R
	Pix[1] = x.G
	Pix[2] = x.B
	Pix[3] = x.A
}

func cropMax(x, max int) int {
	if x > max {
		x = max
	}
	return x
}

func (w *World) Update() error {
	w.Progress()
	return nil
}

func (w *World) Draw(screen *ebiten.Image) {

	w.guard.Lock()
	defer w.guard.Unlock()

	w.drawImage(w.noiseImage)
	screen.ReplacePixels(w.noiseImage.Pix)

	if key, ok := w.keys[ebiten.KeySpace]; ok {
		key.Update()
		if key.KeyDown() {
			w.pause = !w.pause
		}
	}

	if key, ok := w.keys[ebiten.Key0]; ok {
		key.Update()
		if key.KeyDown() {
			w.fillValue(0)
		}
	}

	if key, ok := w.keys[ebiten.Key1]; ok {
		key.Update()
		if key.KeyDown() {
			w.fillValue(1)
		}
	}

	if key, ok := w.keys[ebiten.Key2]; ok {
		key.Update()
		if key.KeyDown() {
			w.fillValue(2)
		}
	}

	if key, ok := w.keys[ebiten.Key3]; ok {
		key.Update()
		if key.KeyDown() {
			w.fillValue(3)
		}
	}

	if key, ok := w.keys[ebiten.Key4]; ok {
		key.Update()
		if key.KeyDown() {
			w.fillValue(4)
		}
	}

	if key, ok := w.keys[ebiten.Key5]; ok {
		key.Update()
		if key.KeyDown() {
			w.fillValue(5)
		}
	}

	if key, ok := w.keys[ebiten.Key6]; ok {
		key.Update()
		if key.KeyDown() {
			w.fillValue(6)
		}
	}

	if key, ok := w.keys[ebiten.Key7]; ok {
		key.Update()
		if key.KeyDown() {
			w.fillValue(7)
		}
	}

	if key, ok := w.keys[ebiten.KeyR]; ok {
		key.Update()
		if key.KeyDown() {
			w.randomSeed(1000)
		}
	}
}

func (w *World) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	//fmt.Println(outsideWidth, outsideHeight)
	return w.size.X, w.size.Y
}
