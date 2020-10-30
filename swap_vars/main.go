package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
	"unsafe"
)

// How to swap two numbers without using a temporary variable?

func main() {
	//testSimple()
	testHard()
	//testRand()
	//testRandInt()
	//testSwapFloat()
	//problemOneVar()
}

func testSimple() {

	var a, b int8 = -128, -1

	fmt.Println(a, b)

	//------------------------
	// Swap "+,-":
	// a = a + b
	// b = a - b
	// a = a - b
	//------------------------
	// Swap "XOR":
	// a = a ^ b
	// b = a ^ b
	// a = a ^ b
	//------------------------
	// Swap "XOR":
	a ^= b
	b ^= a
	a ^= b
	//------------------------

	fmt.Println(a, b)
}

func testHard() {

	start := time.Now()

	//-----------------------------------------
	// const (
	// 	minVal = int(math.MinInt8)
	// 	maxVal = int(math.MaxInt8)
	// )

	// const (
	// 	minVal = int(0)
	// 	maxVal = int(math.MaxUint8)
	// )

	const (
		minVal = int(math.MinInt16)
		maxVal = int(math.MaxInt16)
	)

	// const (
	// 	minVal = int(math.MinInt32)
	// 	maxVal = int(math.MaxInt32)
	// )
	//-----------------------------------------

	for valA := minVal; valA <= maxVal; valA++ {
		for valB := minVal; valB <= maxVal; valB++ {

			//-----------------------------------------
			// var (
			// 	_a = int8(valA)
			// 	_b = int8(valB)
			// )

			// var (
			// 	_a = uint8(valA)
			// 	_b = uint8(valB)
			// )

			var (
				a0 = int16(valA)
				b0 = int16(valB)
			)

			// var (
			// 	_a = int32(valA)
			// 	_b = int32(valB)
			// )
			//-----------------------------------------

			var (
				a = a0
				b = b0
			)

			//------------------------
			// Swap "+,-":
			// a = a + b
			// b = a - b
			// a = a - b
			//------------------------
			// Swap "XOR":
			a = a ^ b
			b = a ^ b
			a = a ^ b
			//------------------------

			if (a != b0) || (b != a0) {
				log.Fatalf("(%d,%d) != (%d,%d)", a, b, a0, b0)
			}
		}
	}

	fmt.Println(time.Since(start))
}

func testRand() {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	randValue := randInt32
	//randValue := randInt64

	for i := 0; i < 1000000; i++ {
		var (
			a0 = randValue(r)
			b0 = randValue(r)
		)

		var (
			a = a0
			b = b0
		)

		//fmt.Println(a, b)
		// if (a > 0) && (b > 0) && (a+b < 0) {
		// 	fmt.Printf("%d + %d = %d\n", a, b, a+b)
		// }

		//------------------------
		// Swap "+,-"
		// a = a + b
		// b = a - b
		// a = a - b
		//------------------------
		// Swap "XOR"
		// a = a ^ b
		// b = a ^ b
		// a = a ^ b
		//------------------------
		// Swap "XOR":
		a ^= b
		b ^= a
		a ^= b
		//------------------------

		if (a != b0) || (b != a0) {
			log.Fatalf("(%d,%d) != (%d,%d)", a, b, a0, b0)
		}
	}
}

func testRandInt() {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	var swap swapIntFunc

	//swap = swapIntTemp
	//swap = swapIntGo
	swap = swapIntAdd
	//swap = swapIntXOR_1
	//swap = swapIntXOR_2

	//randValue := randInt32
	randValue := randInt64

	for i := 0; i < 1000000; i++ {
		var (
			a0 = int(randValue(r))
			b0 = int(randValue(r))
		)

		var (
			a = a0
			b = b0
		)

		swap(&a, &b)

		if (a != b0) || (b != a0) {
			log.Fatalf("(%d,%d) != (%d,%d)", a, b, a0, b0)
		}
	}
}

func randBool(r *rand.Rand) bool {
	return (r.Int() & 1) == 1
}

func randInt32(r *rand.Rand) int32 {
	v := int32(r.Int63() & 0x7FFFFFFF)
	v = v >> uint(r.Intn(32))
	if randBool(r) {
		v = -v
	}
	return v
}

func randInt64(r *rand.Rand) int64 {
	v := r.Int63()
	v = v >> uint(r.Intn(63))
	if randBool(r) {
		v = -v
	}
	return v
}

type swapIntFunc = func(*int, *int)

func swapIntTemp(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func swapIntGo(a, b *int) {
	*a, *b = *b, *a
}

func swapIntAdd(a, b *int) {
	if a == b {
		return
	}
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func swapIntXOR_1(a, b *int) {
	if a == b {
		return
	}
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func swapIntXOR_2(a, b *int) {
	if a == b {
		return
	}
	*a ^= *b
	*b ^= *a
	*a ^= *b
}

func testSwapFloat() {

	var a, b float32
	//var a, b float64

	a, b = 3.1415, -123.4567

	fmt.Println(a, b)

	var (
		pa = unsafe.Pointer(&a)
		pb = unsafe.Pointer(&b)
	)

	var (
		sizeA = unsafe.Sizeof(a)
		sizeB = unsafe.Sizeof(b)
	)

	if (sizeA == 4) && (sizeB == 4) {
		swapUint32((*uint32)(pa), (*uint32)(pb))
	} else if (sizeA == 8) && (sizeB == 8) {
		swapUint64((*uint64)(pa), (*uint64)(pb))
	} else {
		panic("invalid size of variables")
	}

	fmt.Println(a, b)
}

func swapUint32(a, b *uint32) {
	if a == b {
		return
	}
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func swapUint64(a, b *uint64) {
	if a == b {
		return
	}
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func problemOneVar() {

	swapIntBad := func(a, b *int) {
		*a = *a + *b
		*b = *a - *b
		*a = *a - *b
	}

	swapIntGood := func(a, b *int) {
		if a == b {
			return
		}
		*a = *a + *b
		*b = *a - *b
		*a = *a - *b
	}

	fmt.Print("Bad:")
	var x = 10
	fmt.Printf("(%d,%d)", x, x)
	swapIntBad(&x, &x)
	fmt.Printf("->(%d,%d)\n", x, x)

	fmt.Print("Good:")
	x = 10
	fmt.Printf("(%d,%d)", x, x)
	swapIntGood(&x, &x)
	fmt.Printf("->(%d,%d)\n", x, x)
}
