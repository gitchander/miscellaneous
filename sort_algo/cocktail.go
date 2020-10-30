package sort_algo

import "sort"

// Cocktail sort
// Shaker sort
// Сортировка перемешиванием
func CocktailSort(data sort.Interface) {
	var (
		left  = 0
		right = data.Len() - 1
	)
	for left < right {
		for i := left; i < right; i++ {
			if data.Less(i+1, i) {
				data.Swap(i+1, i)
			}
		}
		right--
		for i := right; i > left; i-- {
			if data.Less(i, i-1) {
				data.Swap(i, i-1)
			}
		}
		left++
	}
}

// Cocktail sort with swapping
//func CocktailSort(data sort.Interface) {
//	var (
//		left  = 0
//		right = data.Len() - 1
//	)
//	for left < right {
//		swapped := false
//		for i := left; i < right; i++ {
//			if data.Less(i+1, i) {
//				data.Swap(i+1, i)
//				swapped = true
//			}
//		}
//		if !swapped {
//			break
//		}
//		right--
//		swapped = false
//		for i := right; i > left; i-- {
//			if data.Less(i, i-1) {
//				data.Swap(i, i-1)
//				swapped = true
//			}
//		}
//		if !swapped {
//			break
//		}
//		left++
//	}
//}
