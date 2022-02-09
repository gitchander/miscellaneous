package main

import (
	"bufio"
	"flag"
	"io"
	"log"

	//"math"
	"os"
)

func main() {

	var index int

	flag.IntVar(&index, "index", 0, "index of sample")

	flag.Parse()

	var w io.Writer = os.Stdout

	sample := samplers[index]
	//sample := new(sample2)

	err := generateBuffer(w, sample)
	//err := generateBufio(w, sample)

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func generateBuffer(w io.Writer, sampler Sampler) error {

	buf := make([]byte, 1)

	for {
		buf[0] = sampler.NextSample()

		_, err := w.Write(buf)
		if err != nil {
			return err
		}
	}
}

func generateBufio(w io.Writer, sampler Sampler) error {

	bw := bufio.NewWriter(w)
	defer bw.Flush()

	for {
		b := sampler.NextSample()

		err := bw.WriteByte(b)
		if err != nil {
			return err
		}
	}
}

type Sampler interface {
	NextSample() uint8
}

type funcSampler struct {
	t  int
	fn func(t int) uint8
}

var _ Sampler = &funcSampler{}

func (p *funcSampler) NextSample() uint8 {

	sample := p.fn(p.t)
	p.t++

	return sample
}

func SamplerByFunc(fn func(t int) uint8) Sampler {
	return &funcSampler{
		fn: fn,
	}
}

type sample2 struct {
	t int
	v int
}

func (p *sample2) NextSample() uint8 {

	var (
		t = p.t
		v = p.v
	)

	v = (v >> 1) + (v >> 4) + t*(((t>>16)|(t>>6))&(69&(t>>9)))

	p.t++
	p.v = v

	return uint8(v)
}

var samplers = []Sampler{
	SamplerByFunc(func(t int) uint8 {
		r := (t | (t>>9 | t>>7)) * t & (t>>11 | t>>9)
		return uint8(r)
	}),
	SamplerByFunc(func(t int) uint8 {
		r := (t * ((t >> 5) | (t >> 8))) >> (t >> 16)
		return uint8(r)
	}),
	SamplerByFunc(func(t int) uint8 {
		r := ((t>>6)|t|(t>>uint(t>>16)))*10 + ((t >> 11) & 7)
		return uint8(r)
	}),
	SamplerByFunc(func(t int) uint8 {
		r := (t>>7|t|t>>6)*10 + 4*(t&t>>13|t>>6)
		return uint8(r)
	}),
	SamplerByFunc(func(t int) uint8 {
		r := t*5&(t>>7) | t*3&(t*4>>10)
		return uint8(r)
	}),
	SamplerByFunc(func(t int) uint8 {
		r := (t>>7|t|t>>6)*10 + 4*(t&t>>13|t>>6)
		return uint8(r)
	}),
	SamplerByFunc(func(t int) uint8 {
		r := (t / (t>>16 | t>>8 + 1) & (t>>5 | t>>11)) - 1 | t*(t>>16|t>>8)
		return uint8(r)
	}),
	SamplerByFunc(func(t int) uint8 {
		r := t * (((t >> 12) | (t >> 8)) & (63 & (t >> 4)))
		return uint8(r)
	}),
	SamplerByFunc(func(t int) uint8 {
		r := t * (((t >> 9) & 10) | ((t >> 11) & 24) ^ ((t >> 10) & 15 & (t >> 15)))
		return uint8(r)
	}),
	SamplerByFunc(func(t int) uint8 {
		r := t * (t>>9 | t>>13) & 16
		return uint8(r)
	}),
	SamplerByFunc(func(t int) uint8 {
		r := (t >> uint(t%7)) & (t >> uint(t%5)) & (t >> uint(t%3))
		return uint8(r)
	}),
	SamplerByFunc(func(t int) uint8 {
		r := (t * (t>>8 + t>>9) * 100)
		return uint8(r)
	}),
	SamplerByFunc(func(t int) uint8 {
		r := (t>>7|t|t>>6)*10 + 4*(t&t>>13|t>>6)
		return uint8(r)
	}),
	SamplerByFunc(func(t int) uint8 {
		r := ((t * (t>>8 | t>>9)) & 46 & (t >> 8)) ^ (t&(t>>13) | (t >> 6))
		return uint8(r)
	}),
}
