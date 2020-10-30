package sort_algo

import "sort"

// Insertion sort
// Сортировка вставками
func InsertionSort(data sort.Interface) {
	coreInsertion(data, 1)
}

//func InsertionSort(data sort.Interface) {
//	n := data.Len()
//	for i := 1; i < n; i++ {
//		for j := i; (j > 0) && data.Less(j, j-1); j-- {
//			data.Swap(j, j-1)
//		}
//	}
//}

func coreInsertion(data sort.Interface, step int) {
	n := data.Len()
	for i := step; i < n; i++ {
		for j := i; (j >= step) && data.Less(j, j-step); j -= step {
			data.Swap(j, j-step)
		}
	}
}
