package rational

// x ^ y
func powerInt(x, y int) int {
	exp := 1
	for i := 0; i < y; i++ {
		exp *= x
	}
	return exp
}
