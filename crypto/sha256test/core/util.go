package core

func splitBytes(bs []byte, n int) [][]byte {

	samples := make([][]byte, ceilDiv(len(bs), n))

	i := 0
	for len(bs) >= n {
		samples[i] = bs[:n]
		i++
		bs = bs[n:]
	}

	if len(bs) > 0 {
		samples[i] = bs
	}

	return samples
}

func ceilDiv(a, b int) int {
	return (a + (b - 1)) / b
}
