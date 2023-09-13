package main

import (
	"encoding/binary"
	"fmt"
	"math"

	"mvstick/utils/varint"
)

func main() {
	// textIntsEncode()
	testEncodeInt64()
	testGoVarint()
}

func textIntsEncode() {
	for x := -128; x < 128; x++ {
		a := int8(x)
		b := encodeInt8(a)
		fmt.Printf("%d -> %d\n", a, b)
	}
	fmt.Println()
	for x := 0; x < 256; x++ {
		var (
			a = uint8(x)
			b = decodeInt8(a)
			c = encodeInt8(b)
		)
		fmt.Printf("%4d -> %4d -> %4d\n", a, b, c)
	}
}

func encodeInt8(a int8) uint8 {
	if a < 0 {
		b := uint8(-(a + 1))
		return (b << 1) + 1
	} else {
		b := uint8(a)
		return b << 1
	}
}

func decodeInt8(a uint8) int8 {
	var (
		// b = int8(a / 2)
		b = int8(a >> 1)
	)
	if (a % 2) != 0 {
		b = -b - 1
	}
	return b
}

func decodeInt8_v2(a uint8) int8 {
	var b int8
	if (a % 2) == 0 {
		b = int8(a / 2)
	} else {
		b = -int8(a/2) - 1
	}
	return b
}

func testEncodeInt64() {
	samples := []int64{
		0,
		-1,
		1,
		-2,
		2,
		-100,
		-0xffff,
	}
	data := make([]byte, varint.MaxSize)
	for _, x := range samples {
		n := varint.EncodeInt64(data, x)
		fmt.Printf("%4d => %x\n", x, data[:n])
	}
}

func testGoVarint() {
	data := make([]byte, binary.MaxVarintLen64)
	xs := []uint64{
		0,
		1,
		2,
		3,
		127,
		128,
		255,
		256,
		math.MaxInt16,
		math.MaxUint16,
		math.MaxInt32,
		math.MaxUint32,
		math.MaxInt64,
		math.MaxUint64,
	}
	for _, x := range xs {
		n := binary.PutUvarint(data, x)
		fmt.Printf("%d => (%x)\n", x, data[:n])
	}
}
