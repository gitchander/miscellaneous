package sort_algo

import "sort"

func ShellSort(data sort.Interface) {
	for k := data.Len() / 2; k > 0; k /= 2 {
		coreInsertion(data, k)
	}
}
