package matrix

// power returns a^n
func power(a int, n int) int {
	b := 1
	for i := 0; i < n; i++ {
		b *= a
	}
	return b
}

// fastPower returns a^n
func fastPower(a int, n int) int {
	b := 1
	for n > 0 {
		if (n & 1) == 1 { // n is odd
			b *= a // b = b * a
		}
		n >>= 1 // n = n / 2
		a *= a  // a = a ^ 2
	}
	return b
}
