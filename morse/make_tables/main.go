package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"unicode"
)

func main() {
	makeTables()
}

const (
	dash = iota
	dot
)

var charMapEnglish = map[rune][]int{
	'a':  {dot, dash},
	'b':  {dash, dot, dot, dot},
	'c':  {dash, dot, dash, dot},
	'd':  {dash, dot, dot},
	'e':  {dot},
	'f':  {dot, dot, dash, dot},
	'g':  {dash, dash, dot},
	'h':  {dot, dot, dot, dot},
	'i':  {dot, dot},
	'j':  {dot, dash, dash, dash},
	'k':  {dash, dot, dash},
	'l':  {dot, dash, dot, dot},
	'm':  {dash, dash},
	'n':  {dash, dot},
	'o':  {dash, dash, dash},
	'p':  {dot, dash, dash, dot},
	'q':  {dash, dash, dot, dash},
	'r':  {dot, dash, dot},
	's':  {dot, dot, dot},
	't':  {dash},
	'u':  {dot, dot, dash},
	'v':  {dot, dot, dot, dash},
	'w':  {dot, dash, dash},
	'x':  {dash, dot, dot, dash},
	'y':  {dash, dot, dash, dash},
	'z':  {dash, dash, dot, dot},
	'1':  {dot, dash, dash, dash, dash},
	'2':  {dot, dot, dash, dash, dash},
	'3':  {dot, dot, dot, dash, dash},
	'4':  {dot, dot, dot, dot, dash},
	'5':  {dot, dot, dot, dot, dot},
	'6':  {dash, dot, dot, dot, dot},
	'7':  {dash, dash, dot, dot, dot},
	'8':  {dash, dash, dash, dot, dot},
	'9':  {dash, dash, dash, dash, dot},
	'0':  {dash, dash, dash, dash, dash},
	'.':  {dot, dash, dot, dash, dot, dash},
	',':  {dash, dash, dot, dot, dash, dash},
	':':  {dash, dash, dash, dot, dot, dot},
	'?':  {dot, dot, dash, dash, dot, dot},
	'\'': {dot, dash, dash, dash, dash, dot},
	'-':  {dash, dot, dot, dot, dot, dash},
	'/':  {dash, dot, dot, dash, dot},
	'(':  {dash, dot, dash, dash, dot},
	')':  {dash, dot, dash, dash, dot, dash},
	'+':  {dot, dash, dot, dash, dot},
	'@':  {dot, dash, dash, dot, dash, dot},
}

func makeTables() {

	var ps []decodePair

	for char, vs := range charMapEnglish {

		var b strings.Builder
		for _, v := range vs {
			if v == dash {
				b.WriteByte('-')
			} else if v == dot {
				b.WriteByte('.')
			}
		}

		cu := unicode.ToUpper(char)

		if cu > 127 {
			log.Fatalf("invalid char %q (%U)", cu, cu)
		}

		pair := decodePair{
			Value: b.String(),
			Char:  cu,
		}
		ps = append(ps, pair)
	}

	sort.Sort(byChar(ps))

	fmt.Println("morse encode table:")
	for _, pair := range ps {
		fmt.Printf("{%q, %q},\n", pair.Char, pair.Value)
	}
	fmt.Println()

	fmt.Println("morse encode map:")
	fmt.Println("var morseEncodeMap = map[rune]string {")
	for _, pair := range ps {
		fmt.Printf("%q: %q,\n", pair.Char, pair.Value)
	}
	fmt.Println()

	fmt.Println("morse decode map:")
	fmt.Println("var morseDecodeMap = map[string]rune {")
	for _, pair := range ps {
		fmt.Printf("%q: %q,\n", pair.Value, pair.Char)
	}
	fmt.Println()
}

type decodePair struct {
	Value string
	Char  rune
}

type byChar []decodePair

func (p byChar) Len() int           { return len(p) }
func (p byChar) Less(i, j int) bool { return p[i].Char < p[j].Char }
func (p byChar) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
