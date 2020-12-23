package main

import (
	"fmt"
)

// http://base91.sourceforge.net/#a1
// https://stackoverflow.com/questions/46978133/base91-how-is-it-calculated

// const tableBase91Str = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\""

//------------------------------------------------------------------------------
const bitsPerByte = 8

const (
	mask13bit = 0x1fff
	mask14bit = 0x3fff
)

var alphabet = []byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
	'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
	'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '!', '#', '$',
	'%', '&', '(', ')', '*', '+', ',', '.', '/', ':', ';', '<', '=',
	'>', '?', '@', '[', ']', '^', '_', '`', '{', '|', '}', '~', '"',
}

var (
	encodeTable = alphabet
	decodeTable = makeDecodeTable()
)

func makeDecodeTable() []int {
	dt := make([]int, 256)
	for i := range dt {
		dt[i] = -1
	}
	for i, b := range alphabet {
		dt[b] = i
	}
	return dt
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}

func Encode(d []byte) []byte {

	var res []byte

	var (
		av uint32 // bits accumulator
		an int    // bits length
	)

	for _, di := range d {

		av |= uint32(di) << an
		an += bitsPerByte

		if an < 14 {
			continue
		}

		v := int(av & mask13bit) // 13 bits

		if v > 88 {
			av >>= 13
			an -= 13
		} else {
			v = int(av & mask14bit) // 14 bits
			av >>= 14
			an -= 14
		}

		v1, v0 := quoRem(v, 91)

		res = append(res, encodeTable[v0], encodeTable[v1])
	}

	if an > 0 {
		v := int(av)
		v1, v0 := quoRem(v, 91)
		res = append(res, encodeTable[v0])
		if (an > 7) || (v > 90) {
			res = append(res, encodeTable[v1])
		}
	}

	return res
}

func Decode(d []byte) []byte {

	var res []byte

	var (
		av uint32 // bits accumulator
		an int    // bits length
	)

	v := -1

	for _, di := range d {

		dv := decodeTable[di]
		if dv == -1 { // skip invalid data
			continue
		}

		if v == -1 {
			v = dv
			continue
		}

		v += dv * 91

		var vn int
		if (v & mask13bit) > 88 {
			vn = 13
		} else {
			vn = 14
		}

		av |= uint32(v) << an
		an += vn

		v = -1

		for an >= bitsPerByte {
			res = append(res, byte(av))
			av >>= bitsPerByte
			an -= bitsPerByte
		}
	}

	if v != -1 {
		av |= uint32(v) << an
		res = append(res, byte(av))
	}

	return res
}

//------------------------------------------------------------------------------
func main() {
	//text := repeatByte(0x00, 5)
	//text := repeatByte(0xff, 3)
	//text := []byte("test")
	text := []byte("Hello World!")

	e := Encode(text)
	d := Decode(e)

	fmt.Printf("enc: %s\n", e) // "fPNKd"
	fmt.Printf("dec: %s\n", d)

	fmt.Printf("text: [% x]\n", text)
	fmt.Printf("dec:  [% x]\n", d)
}

func repeatByte(b byte, n int) []byte {
	bs := make([]byte, n)
	for i := range bs {
		bs[i] = b
	}
	return bs
}
