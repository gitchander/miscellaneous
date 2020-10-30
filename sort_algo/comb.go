package sort_algo

import (
	"math"
	"sort"
)

// Comb sort
// Сортировка расчёской
func CombSort(data sort.Interface) {

	// factor := 1 / (1 - math.Exp(-math.Phi))
	const factor = 1.2473309501039787

	n := data.Len()
	dist := float64(n)

	for {
		dist /= factor

		i := 0
		j := int(math.Floor(dist))

		if j <= 1 {
			break
		}

		for j < n {
			if data.Less(j, i) {
				data.Swap(j, i)
			}
			i++
			j++
		}
	}

	BubbleSort(data)
}
