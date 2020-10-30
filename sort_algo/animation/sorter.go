package main

import (
	"image"
	"image/color"
	"image/gif"
	"sort"
)

const (
	colorBackground = iota
	colorValue
	colorIndex
	colorVertical
)

var sortPalette = []color.Color{
	colorBackground: rgb(0xFF, 0xFF, 0xFF), // White
	colorValue:      rgb(0x00, 0x00, 0x00), // Black
	colorIndex:      rgb(0xFF, 0x88, 0x88),
	colorVertical:   rgb(0xCC, 0xCC, 0xCC),
}

func rgb(r, g, b byte) color.RGBA {
	return color.RGBA{r, g, b, 0xFF}
}

type params struct {
	n        int
	delay    int // 100 = 1 second
	width    int
	f        func(sort.Interface)
	filename string
}

func (p params) run() {
	baseSort(p)
}

type Drawer struct {
	p         params
	vs        []int
	anim      *gif.GIF
	rect      image.Rectangle
	imageSize image.Point
	pal       []color.Color
}

func NewDrawer(p params, vs []int, anim *gif.GIF) *Drawer {

	minY, maxY := minMaxInt(vs)
	rect := image.Rectangle{
		Min: image.Point{X: 0, Y: minY},
		Max: image.Point{X: len(vs), Y: maxY + 1},
	}

	return &Drawer{
		p:         p,
		vs:        vs,
		anim:      anim,
		rect:      rect,
		imageSize: rect.Size().Mul(p.width),
		pal:       sortPalette,
	}
}

func (d *Drawer) Draw(i, j int) {
	r := image.Rectangle{Max: d.imageSize}
	ip := image.NewPaletted(r, d.pal)
	d.render(ip, i, j)
	d.anim.Delay = append(d.anim.Delay, d.p.delay)
	d.anim.Image = append(d.anim.Image, ip)
}

func (d *Drawer) render(ip *image.Paletted, i, j int) {

	yn := d.rect.Size().Y

	drawVertical := true

	y0 := yn - 1 + d.rect.Min.Y
	if d.p.width > 1 {
		for x, v := range d.vs {
			y := y0 - v
			if (x == i) || (x == j) {
				for k := 0; k < yn; k++ {
					drawUnit(ip, x, k, d.p.width, colorIndex)
				}
			} else if drawVertical {
				for k := y; k < yn; k++ {
					drawUnit(ip, x, k, d.p.width, colorVertical)
				}
			}
			drawUnit(ip, x, y, d.p.width, colorValue)
		}
	} else {
		for x, v := range d.vs {
			y := y0 - v
			if (x == i) || (x == j) {
				for k := 0; k < yn; k++ {
					ip.SetColorIndex(x, k, colorIndex)
				}
			} else if drawVertical {
				for k := y; k < yn; k++ {
					ip.SetColorIndex(x, k, colorVertical)
				}
			}
			ip.SetColorIndex(x, y, colorValue)
		}
	}
}

func drawUnit(ip *image.Paletted, posX, posY, width int, index uint8) {
	var (
		x0 = (posX + 0) * width
		x1 = (posX + 1) * width

		y0 = (posY + 0) * width
		y1 = (posY + 1) * width
	)
	for x := x0; x < x1; x++ {
		for y := y0; y < y1; y++ {
			ip.SetColorIndex(x, y, index)
		}
	}
}

func minMaxInt(a []int) (min, max int) {
	n := len(a)
	if n == 0 {
		return
	}
	min = a[0]
	max = a[0]
	for i := 1; i < n; i++ {
		if min > a[i] {
			min = a[i]
		}
		if max < a[i] {
			max = a[i]
		}
	}
	return
}

type IntSlice struct {
	vs []int
	d  *Drawer
}

func NewIntSlice(vs []int, d *Drawer) *IntSlice {
	return &IntSlice{vs: vs, d: d}
}

func (p *IntSlice) Len() int { return len(p.vs) }

func (p *IntSlice) Less(i, j int) bool {
	p.d.Draw(i, j)
	return p.vs[i] < p.vs[j]
}

func (p *IntSlice) Swap(i, j int) {
	vs := p.vs
	vs[i], vs[j] = vs[j], vs[i]
	p.d.Draw(i, j)
}
