package core

import (
	"bytes"
)

func CalcResult(r Result) (string, error) {

	h := newHash()

	var b bytes.Buffer
	for _, gb := range r.Blocks {

		h.Reset()
		h.Write(gb.Data)
		h.Sum(nil)

		var (
			sum = h.Sum(nil)
			pos = gb.Pos
		)

		b.Write(sum[pos[0]:pos[1]])
	}
	return b.String(), nil
}
