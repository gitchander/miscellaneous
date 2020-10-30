package sort_algo

import "sort"

// Gnome sort
// Гномья сортировка
func GnomeSort(data sort.Interface) {
	n := data.Len()
	for i := 1; i < n; {
		if (i == 0) || !data.Less(i, i-1) { // [i-1] <= [i]
			i++
		} else {
			data.Swap(i, i-1)
			i--
		}
	}
}
