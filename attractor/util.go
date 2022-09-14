package attractor

func not(b bool) bool {
	return !b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func clamp(a float64, min, max float64) float64 {
	if a < min {
		a = min
	}
	if a > max {
		a = max
	}
	return a
}
