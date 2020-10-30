package bits

import (
	"fmt"
)

const sizeOfUint64 = 64

type IBits interface {
	Len() int
	Cap() int
	Bit(i int) uint
	SetBit(i int, b uint)
}

// sort: IntSlice

// BitSlice
type Bits struct {
	vs []uint64

	start int
	stop  int
}

var _ IBits = Bits{}

// func MakeBits(len_, cap_ int) Bits {
// 	return Bits{}
// }

func MakeBits(length int) Bits {

	if length < 0 {
		panic("makeslice: len out of range")
	}

	const n = sizeOfUint64

	size := (length + (n - 1)) / n

	return Bits{
		vs:    make([]uint64, size),
		start: 0,
		stop:  length,
	}
}

func (b Bits) Len() int {
	return (b.stop - b.start)
}

func (b Bits) Cap() int {
	return ((len(b.vs) * sizeOfUint64) - b.start)
}

func (b Bits) checkIndex(i int) error {
	if i < 0 {
		return fmt.Errorf("invalid slice index %d (index must be non-negative)", i)
	}
	length := b.Len()
	if i >= length {
		return fmt.Errorf("index out of range [%d] with length %d", i, length)
	}
	return nil
}

func (b Bits) checkBounds(lo, hi int) error {
	capacity := b.Cap()
	if (lo < 0) || (lo > hi) || (hi > capacity) {
		return fmt.Errorf("slice bounds out of range [%d:%d]", lo, hi)
	}
	return nil
}

// Bit gets i-th bit of bit array (p[i])
func (p Bits) Bit(i int) uint {

	err := p.checkIndex(i)
	if err != nil {
		panic(err)
	}

	index, ibit := quoRemInt((p.start + i), sizeOfUint64)

	return getBitUint64(p.vs[index], ibit)
}

// SetBit sets i-th bit of bit array (p[i] = b)
func (p Bits) SetBit(i int, b uint) {

	err := p.checkIndex(i)
	if err != nil {
		panic(err)
	}

	index, ibit := quoRemInt((p.start + i), sizeOfUint64)

	p.vs[index] = setBitUint64(p.vs[index], ibit, b)
}

func (b Bits) String() string {

	length := b.Len()

	data := make([]byte, length)

	for i := 0; i < length; i++ {

		index, ibit := quoRemInt((b.start + i), sizeOfUint64)

		b := getBitUint64(b.vs[index], ibit)

		switch b {
		case 0:
			data[i] = '0'
		case 1:
			data[i] = '1'
		}
	}

	return string(data)
}

func (b Bits) Slice(lo, hi int) Bits {
	err := b.checkBounds(lo, hi)
	if err != nil {
		panic(err)
	}
	var (
		start = b.start + lo
		stop  = b.start + hi
	)
	return Bits{
		vs:    b.vs,
		start: start,
		stop:  stop,
	}
}

func (p Bits) Bools() []bool {
	vs := make([]bool, p.Len())
	for i := range vs {

		index, ibit := quoRemInt((p.start + i), sizeOfUint64)

		b := getBitUint64(p.vs[index], ibit)

		switch b {
		case 0:
			vs[i] = false
		case 1:
			vs[i] = true
		}
	}
	return vs
}

func (b Bits) Bytes(ba *BitArray) {

}

type BitArray struct {
	Bytes     []byte
	BitLength int
}

func Append(b Bits, vs ...bool) Bits {

	for _, v := range vs {
		_ = v
	}

	return b
}

func Parse(s string) (b Bits, err error) {

	return
}

const oneUint64 uint64 = 1

func getBitUint64(x uint64, i int) uint {
	if (x & (oneUint64 << i)) != 0 {
		return 1
	}
	return 0
}

func setBitUint64(x uint64, i int, b uint) uint64 {
	switch b {
	case 0:
		x &= ^(oneUint64 << i)
	case 1:
		x |= (oneUint64 << i)
	default:
		err := ErrInvalidBit(b)
		panic(err)
	}
	return x
}

func quoRemInt(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}

func ErrInvalidBit(b uint) error {
	return fmt.Errorf("invalid bit value %d", b)
}
