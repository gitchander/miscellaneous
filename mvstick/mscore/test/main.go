package main

import (
	"fmt"
	"image"
	"log"

	"mvstick/mscore"
)

func _() {
	var r image.Rectangle
	_ = r
}

func main() {
	//textIntsEncode()
	testParseLines()
}

func test() {
	m := mscore.Move{BrickIndex: 2, Offset: -5}
	fmt.Printf("%s\n", m)

	m, err := mscore.ParseMove("J+28")
	checkError(err)

	fmt.Printf("%s, %#v\n", m, m)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testParseLines() {

	// // 51 moves
	// lines := []string{
	// 	"BCC-F-",
	// 	"BDE-FG",
	// 	"BDEAAG",
	// 	"HHHI-G",
	// 	"--KIJJ",
	// 	"LLKMM-",
	// }
	// g, err := mscore.ParseLines(lines, '-')
	// checkError(err)

	// https://www.michaelfogleman.com/rush/

	samples := []string{
		"IBBxooIooLDDJAALooJoKEEMFFKooMGGHHHM", // 60 moves
		"BBoKMxDDDKMoIAALooIoJLEEooJFFNoGGoxN", // 58 moves
		"GBBoLoGHIoLMGHIAAMCCCKoMooJKDDEEJFFo", // 51 moves
	}

	g, err := mscore.ParseGrid6x6(samples[2])
	checkError(err)

	fmt.Println(g.Printable())

	mscore.Solve(g)

	fmt.Println(g.Printable())
}
