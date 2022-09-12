package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func NewImageBySize(imageSize image.Point) draw.Image {
	r := image.Rectangle{Max: imageSize}
	return image.NewRGBA(r)
}

func FillImage(m draw.Image, c color.Color) {
	b := m.Bounds()
	draw.Draw(m, b, image.NewUniform(c), image.ZP, draw.Src)
}

func saveImagePNG(filename string, m image.Image) error {
	var b bytes.Buffer
	err := png.Encode(&b, m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b.Bytes(), 0666)
}

func SaveImage(filename string, m image.Image) error {
	ext := filepath.Ext(filename)
	var b bytes.Buffer
	switch ext {
	case ".png":
		err := png.Encode(&b, m)
		if err != nil {
			return err
		}
	case ".jpeg":
		err := jpeg.Encode(&b, m, nil)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid image ext %q", ext)
	}
	return ioutil.WriteFile(filename, b.Bytes(), 0666)
}

func ParseSize(s string) (image.Point, error) {
	var p image.Point
	vs := strings.Split(s, "x")
	if len(vs) != 2 {
		return p, fmt.Errorf("parse size: invalid size value %s", s)
	}
	x, err := strconv.Atoi(vs[0])
	if err != nil {
		return p, fmt.Errorf("parse size %c value: %s", 'x', err)
	}
	y, err := strconv.Atoi(vs[1])
	if err != nil {
		return p, fmt.Errorf("parse size %c value: %s", 'y', err)
	}
	p = image.Pt(x, y)
	return p, nil
}
