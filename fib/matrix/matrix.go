package matrix

type Matrix [2][2]int

func (a Matrix) Mul(b Matrix) (c Matrix) {
	c[0][0] = a[0][0]*b[0][0] + a[1][0]*b[0][1]
	c[1][0] = a[0][0]*b[1][0] + a[1][0]*b[1][1]
	c[0][1] = a[0][1]*b[0][0] + a[1][1]*b[0][1]
	c[1][1] = a[0][1]*b[1][0] + a[1][1]*b[1][1]
	return
}

func MatrixIdentity() Matrix {
	return Matrix{
		{1, 0},
		{0, 1},
	}
}

// Power returns a^n
func (a Matrix) Power(n int) Matrix {
	b := MatrixIdentity()
	for i := 0; i < n; i++ {
		b = b.Mul(a)
	}
	return b
}

func (a Matrix) FastPower(n int) Matrix {
	b := MatrixIdentity()
	for n > 0 {
		if (n & 1) == 1 { // n is odd
			b = b.Mul(a) // b = b * a
		}
		n >>= 1      // n = n / 2
		a = a.Mul(a) // a = a ^ 2
	}
	return b
}

// fibN returns n-th number of Fibonacci Series
func FibN(n int) int {

	a := Matrix{
		{0, 1},
		{1, 1},
	}

	//b := a.Power(n)
	b := a.FastPower(n)

	return b[0][1]
}
