package cbits

import (
	"errors"
	"math/bits"
)

const useGoVersion = false

const (
	shift16 = 16
	two16   = 1 << shift16
	mask16  = two16 - 1
)

const (
	shift32 = 32
	two32   = 1 << shift32
	mask32  = two32 - 1
)

func splitUint32(x uint32) (hi, lo uint32) {
	hi = x >> shift16
	lo = x & mask16
	return
}

func mergeUint32(hi, lo uint32) (x uint32) {
	x |= (hi << shift16)
	x |= (lo & mask16)
	return
}

func splitUint64(x uint64) (hi, lo uint64) {
	hi = x >> shift32
	lo = x & mask32
	return
}

func mergeUint64(hi, lo uint64) (x uint64) {
	x |= (hi << shift32)
	x |= (lo & mask32)
	return
}

// var (
// 	//notUint32 = notUint32v1
// 	notUint32 = notUint32v2
// )

// func notUint32v1(a uint32) uint32 {
// 	return ^a
// }

// func notUint32v2(a uint32) uint32 {
// 	return mask32 ^ a
// }

// ------------------------------------------------------------------------------
// Add32
// ------------------------------------------------------------------------------
// Add32 returns the sum with carry of x, y and carry: sum = x + y + carry.
// The carry input must be 0 or 1; otherwise the behavior is undefined.
// The carryOut output is guaranteed to be 0 or 1.
//
// This function's execution time does not depend on the inputs.
func Add32(x, y, carry uint32) (sum, carryOut uint32) {
	if useGoVersion {
		return add32_Go(x, y, carry)
	} else {
		return add32_My(x, y, carry)
	}
}

func add32_Go(x, y, carry uint32) (sum, carryOut uint32) {
	sum = x + y + carry
	// The sum will overflow if both top bits are set (x & y) or if one of them
	// is (x | y), and a carry from the lower place happened. If such a carry
	// happens, the top bit will be 1 + 0 + 1 = 0 (&^ sum).

	carryOut = ((x & y) | ((x | y) &^ sum)) >> 31
	//carryOut = ((x & y) | ((x | y) & notUint32(sum))) >> 31

	return
}

func add32_My(x, y, carry uint32) (sum, carryOut uint32) {

	x1, x0 := splitUint32(x)
	y1, y0 := splitUint32(y)

	c := carry

	c, z0 := splitUint32(x0 + y0 + c)
	c, z1 := splitUint32(x1 + y1 + c)

	sum = mergeUint32(z1, z0)
	carryOut = c

	return
}

// ------------------------------------------------------------------------------
// Sub32
// ------------------------------------------------------------------------------
// Sub32 returns the difference of x, y and borrow, diff = x - y - borrow.
// The borrow input must be 0 or 1; otherwise the behavior is undefined.
// The borrowOut output is guaranteed to be 0 or 1.
//
// This function's execution time does not depend on the inputs.
func Sub32(x, y, borrow uint32) (diff, borrowOut uint32) {
	if useGoVersion {
		return sub32_Go(x, y, borrow)
	} else {
		return sub32_My(x, y, borrow)
	}
}

func sub32_Go(x, y, borrow uint32) (diff, borrowOut uint32) {
	diff = x - y - borrow
	// The difference will underflow if the top bit of x is not set and the top
	// bit of y is set (^x & y) or if they are the same (^(x ^ y)) and a borrow
	// from the lower place happens. If that borrow happens, the result will be
	// 1 - 1 - 1 = 0 - 0 - 1 = 1 (& diff).
	borrowOut = ((^x & y) | (^(x ^ y) & diff)) >> 31
	return
}

func sub32_My(x, y, borrow uint32) (diff, borrowOut uint32) {
	// todo
	return sub32_Go(x, y, borrow)
}

// ------------------------------------------------------------------------------
// Mul32
// ------------------------------------------------------------------------------
// Mul32 returns the 64-bit product of x and y: (hi, lo) = x * y
// with the product bits' upper half returned in hi and the lower
// half returned in lo.
//
// This function's execution time does not depend on the inputs.
func Mul32(x, y uint32) (hi, lo uint32) {
	if useGoVersion {
		return mul32_Go(x, y)
	} else {
		return mul32_My(x, y)
	}
}

