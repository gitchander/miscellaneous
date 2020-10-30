package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"image/color"
	"math/rand"
)

// func randColor(r *rand.Rand) color.Color {
// 	return color.RGBA{
// 		R: byte(32 * r.Intn(8)),
// 		G: byte(32 * r.Intn(8)),
// 		B: byte(32 * r.Intn(8)),
// 		A: 255,
// 	}
// }

func randColor(r *rand.Rand) color.Color {
	return color.RGBA{
		R: byte(16 * r.Intn(16)),
		G: byte(16 * r.Intn(16)),
		B: byte(16 * r.Intn(16)),
		A: 255,
	}
}

// func firstByteIs(bs []byte, b byte) bool {
// 	return (len(bs) > 0) && (bs[0] == b)
// }

func parseHexColor(s string) (color.Color, error) {
	if (len(s) == 0) || (s[0] != '#') {
		return nil, errors.New("invalid hex color format")
	}
	bs := []byte(s[1:])
	switch len(bs) {
	case 3: // rgb
		fallthrough
	case 4: // rgba
		return nibblesToColor(bs)
	case 6: // rrggbb
		fallthrough
	case 8: // rrggbbaa
		return hexToColor(bs)
	default:
		return nil, fmt.Errorf("invalid color %q", s)
	}
}

func nibblesToColor(bs []byte) (color.Color, error) {

	ns, err := hexToNibbles(bs)
	if err != nil {
		return nil, err
	}

	if len(ns) == 3 {
		ns = append(ns, 0xf)
	}

	var (
		r = ns[0]
		g = ns[1]
		b = ns[2]
		a = ns[3]
	)

	r |= r << 4
	g |= g << 4
	b |= b << 4
	a |= a << 4

	return color.RGBA{R: r, G: g, B: b, A: a}, nil
}

func hexToColor(bs []byte) (color.Color, error) {

	hs := make([]byte, hex.DecodedLen(len(bs)))
	_, err := hex.Decode(hs, bs)
	if err != nil {
		return nil, err
	}

	if len(hs) == 3 {
		hs = append(hs, 0xff)
	}

	var (
		r = hs[0]
		g = hs[1]
		b = hs[2]
		a = hs[3]
	)

	return color.RGBA{R: r, G: g, B: b, A: a}, nil
}

func hexToNibbles(bs []byte) ([]byte, error) {
	ns := make([]byte, len(bs))
	for i, b := range bs {
		n, ok := hexToNibble(b)
		if !ok {
			return ns, fmt.Errorf("invalid hex byte %c", b)
		}
		ns[i] = n
	}
	return ns, nil
}

func hexToNibble(b byte) (byte, bool) {
	if ('0' <= b) && (b <= '9') {
		return (b - '0'), true
	}
	if ('a' <= b) && (b <= 'f') {
		return (b - 'a' + 10), true
	}
	if ('A' <= b) && (b <= 'F') {
		return (b - 'A' + 10), true
	}
	return 0, false
}

func nibbleToHex(n byte) (byte, bool) {
	if (0 <= n) && (n < 10) {
		return (n + '0'), true
	}
	if (10 <= n) && (n < 16) {
		return (n + 'a' - 10), true
	}
	return 0, false
}
