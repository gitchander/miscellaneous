package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
)

func main() {

	var (
		//		par   = logoSmile
		//		width = 16

		par   = logoChander
		width = 32
	)

	ps, _ := parseLines(par.Lines, '0')
	pal := paletteGithub
	const colorSet = 1

	rd := image.Rectangle{Max: par.Size.Mul(width)}
	pd := image.NewPaletted(rd, pal)
	for _, p := range ps {
		p = p.Add(par.Pos)
		if width <= 1 {
			pd.SetColorIndex(p.X, p.Y, colorSet)
		} else {
			drawUnit(pd, p, width, colorSet)
		}
	}
	err := saveImagePNG(pd, "logo.png")
	checkError(err)
}

func saveImagePNG(im image.Image, filename string) error {
	var buf bytes.Buffer
	err := png.Encode(&buf, im)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, buf.Bytes(), 0644)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var one = image.Point{X: 1, Y: 1}

func drawUnit(pd *image.Paletted, p image.Point, width int, colorIndex uint8) {
	var (
		p0 = p.Mul(width)
		p1 = p.Add(one).Mul(width)
	)
	for x := p0.X; x < p1.X; x++ {
		for y := p0.Y; y < p1.Y; y++ {
			pd.SetColorIndex(x, y, colorIndex)
		}
	}
}

var (
	paletteWB = color.Palette{
		0: color.White,
		1: color.RGBA{R: 32, G: 32, B: 32, A: 255},
	}

	paletteGithub = color.Palette{
		0: color.RGBA{R: 0xF0, G: 0xF0, B: 0xF0, A: 0xFF}, // f0f0f0
		1: color.RGBA{R: 0x6E, G: 0xC8, B: 0xD8, A: 0xFF}, // 6ec8d8
	}
)

type Params struct {
	Lines []string
	Size  image.Point
	Pos   image.Point
}

var (
	logoSmile = Params{
		Lines: []string{
			"----0000----",
			"--00----00--",
			"-0--------0-",
			"-0--------0-",
			"0--00--00--0",
			"0----------0",
			"0----------0",
			"0--0----0--0",
			"-0--0000--0-",
			"-0--------0-",
			"--00----00--",
			"----0000----",
		},
		Size: image.Point{X: 16, Y: 16},
		Pos:  image.Point{X: 2, Y: 2},
	}

	logoChander = Params{
		Lines: []string{
			"-------00--00-",
			"------0-0-0---",
			"------0-0-0---",
			"------0-0-0---",
			"-------00--00-",
			"------0-0-00--",
			"-----0--00-0--",
			"----0---0--0--",
			"---0-------0--",
			"--0--------0--",
			"0000-----00000",
		},
		Size: image.Point{X: 16, Y: 16},
		Pos:  image.Point{X: 1, Y: 3},
	}
)

func parseLines(lines []string, target byte) ([]image.Point, image.Rectangle) {
	var ps []image.Point
	for y, line := range lines {
		bs := []byte(line)
		for x, b := range bs {
			if b == target {
				ps = append(ps, image.Pt(x, y))
			}
		}
	}
	return ps, PointsBounds(ps)
}

func PointsBounds(ps []image.Point) image.Rectangle {
	var min, max image.Point
	for i, p := range ps {
		if i == 0 {
			min = p

			max.X = p.X + 1
			max.Y = p.Y + 1
		} else {
			min.X = minInt(min.X, p.X)
			min.Y = minInt(min.Y, p.Y)

			max.X = maxInt(max.X, p.X+1)
			max.Y = maxInt(max.Y, p.Y+1)
		}
	}
	return image.Rectangle{Min: min, Max: max}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