func mul32_Go(x, y uint32) (hi, lo uint32) {

	const (
		mask  = 1<<16 - 1
		shift = 16
	)

	var (
		x0 = x & mask
		x1 = x >> shift
	)

	var (
		y0 = y & mask
		y1 = y >> shift
	)

	w0 := x0 * y0
	t := x1*y0 + (w0 >> shift)
	w1 := t & mask
	w2 := t >> shift
	w1 += x0 * y1

	hi = x1*y1 + w2 + (w1 >> shift)
	lo = x * y

	return
}

func mul32_My(x, y uint32) (hi, lo uint32) {

	x1, x0 := splitUint32(x)
	y1, y0 := splitUint32(y)

	c00, v00 := splitUint32(x0 * y0)
	c01, v01 := splitUint32(x0 * y1)
	c10, v10 := splitUint32(x1 * y0)
	c11, v11 := splitUint32(x1 * y1)

	z0 := v00
	c, z1 := splitUint32(v01 + v10 + c00)
	c, z2 := splitUint32(v11 + c01 + c10 + c)
	z3 := c11 + c

	hi = mergeUint32(z3, z2)
	lo = mergeUint32(z1, z0)

	return
}

// ------------------------------------------------------------------------------
var (
	divideError   = errors.New("divideError")
	overflowError = errors.New("overflowError")
)

// ------------------------------------------------------------------------------
// Div32
// ------------------------------------------------------------------------------
// Div32 returns the quotient and remainder of (hi, lo) divided by y:
// quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits' upper
// half in parameter hi and the lower half in parameter lo.
// Div32 panics for y == 0 (division by zero) or y <= hi (quotient overflow).
func Div32(hi, lo, y uint32) (quo, rem uint32) {
	if useGoVersion {
		return div32_Go(hi, lo, y)
	} else {
		return div32_My(hi, lo, y)
	}
}

func div32_Go(hi, lo, y uint32) (quo, rem uint32) {
	const (
		two16  = 1 << 16
		mask16 = two16 - 1
	)
	if y == 0 {
		panic(divideError)
	}
	if y <= hi {
		panic(overflowError)
	}

	s := uint(bits.LeadingZeros32(y))
	y <<= s

	yn1 := y >> 16
	yn0 := y & mask16
	un16 := hi<<s | lo>>(32-s)
	un10 := lo << s
	un1 := un10 >> 16
	un0 := un10 & mask16
	q1 := un16 / yn1
	rhat := un16 - q1*yn1

	for q1 >= two16 || q1*yn0 > two16*rhat+un1 {
		q1--
		rhat += yn1
		if rhat >= two16 {
			break
		}
	}

	un21 := un16*two16 + un1 - q1*y
	q0 := un21 / yn1
	rhat = un21 - q0*yn1

	for q0 >= two16 || q0*yn0 > two16*rhat+un0 {
		q0--
		rhat += yn1
		if rhat >= two16 {
			break
		}
	}

	return q1*two16 + q0, (un21*two16 + un0 - q0*y) >> s
}

func div32_My(hi, lo, y uint32) (quo, rem uint32) {
	// todo
	return div32_Go(hi, lo, y)
}

// ------------------------------------------------------------------------------
// Rem32
// ------------------------------------------------------------------------------
// Rem32 returns the remainder of (hi, lo) divided by y. Rem32 panics
// for y == 0 (division by zero) but, unlike Div32, it doesn't panic
// on a quotient overflow.
func Rem32(hi, lo, y uint32) uint32 {
	// We scale down hi so that hi < y, then use Div32 to compute the
	// rem with the guarantee that it won't panic on quotient overflow.
	// Given that
	//   hi ≡ hi%y    (mod y)
	// we have
	//   hi<<32 + lo ≡ (hi%y)<<32 + lo    (mod y)
	_, rem := Div32(hi%y, lo, y)
	return rem
}
