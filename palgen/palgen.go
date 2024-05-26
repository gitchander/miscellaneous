package palgen

import (
	"image/color"
	"math"
)

const tau = 2 * math.Pi

// http://www.iquilezles.org/www/articles/palettes/palettes.htm

// t = [0..1]
func ColorByParams(p Params, t float64) color.Color {
	v := palette(p.A, p.B, p.C, p.D, t)
	c := RGB{
		R: clamp01(v[0]),
		G: clamp01(v[1]),
		B: clamp01(v[2]),
	}
	return c.Color()
}

func palette(a, b, c, d Vec3, t float64) Vec3 {

	// a + b * cos( 2*pi * (c*t + d) )

	angle := c.MulScalar(t).Add(d).MulScalar(tau)

	cos := CosVec3(angle)

	return a.Add(b.Mul(cos))
}

func clamp01(x float64) float64 {
	return clampFloat64(x, 0, 1)
}

func clampFloat64(x float64, min, max float64) float64 {
	if max < min {
		panic("invalid interval")
	}
	if x < min {
		x = min
	}
	if x > max {
		x = max
	}
	return x
}
