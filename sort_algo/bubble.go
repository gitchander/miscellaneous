package sort_algo

import "sort"

// Bubble sort
// Сортировка пузырьком
func BubbleSort(data sort.Interface) {
	n := data.Len()
	for {
		n--
		swapped := false
		for i := 0; i < n; i++ {
			if data.Less(i+1, i) {
				data.Swap(i+1, i)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
