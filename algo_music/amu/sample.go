package amu

import (
	"bufio"
	"io"
)

type Sample interface {
	Channels() []int32
}

type Sampler interface {
	NextSample() byte
}

func WriteSamples(w io.Writer, sampler Sampler) error {

	bw := bufio.NewWriter(w)
	defer bw.Flush()

	for {
		b := sampler.NextSample()

		err := bw.WriteByte(b)
		if err != nil {
			return err
		}
	}

	return nil
}
