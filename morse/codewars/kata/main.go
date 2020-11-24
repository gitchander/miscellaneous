package main

import (
	"fmt"
	"strings"

	"github.com/gitchander/miscellaneous/morse"
)

// https://www.codewars.com/kata/54b724efac3d5402db00065e

var decodeMap map[string]string

func init() {
	decodeMap = morse.MakeDecodeMap()
}

func DecodeMorse(morseCode string) string {
	morseCode = strings.TrimSpace(morseCode)
	var b strings.Builder
	words := strings.Split(morseCode, "   ")
	for i, word := range words {
		if i > 0 {
			b.WriteByte(' ')
		}
		chars := strings.Split(word, " ")
		for _, char := range chars {
			b.WriteString(decodeMap[char])
		}
	}
	return b.String()
}

func main() {
	fmt.Println(DecodeMorse(".... . -.--   .--- ..- -.. ."))
}
