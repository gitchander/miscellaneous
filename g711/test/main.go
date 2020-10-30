package main

import (
	"fmt"
	"log"

	"github.com/gitchander/miscellaneous/g711"
)

func main() {
	testAlaw()
	testUlaw()
}

func testAlaw() {

	for i := 0; i < 256; i++ {
		b := byte(i)

		valuePCM := g711.AlawToLinear(b)

		fmt.Printf("%x: %d\n", b, valuePCM)

		x := g711.LinearToAlaw(valuePCM)
		if b != x {
			log.Fatalf("%x != %x\n", b, x)
			return
		}
	}
}

func testUlaw() {
	for i := 0; i < 256; i++ {
		b := byte(i)

		valuePCM := g711.UlawToLinear(b)

		fmt.Printf("%x: %d\n", b, valuePCM)

		x := g711.LinearToUlaw(valuePCM)
		if b != x {
			//log.Fatalf("%x != %x\n", b, x)
			log.Printf("%x != %x\n", b, x)
		}
	}
}
