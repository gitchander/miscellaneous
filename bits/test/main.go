package main

import (
	"fmt"

	"bits"
)

func main() {
	//testSlice()
	testBits()
}

func testSlice() {
	n := 127
	var a = make([]int, n)
	fmt.Println("len:", len(a))
	fmt.Println("cap:", cap(a))

	//a[0] = 1

	fmt.Println(a)

	k := 5

	b := a[k:2]
	fmt.Println(b)
}

func testBits() {
	var b bits.Bits
	b = bits.MakeBits(11)
	fmt.Println("len:", b.Len())
	fmt.Println("cap:", b.Cap())

	b.SetBit(5, 1)
	fmt.Println(b)

	c := b.Slice(15, 64)
	fmt.Println(c)

	d := c.Slice(0, 40)
	fmt.Println(d.Len(), d.Cap())
	fmt.Println(d)

	c.SetBit(2, 1)

	vs := c.Bools()
	fmt.Println(vs)
}
