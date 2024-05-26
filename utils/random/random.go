package random

import (
	"math/rand"
)

type Rand = rand.Rand

type Random struct {
	baseRand *Rand

	bitsAccum word
	bitsCount int
}

func NewRandom(baseRand *Rand) *Random {
	return &Random{
		baseRand: baseRand,
	}
}

func (r *Random) Rand() *Rand {
	return r.baseRand
}

func (r *Random) need(n int) {
	if r.bitsCount < n {
		r.bitsAccum = randWord(r.baseRand)
		r.bitsCount = bitsPerWord
	}
}

func (r *Random) Bit() Bit {

	n := 1
	r.need(n)

	v := Bit(r.bitsAccum & 1)

	r.bitsAccum >>= n
	r.bitsCount -= n

	return v
}

func (r *Random) Bool() bool {
	return r.Bit() == 1
}

func (r *Random) Byte() byte {

	n := bitsPerByte
	r.need(n)

	v := byte(r.bitsAccum & 0xff)

	r.bitsAccum >>= n
	r.bitsCount -= n

	return v
}

func (r *Random) Uint32() uint32 {
	return r.baseRand.Uint32()
}

func (r *Random) Uint64() uint64 {
	return r.baseRand.Uint64()
}
