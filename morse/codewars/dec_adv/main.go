package main

import (
	"log"
)

func main() {
	testSamples()
}

type Sample struct {
	Value  string
	Result string
}

func testSamples() {

	var samples = []Sample{
		{
			Value:  "1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011",
			Result: "HEY JUDE",
		},
	}

	for _, sample := range samples {

		code := DecodeBits(sample.Value)
		res := DecodeMorse(code)

		if res != sample.Result {
			log.Fatalf("%q != %q", res, sample.Result)
		}
	}
}
