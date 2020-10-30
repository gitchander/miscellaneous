package matrix

import (
	"math/big"
	"sync"
)

/*

|             |
|  v00   v10  |
|             |
|             |
|  v01   v11  |
|             |

*/

type MatrixBig struct {
	v00, v10 *big.Int
	v01, v11 *big.Int

	ts []*big.Int // for multiplication
}

func NewMatrixBig() *MatrixBig {
	return &MatrixBig{
		v00: new(big.Int),
		v10: new(big.Int),
		v01: new(big.Int),
		v11: new(big.Int),

		ts: makeBigInts(8),
	}
}

func makeBigInts(n int) []*big.Int {
	as := make([]*big.Int, n)
	for i := range as {
		as[i] = new(big.Int)
	}
	return as
}

func (p *MatrixBig) InitIdendity() *MatrixBig {
	p.v00.SetInt64(1)
	p.v10.SetInt64(0)
	p.v01.SetInt64(0)
	p.v11.SetInt64(1)
	return p
}

func (p *MatrixBig) Init(v00, v10, v01, v11 int64) *MatrixBig {
	p.v00.SetInt64(v00)
	p.v10.SetInt64(v10)
	p.v01.SetInt64(v01)
	p.v11.SetInt64(v11)
	return p
}

// c = a * b
func (c *MatrixBig) Mul(a, b *MatrixBig) *MatrixBig {
	if a.v00.BitLen() < 100000 {
		return c.consistentlyMul(a, b)
	} else {
		return c.parallelMul(a, b)
	}
}

// c = a * b
func (c *MatrixBig) consistentlyMul(a, b *MatrixBig) *MatrixBig {

	ts := c.ts

	ts[0].Mul(a.v00, b.v00)
	ts[1].Mul(a.v10, b.v01)
	ts[2].Mul(a.v00, b.v10)
	ts[3].Mul(a.v10, b.v11)
	ts[4].Mul(a.v01, b.v00)
	ts[5].Mul(a.v11, b.v01)
	ts[6].Mul(a.v01, b.v10)
	ts[7].Mul(a.v11, b.v11)

	c.v00.Add(ts[0], ts[1])
	c.v10.Add(ts[2], ts[3])
	c.v01.Add(ts[4], ts[5])
	c.v11.Add(ts[6], ts[7])

	return c
}

// c = a * b
func (c *MatrixBig) parallelMul(a, b *MatrixBig) *MatrixBig {

	ts := c.ts

	// run muls parallel
	{
		var wg sync.WaitGroup

		wg.Add(len(ts))

		go mul(&wg, ts[0], a.v00, b.v00)
		go mul(&wg, ts[1], a.v10, b.v01)
		go mul(&wg, ts[2], a.v00, b.v10)
		go mul(&wg, ts[3], a.v10, b.v11)
		go mul(&wg, ts[4], a.v01, b.v00)
		go mul(&wg, ts[5], a.v11, b.v01)
		go mul(&wg, ts[6], a.v01, b.v10)
		go mul(&wg, ts[7], a.v11, b.v11)

		wg.Wait()
	}

	c.v00.Add(ts[0], ts[1])
	c.v10.Add(ts[2], ts[3])
	c.v01.Add(ts[4], ts[5])
	c.v11.Add(ts[6], ts[7])

	return c
}

func mul(wg *sync.WaitGroup, a, b, c *big.Int) {
	a.Mul(b, c)
	wg.Done()
}

func (b *MatrixBig) Power(a *MatrixBig, n int) *MatrixBig {
	b.InitIdendity()
	for i := 0; i < n; i++ {
		b.Mul(b, a)
	}
	return b
}

func (b *MatrixBig) FastPower(a *MatrixBig, n int) *MatrixBig {
	b.InitIdendity()
	for n > 0 {
		if (n & 1) == 1 { // n is odd
			b.Mul(b, a) // b = b * a
		}
		n >>= 1     // n = n / 2
		a.Mul(a, a) // a = a ^ 2
	}
	return b
}

// fibNBig returns n-th number of fibonacci series
func FibNBig(n int) *big.Int {

	a := NewMatrixBig().Init(
		0, 1,
		1, 1,
	)

	b := NewMatrixBig()

	//b.Power(a, n)
	b.FastPower(a, n)

	return b.v01
}
