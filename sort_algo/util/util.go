package util

import "sort"

func Scramble(r Randomer, data sort.Interface) {
	ScrambleN(r, data, data.Len())
}

func ScrambleN(r Randomer, data sort.Interface, n int) {
	for ; n > 1; n-- {
		data.Swap(r.Intn(n), n-1)
	}
}

func Reverse(data sort.Interface) {
	i, j := 0, data.Len()-1
	for i < j {
		data.Swap(i, j)
		i, j = i+1, j-1
	}
}
