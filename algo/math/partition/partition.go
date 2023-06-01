package partition

// https://en.wikipedia.org/wiki/Partition_(number_theory)

// 1, 1, 2, 3, 5, 7, 11, 15, 22, 30, 42, 56, 77, 101, 135, 176, 231, 297, 385,
// 490, 627, 792, 1002, 1255, 1575, 1958, 2436, 3010, 3718, 4565, 5604, ...

// ------------------------------------------------------------------------------
func Partition(n int) int {
	return part(n, n)
}

func part(n, k int) int {
	if n == 0 {
		return 1
	}
	if (k <= 0) || (n < 0) {
		return 0
	}
	return part(n-k, k) + part(n, k-1)
}

// ------------------------------------------------------------------------------
func PartitionWalk(n int, f func([]int)) int {
	if n < 0 {
		return 0
	}
	if f == nil {
		f = func([]int) {}
	}
	as := make([]int, 0, n)
	return partWalk(n, n, as, f)
}

// func partWalk(n, k int, as []int, f func([]int)) int {
// 	if n == 0 {
// 		f(as)
// 		return 1
// 	}
// 	if (k <= 0) || (n < 0) {
// 		return 0
// 	}
// 	if k > n {
// 		k = n
// 	}
// 	return partWalk(n-k, k, append(as, k), f) + partWalk(n, k-1, as, f)
// }

func partWalk(n, k int, as []int, f func([]int)) int {

	if n == 0 {
		f(as)
		return 1
	}

	var count int

	if k > n {
		k = n
	}

	count += partWalk(n-k, k, append(as, k), f)

	if k > 1 {
		count += partWalk(n, k-1, as, f)
	}

	return count
}

// ------------------------------------------------------------------------------
func PartitionTest(n int, f func([]int)) int {
	if f == nil {
		f = func([]int) {}
	}
	as := make([]int, 0, n)
	return partMy(n, n, as, f)
}

func partMy(n, k int, as []int, f func([]int)) int {
	// if n < 0 {
	// 	return 0
	// }
	if n == 0 {
		f(as)
		return 1
	}
	if k > n {
		k = n
	}
	var count int
	for k > 0 {
		count += partMy(n-k, k, append(as, k), f)
		k--
	}
	return count
}
