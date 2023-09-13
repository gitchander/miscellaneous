package rational

// GCD - Greatest Common Denominator: largest number that can devide two numbers.
// GCD - Greatest Common Divisor
// https://en.wikipedia.org/wiki/Greatest_common_divisor

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func gcd_Euclidean(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func gcd_Remainder(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
