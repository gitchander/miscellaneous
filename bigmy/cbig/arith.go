package cbig

import (
	"github.com/gitchander/miscellaneous/bigmy/cbits"
)

// ----------------------------------------------------------------------------
// Elementary operations on words
//
// These operations are used by the vector operations below.

// z1<<_W + z0 = x*y
func mulWW(x, y Word) (z1, z0 Word) {
	hi, lo := cbits.Mul32(uint32(x), uint32(y))
	return Word(hi), Word(lo)
}

// z1<<_W + z0 = x*y + c
func mulAddWWW(x, y, c Word) (z1, z0 Word) {
	hi, lo := cbits.Mul32(uint32(x), uint32(y))
	var cc uint32
	lo, cc = cbits.Add32(lo, uint32(c), 0)
	return Word(hi + cc), Word(lo)
}

// q = (u1<<_W + u0 - r)/v
func divWW(u1, u0, v Word) (q, r Word) {
	qq, rr := cbits.Div32(uint32(u1), uint32(u0), uint32(v))
	return Word(qq), Word(rr)
}

// The resulting carry c is either 0 or 1.
func addVV(z, x, y []Word) (c Word) {
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x) && i < len(y); i++ {
		zi, cc := cbits.Add32(uint32(x[i]), uint32(y[i]), uint32(c))
		z[i] = Word(zi)
		c = Word(cc)
	}
	return
}

// The resulting carry c is either 0 or 1.
func addVW(z, x []Word, y Word) (c Word) {
	c = y
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x); i++ {
		zi, cc := cbits.Add32(uint32(x[i]), uint32(c), 0)
		z[i] = Word(zi)
		c = Word(cc)
	}
	return
}

// The resulting carry c is either 0 or 1.
func subVV(z, x, y []Word) (c Word) {
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x) && i < len(y); i++ {
		zi, cc := cbits.Sub32(uint32(x[i]), uint32(y[i]), uint32(c))
		z[i] = Word(zi)
		c = Word(cc)
	}
	return
}

func subVW(z, x []Word, y Word) (c Word) {
	c = y
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x); i++ {
		zi, cc := cbits.Sub32(uint32(x[i]), uint32(c), 0)
		z[i] = Word(zi)
		c = Word(cc)
	}
	return
}

func mulAddVWW(z, x []Word, y, r Word) (c Word) {
	c = r
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x); i++ {
		c, z[i] = mulAddWWW(x[i], y, c)
	}
	return
}

func addMulVVW(z, x []Word, y Word) (c Word) {
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x); i++ {
		z1, z0 := mulAddWWW(x[i], y, z[i])
		lo, cc := cbits.Add32(uint32(z0), uint32(c), 0)
		c, z[i] = Word(cc), Word(lo)
		c += z1
	}
	return
}
