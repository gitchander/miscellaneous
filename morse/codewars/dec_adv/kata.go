package main

import (
	"log"
	"strings"
)

// https://www.codewars.com/kata/54b72c16cd7f5154e9000457

var morseDecodeMap = map[string]rune{
	".----.": '\'',
	"-.--.":  '(',
	"-.--.-": ')',
	".-.-.":  '+',
	"--..--": ',',
	"-....-": '-',
	".-.-.-": '.',
	"-..-.":  '/',
	"-----":  '0',
	".----":  '1',
	"..---":  '2',
	"...--":  '3',
	"....-":  '4',
	".....":  '5',
	"-....":  '6',
	"--...":  '7',
	"---..":  '8',
	"----.":  '9',
	"---...": ':',
	"..--..": '?',
	".--.-.": '@',
	".-":     'A',
	"-...":   'B',
	"-.-.":   'C',
	"-..":    'D',
	".":      'E',
	"..-.":   'F',
	"--.":    'G',
	"....":   'H',
	"..":     'I',
	".---":   'J',
	"-.-":    'K',
	".-..":   'L',
	"--":     'M',
	"-.":     'N',
	"---":    'O',
	".--.":   'P',
	"--.-":   'Q',
	".-.":    'R',
	"...":    'S',
	"-":      'T',
	"..-":    'U',
	"...-":   'V',
	".--":    'W',
	"-..-":   'X',
	"-.--":   'Y',
	"--..":   'Z',
}

func DecodeBits(bits string) string {

	// '0' - signal off
	// '1' - signal on

	trimBits := strings.Trim(bits, "0")
	bs := []byte(trimBits)

	var (
		m0 = make(map[int]int)
		m1 = make(map[int]int)
	)

	walkSolid(bs,
		func(value byte, length int) bool {
			switch value {
			case '0':
				m0[length]++
			case '1':
				m1[length]++
			default:
			}
			return true
		})

	unit0, _ := maxValueKey(m0)
	unit1, _ := maxValueKey(m1)

	var unit int

	if unit0 == 0 {
		unit = unit1
	} else {
		if unit0 == unit1 {
			unit = unit0
		} else {
			unit = minInt(unit0, unit1)
		}
	}
	if unit == 0 {
		unit = 1
	}

	var b strings.Builder

	walkSolid(bs,
		func(value byte, length int) bool {
			k := length / unit
			switch value {
			case '0':
				{
					switch k {
					case 1:
						// skip
					case 3:
						b.WriteString(" ") // char delimiter
					case 7:
						b.WriteString("   ") // word delimiter
					default:
						log.Fatalf("invalid %q length %d", '0', k)
					}
				}
			case '1':
				{
					switch k {
					case 1:
						b.WriteByte('.')
					case 3:
						b.WriteByte('-')
					default:
						log.Fatalf("invalid %q length %d", '1', k)
					}
				}
			default:
				log.Fatalf("invalid value %q", value)
			}

			return true
		})

	return b.String()
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func walkSolid(bs []byte, f func(value byte, length int) bool) {

	var value byte
	var length int = 0 // there is not value

	for _, b := range bs {
		if (length > 0) && (b == value) {
			length++
		} else {
			if length > 0 {
				if !f(value, length) {
					return
				}
			}
			value = b
			length = 1
		}
	}
	if length > 0 {
		if !f(value, length) {
			return
		}
	}
}

func maxValueKey(m map[int]int) (int, bool) {
	var hasKey bool
	var resKey int
	for key, val := range m {
		if hasKey {
			if m[resKey] < val {
				resKey = key
			}
		} else {
			resKey = key
			hasKey = true
		}
	}
	return resKey, hasKey
}

func DecodeMorse(morseCode string) string {
	const (
		wordDelimiter = "   " // 3 spaces
		charDelimiter = " "   // 1 space
	)
	morseCode = strings.TrimSpace(morseCode)
	var b strings.Builder
	words := strings.Split(morseCode, wordDelimiter)
	for i, word := range words {
		if i > 0 {
			b.WriteByte(' ')
		}
		chars := strings.Split(word, charDelimiter)
		for _, char := range chars {
			r, ok := morseDecodeMap[char]
			if ok {
				b.WriteRune(r)
			}
		}
	}
	return b.String()
}
