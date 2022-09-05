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
