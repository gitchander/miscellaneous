package sort_algo

import "sort"

// Quick sort
// Быстрая сортировка
func QuickSort(data sort.Interface) {
	quicksort(data, 0, data.Len())
}

func quicksort(data sort.Interface, lo, hi int) {
	if lo < hi {
		p := partition(data, lo, hi)
		quicksort(data, lo, p)
		quicksort(data, p+1, hi)
	}
}

var (
	partition = partitionFirst
	//partition = partitionLast
)

func partitionFirst(data sort.Interface, lo, hi int) int {

	pivot := lo

	i := lo + 1
	j := hi

	for {
		for (i < j) && !data.Less(pivot, i) { // [p] >= [i]
			i++
		}

		for (i < j) && data.Less(pivot, j-1) { // [p] < [j]
			j--
		}

		if i >= j {
			break
		}

		data.Swap(i, j-1)

		i++
		j--
	}

	data.Swap(pivot, i-1)

	return i - 1
}

func partitionLast(data sort.Interface, lo, hi int) int {

	pivot := hi - 1

	i := lo
	j := hi - 1

	for {
		for (i < j) && !data.Less(pivot, i) { // [p] >= [i]
			i++
		}

		for (i < j) && data.Less(pivot, j-1) { // [p] < [j]
			j--
		}

		if i >= j {
			break
		}

		data.Swap(i, j-1)

		i++
		j--
	}

	data.Swap(pivot, j)

	return j
}
