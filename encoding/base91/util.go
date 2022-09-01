package base91

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}

func ceilDiv(a, b int) int {
	return (a + (b - 1)) / b
}
