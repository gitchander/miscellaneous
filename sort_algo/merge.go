package sort_algo

import (
	"sort"

	. "github.com/gitchander/miscellaneous/sort_algo/util"
)

// data: [0 ... n)
// temp: [n ... 2n)
func MergeSort(data sort.Interface, n int) {

	if data.Len() < 2*n {
		panic("mergesort: insufficient data length.")
	}

	var vs = [2]Interval{
		Interval{Min: 0, Max: n},     // data
		Interval{Min: n, Max: 2 * n}, // temp
	}

	k := merge_sort(data, vs)
	if k == 1 {
		for i := 0; i < n; i++ {
			data.Swap(i, n+i)
		}
	}
}

func merge_sort(data sort.Interface, vs [2]Interval) int {

	if vs[0].Len() < 2 {
		return 0
	}

	var (
		st0 = bisect(vs[0])
		st1 = bisect(vs[1])
	)

	var (
		left  = [2]Interval{st0[0], st1[0]}
		right = [2]Interval{st0[1], st1[1]}
	)

	var (
		i = merge_sort(data, left)
		j = merge_sort(data, right)
	)

	// Calc result index
	var k int
	if i == 0 {
		k = 1
	}
	merge(data, left[i], right[j], vs[k])

	return k
}

// Bisect the interval.
func bisect(a Interval) []Interval {
	middle := (a.Min + a.Max) / 2
	return []Interval{
		Interval{Min: a.Min, Max: middle},
		Interval{Min: middle, Max: a.Max},
	}
}

// Merge left and right to result
func merge(data sort.Interface, left, right, result Interval) {

	var (
		i = left.Min
		j = right.Min
		k = result.Min
	)

	for (i < left.Max) && (j < right.Max) {
		if data.Less(i, j) {
			// left
			data.Swap(i, k)
			i++
			k++
		} else {
			// right
			data.Swap(j, k)
			j++
			k++
		}
	}

	for i < left.Max {
		data.Swap(i, k)
		i++
		k++
	}

	for j < right.Max {
		data.Swap(j, k)
		j++
		k++
	}
}
