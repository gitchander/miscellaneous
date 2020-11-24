package main

import (
	"fmt"
	"math"
	"sort"
)

// const (
// 	unitSignalON  = '1'
// 	unitSignalOFF = '0'
// )

// const (
// 	unitSignalON  = '='
// 	unitSignalOFF = '.'
// )

// func decodeSignal(b byte) (signal bool, ok bool) {
// 	switch b {
// 	case '.':
// 		signal = false
// 		ok = true
// 	case '=':
// 		signal = true
// 		ok = true
// 	default:
// 		ok = false
// 	}
// 	return
// }

func decodeSignal(b byte) (signal bool, ok bool) {
	switch b {
	case '0':
		signal = false
		ok = true
	case '1':
		signal = true
		ok = true
	default:
		ok = false
	}
	return
}

const (
	durDot  = 1 // unit
	durDash = 3 // units
)

const (
	durSymbolSpace = 1
	durLetterSpace = 3
	durWordSpace   = 7
)

func mapKeys(m map[int]int) []int {
	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

type solid struct {
	signalIsON bool
	width      int
}

func parseSolids(bits string) []solid {

	var ds []solid

	walkSolidBytes(bits, func(b byte, width int) bool {
		signal, ok := decodeSignal(b)
		if !ok {
			panic("invalid bits")
		}
		d := solid{
			signalIsON: signal,
			width:      width,
		}
		ds = append(ds, d)
		return true
	})

	return ds
}

func walkSolidBytes(bits string, f func(b byte, width int) bool) {

	var (
		k = 0
		n = len(bits)
	)

	for i := 1; i < n; i++ {
		if bits[i] == bits[k] {
			continue
		}
		width := i - k
		if !f(bits[k], width) {
			return
		}

		k = i
	}

	if k < n {
		width := n - k
		if !f(bits[k], width) {
			return
		}
	}
}

func trimSolids(ds []solid) []solid {

	for len(ds) > 0 {
		if ds[0].signalIsON {
			break
		}
		ds = ds[1:]
	}

	for len(ds) > 0 {
		n := len(ds)
		if ds[n-1].signalIsON {
			break
		}
		ds = ds[:n-1]
	}

	return ds
}

func calcN(ds []solid) int {

	var (
		mapOFF = make(map[int]int)
		mapON  = make(map[int]int)
	)

	for _, d := range ds {
		if d.signalIsON {
			mapON[d.width]++
		} else {
			mapOFF[d.width]++
		}
	}

	fmt.Println("map OFF:", mapOFF)
	fmt.Println("map ON: ", mapON)

	csOFF := mapKeys(mapOFF)
	csON := mapKeys(mapON)

	sort.Ints(csOFF)
	sort.Ints(csON)

	fmt.Println("keys OFF:", csOFF)
	fmt.Println("keys ON: ", csON)

	var sum float64
	var count int

	if m := len(csON); m > 0 {
		sum += float64(csON[0]) / durDot
		count++
		if m > 1 {
			sum += float64(csON[1]) / durDash
			count++
		}
	}
	if m := len(csOFF); m > 0 {
		sum += float64(csOFF[0]) / durSymbolSpace
		count++
		if m > 1 {
			sum += float64(csOFF[1]) / durLetterSpace
			count++
			if m > 2 {
				sum += float64(csOFF[2]) / durWordSpace
				count++
			}
		}
	}

	Nf := sum / float64(count)
	N := int(math.Round(Nf))
	fmt.Println("Nf:", Nf)
	fmt.Println("N:", N)

	return N
}

func DecodeBits(bits string) string {

	ds := parseSolids(bits)
	ds = trimSolids(ds)

	n := calcN(ds)
	fmt.Println(n)

	//fmt.Println("solids:", ds)

	vs := make([]byte, 0, len(ds))

	for _, d := range ds {

		w := (d.width / n)

		if d.signalIsON {
			switch w {
			case durDot: // 1
				vs = append(vs, '.')
			case durDash: // 3 // units
				vs = append(vs, '-')
			default:
				vs = append(vs, '.')
			}
		} else {
			switch w {
			case durSymbolSpace: // 1
			case durLetterSpace: // 3
				vs = append(vs, ' ')
			case durWordSpace: // 7
				vs = append(vs, []byte("   ")...)
			default:
			}
		}
	}

	return string(vs)
}

func DecodeMorse(morseCode string) string {
	return ""
}

func main() {
	s := "1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"
	//s := "=.=.=.=...=...===.=.===.===.......=.===.===.===...=.=.===...===.=.=...="
	//s := "=.===...=.===.=.=...=.===.=.=.......===...=.=.=.=...=.......===...=...=.=.=...===.......=.=.=...===...=.===.=...=.=...===.=...===.===.=...=.=.=.......=.===.===...===.===.===...=.=.===...=.===.=.=...===.=.=.......===.=.=.=...=.......=.=.=.===...=.===...=.===.=.=...=.=...===.=.=.......===...===.===.===.......===...=.=.=.=...=.......=.===.===.=...===.===.===...=.=...===.=...===.......===...=.=.=.=...=.===...===.......===...=.=.=.=...=...===.=.===.===.......===.=.===.=...===.===.===...=.=.===...=.===.=.=...===.=.=.......===.=.=.=...=.......=.===.=...=...=.===.=.=...=.=...=.===...===.=.=.=...=.===.=.=...===.=.===.===.......===.=.=...=...===.=.===.=...===.===.===...===.=.=...=...===.=.=.......=.===...=.=.=.......===.=.=...=...=.=.=...===.=.===.=...=.===.=...=.=...===.=.=.=...=...===.=.=.......=.===...===.=.=.=...===.===.===...=.=.=.===...=...===.===.=.=.===.===.......=.=.=...===.===.===.......===.=.===.===...===.===.===...=.=.===.......===.===...=.===...===.=.===.===.......=.=.=...===.=.===...=.=...=.===.===.=.......===.=.===.=...=.=.=.=...=...===.=.===.=...===.=.===...=.=...===.=...===.===.=.......=.=.===.=...===.===.===...=.===.=.......=...=.===.=...=.===.=...===.===.===...=.===.=...=.=.=.......=.===...===.=...===.=.=.......=...===.=.=.===...===.=.===.=...=...=.===.===.=...===...=.=...===.===.===...===.=...=.=.=...===.===.=.=.===.===.......=.===.===.===...=.=.===...=.=.=...===.......===.=.=...===.===.===.......===.=.===.===...===.===.===...=.=.===...=.===.=.......===.=.=.=...=...=.=.=...===.......=.=...===.=.......=.=.===.=...=.=...===.===.=...=.=.===...=.===.=...=.=...===.=...===.===.=.......===.===.===...=.=.===...===.......=.===.===...=.=.=.=...=.===...===.......===...=.=.=.=...=.......===.===...=...=.=.=...=.=.=...=.===...===.===.=...=.......=.=...=.=.="

	p := DecodeBits(s)

	fmt.Println(p)
}
