package main

import (
	// "fmt"
	"log"
	"math/bits"
	"math/rand"
	"time"

	"bigmy/cbits"
)

func main() {
	testAdd()
	testSub()
	testMul()
	testMulSample()
	testDiv()
}

func testAdd() {
	r := newRandNow()
	for i := 0; i < 1000; i++ {
		var (
			x     = randUint32(r)
			y     = randUint32(r)
			carry = randCarry(r)
		)

		sum1, carryOut1 := cbits.Add32(x, y, carry)
		sum2, carryOut2 := bits.Add32(x, y, carry)

		if sum1 != sum2 {
			log.Fatalf("sums: %d != %d", sum1, sum2)
		}
		if carryOut1 != carryOut2 {
			log.Fatalf("carryOuts: %d != %d", carryOut1, carryOut2)
		}
	}
}

func testSub() {
	r := newRandNow()
	for i := 0; i < 1000; i++ {
		var (
			x     = randUint32(r)
			y     = randUint32(r)
			carry = randCarry(r)
		)

		diff1, borrowOut1 := cbits.Sub32(x, y, carry)
		diff2, borrowOut2 := bits.Sub32(x, y, carry)

		if diff1 != diff2 {
			log.Fatalf("diffs: %d != %d", diff1, diff2)
		}
		if borrowOut1 != borrowOut2 {
			log.Fatalf("borrowOuts: %d != %d", borrowOut1, borrowOut2)
		}
	}
}

func testMul() {
	r := newRandNow()
	for i := 0; i < 1000; i++ {
		var (
			x = randUint32(r)
			y = randUint32(r)
		)

		hi1, lo1 := cbits.Mul32(x, y)
		hi2, lo2 := bits.Mul32(x, y)

		if hi1 != hi2 {
			// log.Printf("x: %x, y: %x, x*y: %x", x, y, uint64(x)*uint64(y))
			// log.Printf("%x-%x", hi1, lo1)
			log.Fatalf("his: %x != %x", hi1, hi2)
		}
		if lo1 != lo2 {
			log.Fatalf("los: %d != %d", lo1, lo2)
		}
	}
}

func testMulSample() {

	var (
		x uint32 = 0xff308a06
		y uint32 = 0x94d0b50f
	)

	hi1, lo1 := cbits.Mul32(x, y)
	hi2, lo2 := bits.Mul32(x, y)

	if hi1 != hi2 {
		log.Fatalf("sums: %x != %x", hi1, hi2)
	}
	if lo1 != lo2 {
		log.Fatalf("carryOuts: %d != %d", lo1, lo2)
	}
}

func testDiv() {
	r := newRandNow()
	for i := 0; i < 1000; i++ {

		var xHi, xLo uint32
		var y32 uint32

		for {
			xHi = randUint32(r)
			xLo = randUint32(r)
			y32 = randUint32(r)
			if (y32 != 0) && (y32 > xHi) {
				break
			}
		}

		quo, rem := cbits.Div32(xHi, xLo, y32)

		var (
			quo1 = uint64(quo)
			rem1 = uint64(rem)
		)

		x := (uint64(xHi) << 32) | uint64(xLo)
		y := uint64(y32)

		//fmt.Printf("x:%d, y:%d\n", x, y)

		var (
			quo2 = x / uint64(y)
			rem2 = x % uint64(y)
		)

		if quo1 != quo2 {
			log.Fatalf("quo: (%d != %d)", quo1, quo2)
		}
		if rem1 != rem2 {
			log.Fatalf("rem: (%d != %d)", rem1, rem2)
		}
	}
}

//------------------------------------------------------------------------------
// Rand
func newRandNow() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randBool(r *rand.Rand) bool {
	return (r.Int() & 1) == 1
}

func randUint32(r *rand.Rand) uint32 {
	return r.Uint32() >> (r.Intn(32))
}

func randCarry(r *rand.Rand) uint32 {
	carry := uint32(r.Intn(2)) // [0, 1]
	return carry
}

//------------------------------------------------------------------------------
