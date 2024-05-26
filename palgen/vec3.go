package palgen

import "math"

type Vec3 [3]float64

func (a Vec3) Add(b Vec3) (c Vec3) {
	vecAdd(a[:], b[:], c[:], 3)
	return
}

func (a Vec3) Sub(b Vec3) (c Vec3) {
	vecSub(a[:], b[:], c[:], 3)
	return
}

func (a Vec3) MulScalar(scalar float64) (b Vec3) {
	vecMulScalar(a[:], b[:], scalar, 3)
	return
}

func (a Vec3) DivScalar(scalar float64) (b Vec3) {
	vecDivScalar(a[:], b[:], scalar, 3)
	return
}

func (a Vec3) Mul(b Vec3) (c Vec3) {
	vecMul(a[:], b[:], c[:], 3)
	return
}

func CosVec3(a Vec3) (b Vec3) {
	vecCos(a[:], b[:], 3)
	return
}

// z = x + y
func vecAdd(x, y, z []float64, n int) {
	for i := 0; i < n; i++ {
		z[i] = x[i] + y[i]
	}
}

// z = x - y
func vecSub(x, y, z []float64, n int) {
	for i := 0; i < n; i++ {
		z[i] = x[i] - y[i]
	}
}

// y = x * t
func vecMulScalar(x, y []float64, t float64, n int) {
	for i := 0; i < n; i++ {
		y[i] = x[i] * t
	}
}

// y = x / t
func vecDivScalar(x, y []float64, t float64, n int) {
	for i := 0; i < n; i++ {
		y[i] = x[i] / t
	}
}

// z = x * y
// It represents a component-wise product of two vectors
func vecMul(x, y, z []float64, n int) {
	for i := 0; i < n; i++ {
		z[i] = x[i] * y[i]
	}
}

// y = cos(x)
func vecCos(x, y []float64, n int) {
	for i := 0; i < n; i++ {
		y[i] = math.Cos(x[i])
	}
}
