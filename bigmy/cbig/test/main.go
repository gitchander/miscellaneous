package main

import (
	"fmt"
	"math/big"
)

func testBig() {
	var a = big.NewInt(123)
	fmt.Printf("%x\n", a)
}

func main() {
	testBig()
}
