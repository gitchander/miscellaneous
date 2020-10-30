package sort_algo

import "sort"

// Selection sort
// Сортировка выбором
func SelectionSort(data sort.Interface) {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		if min != i {
			data.Swap(i, min)
		}
	}
}
