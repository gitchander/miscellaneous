package base91

import (
	"fmt"
)

// http://base91.sourceforge.net/#a1
// https://stackoverflow.com/questions/46978133/base91-how-is-it-calculated
// const tableBase91Str = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\""

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

func makeDecodeTable() []byte {
	dt := make([]byte, 256)
	for i := range dt {
		dt[i] = 0xFF
	}
	for i, b := range alphabet {
		dt[b] = byte(i)
	}
	return dt
}

func EncodedLenMax(x int) int { return ceilDiv(x*16, 13) } // + 23%

func DecodedLenMax(x int) int { return x }

func Encode(dst, src []byte) int {

	j := 0

	var (
		av uint32 // bits accumulator
		an int    // bits length
	)

	for _, b := range src {

		av |= uint32(b) << an
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

		dst[j] = encodeTable[v0]
		j++

		dst[j] = encodeTable[v1]
		j++
	}

	if an > 0 {
		v := int(av)
		v1, v0 := quoRem(v, 91)
		dst[j] = encodeTable[v0]
		j++
		if (an > 7) || (v > 90) {
			dst[j] = encodeTable[v1]
			j++
		}
	}

	return j
}

func Decode(dst, src []byte) (int, error) {

	j := 0

	var (
		av uint32 // bits accumulator
		an int    // bits length
	)

	v := -1

	for _, b := range src {

		dv := decodeTable[b]
		if dv == 0xFF {
			return j, fmt.Errorf("base91: invalid byte: %#U", rune(b))
		}

		if v == -1 {
			v = int(dv)
			continue
		}

		v += int(dv) * 91

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
			dst[j] = byte(av)
			j++

			av >>= bitsPerByte
			an -= bitsPerByte
		}
	}

	if v != -1 {
		av |= uint32(v) << an
		dst[j] = byte(av)
		j++
	}

	return j, nil
}

func EncodeToString(src []byte) string {
	dst := make([]byte, EncodedLenMax(len(src)))
	n := Encode(dst, src)
	return string(dst[:n])
}

func DecodeString(s string) ([]byte, error) {
	src := []byte(s)
	n, err := Decode(src, src)
	return src[:n], err
}
