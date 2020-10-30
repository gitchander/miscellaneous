package sort_algo

import "sort"

// Stooge sort
func StoogeSort(data sort.Interface) {
	stooge_sort(data, 0, data.Len())
}

func stooge_sort(data sort.Interface, i, j int) {
	count := j - i
	if count > 1 {
		if data.Less(j-1, i) {
			data.Swap(j-1, i)
		}
		if count > 2 {
			t := count / 3
			stooge_sort(data, i, j-t)
			stooge_sort(data, i+t, j)
			stooge_sort(data, i, j-t)
		}
	}
}
