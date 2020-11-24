package morse

import (
	"unicode"
)

func MakeEncodeMap() map[string]string {
	m := make(map[string]string)
	for r, ds := range charMapEnglish {
		var (
			key = string(unicode.ToUpper(r))
			val = dashDotsToString(ds)
		)
		m[key] = val
	}
	return m
}

func MakeDecodeMap() map[string]string {
	m := make(map[string]string)
	for r, ds := range charMapEnglish {
		var (
			key = dashDotsToString(ds)
			val = string(unicode.ToUpper(r))
		)
		m[key] = val
	}
	return m
}

func dashDotsToString(ds []int) string {
	vs := make([]byte, len(ds))
	for i, d := range ds {
		switch d {
		case dot:
			vs[i] = '.'
		case dash:
			vs[i] = '-'
		default:
			panic("invalid char map")
		}
	}
	return string(vs)
}
