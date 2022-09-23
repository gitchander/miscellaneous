package utils

import (
	"fmt"
	"image/color"
)

func ParseColor(s string) (color.Color, error) {

	bs := []byte(s)
	if (len(bs) == 0) || (bs[0] != '#') {
		return nil, fmt.Errorf("invalid color: there is not prefix %c", '#')
	}
	bs = bs[1:]

	ds, err := parseHexDigits(bs)
	if err != nil {
		return nil, err
	}
	c := color.RGBA{A: 255}
	switch len(ds) {
	case 3:
		{
			c.R = nibblesToByte(ds[0], ds[0])
			c.G = nibblesToByte(ds[1], ds[1])
			c.B = nibblesToByte(ds[2], ds[2])
		}
	case 6:
		{
			c.R = nibblesToByte(ds[0], ds[1])
			c.G = nibblesToByte(ds[2], ds[3])
			c.B = nibblesToByte(ds[4], ds[5])
		}
	default:
		return nil, fmt.Errorf("invalid color %q", s)
	}
	return c, nil
}

func parseHexDigit(b byte) (digit int, ok bool) {
	if ('0' <= b) && (b <= '9') {
		return int(b - '0'), true
	}
	if ('a' <= b) && (b <= 'f') {
		return int(b-'a') + 10, true
	}
	if ('A' <= b) && (b <= 'F') {
		return int(b-'A') + 10, true
	}
	return 0, false
}

func parseHexDigits(bs []byte) ([]uint8, error) {
	ds := make([]uint8, len(bs))
	for i, b := range bs {
		d, ok := parseHexDigit(b)
		if !ok {
			return ds, fmt.Errorf("invalid hex digit %#U", b)
		}
		ds[i] = uint8(d)
	}
	return ds, nil
}

func ColorOver(dc, sc color.Color) color.Color {
	var (
		dc1 = color.RGBA64Model.Convert(dc).(color.RGBA64)
		sc1 = color.RGBA64Model.Convert(sc).(color.RGBA64)
	)
	return colorOverRGBA64(dc1, sc1)
}

func colorOverRGBA64(dc, sc color.RGBA64) color.RGBA64 {

	// m is the maximum color value returned by image.Color.RGBA.
	const m = 1<<16 - 1

	a := m - uint32(sc.A)

	return color.RGBA64{
		R: uint16((uint32(dc.R)*a)/m) + sc.R,
		G: uint16((uint32(dc.G)*a)/m) + sc.G,
		B: uint16((uint32(dc.B)*a)/m) + sc.B,
		A: uint16((uint32(dc.A)*a)/m) + sc.A,
	}
}

func LerpColor(c0, c1 color.Color, t float64) color.Color {

	var (
		v0 = color.RGBAModel.Convert(c0).(color.RGBA)
		v1 = color.RGBAModel.Convert(c1).(color.RGBA)
	)

	var (
		r = uint8(Round(Lerp(float64(v0.R), float64(v1.R), t)))
		g = uint8(Round(Lerp(float64(v0.G), float64(v1.G), t)))
		b = uint8(Round(Lerp(float64(v0.B), float64(v1.B), t)))
		a = uint8(Round(Lerp(float64(v0.A), float64(v1.A), t)))
	)

	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
