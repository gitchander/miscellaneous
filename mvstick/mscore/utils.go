package mscore

import (
	"sort"
	"strconv"
)

func not(b bool) bool {
	return !b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// ------------------------------------------------------------------------------
type runeSlice []rune

var _ sort.Interface = runeSlice(nil)

func (p runeSlice) Len() int           { return len(p) }
func (p runeSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p runeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p runeSlice) Sort() { sort.Sort(p) }

//------------------------------------------------------------------------------

//const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Sign, Letter, Character, Label, Symbol

func indexToLetter(i int) (byte, bool) {
	if (0 <= i) && (i < 26) {
		return byte(i + 'A'), true
	}
	return 0, false
}

func letterToIndex(letter byte) (int, bool) {
	if ('A' <= letter) && (letter <= 'Z') {
		return int(letter - 'A'), true
	}
	return 0, false
}

//------------------------------------------------------------------------------

func cloneSlice[T any](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	return b
}

func reverseSlice[T any](a []T) {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i, j = i+1, j-1
	}
}
