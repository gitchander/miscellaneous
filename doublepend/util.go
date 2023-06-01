package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

const Tau = 2 * math.Pi

func DegToRad(deg float64) float64 {
	return deg * Tau / 360
}

func RadToDeg(rad float64) float64 {
	return rad * 360 / Tau
}

// func DegToRad(deg float64) float64 {
// 	return deg * math.Pi / 180
// }

// func RadToDeg(rad float64) float64 {
// 	return rad * 180 / math.Pi
// }

func clampFloat64(a float64, min, max float64) float64 {
	if max < min { // empty range
		return 0
	}
	if a < min {
		a = min
	}
	if a > max {
		a = max
	}
	return a
}

func ceilPowerOfTwo(x int) int {
	d := 1
	for d < x {
		d *= 2
	}
	return d
}

var (
	//angleNormalize = angleNormalizeV1
	angleNormalize = angleNormalizeV2
)

// [-Pi, +Pi)
func angleNormalizeV1(angle float64) float64 {
	for angle < -math.Pi {
		angle += Tau
	}
	for angle >= math.Pi {
		angle -= Tau
	}
	return angle
}

// a % b
func modFloat64(a, b float64) float64 {
	m := a - math.Floor(a/b)*b
	if m < 0 {
		m += b
	}
	return m
}

func angleNormalizeV2(angle float64) float64 {

	angle += math.Pi

	// [0, 2*Pi) = [0, Tau)
	angle = modFloat64(angle, Tau)
	//angle = math.Mod(angle, Tau)

	angle -= math.Pi

	return angle
}

// Lerp - Linear interpolation
// t= [0..1]
// (t == 0) => v0
// (t == 1) => v1
func lerp(v0, v1 float64, t float64) float64 {
	return (1.0-t)*v0 + t*v1
}

func mod(x, y int) int {
	m := x % y
	if m < 0 {
		m += y
	}
	return m
}

func checkMax(prefix string, x, max float64) {
	if math.Abs(x) > max {
		err := fmt.Errorf("%s: %v", prefix, x)
		panic(err)
	}
}

const (
	minVelocity = -1e+20
	maxVelocity = +1e+20
)

func clampVelocity(velocity float64) float64 {
	return clampFloat64(velocity, minVelocity, maxVelocity)
}

// ------------------------------------------------------------------------------
var (
	//strCat = strCat1
	strCat = strCat2
)

func strCat1(vs ...string) string {
	var b strings.Builder
	for _, v := range vs {
		b.WriteString(v)
	}
	return b.String()
}

func strCat2(vs ...string) string {
	return strings.Join(vs, "")
}

// ------------------------------------------------------------------------------
var (
	// time.RFC3339 - "2006-01-02T15:04:05"

	dateTimeFormat = strCat(
		"2006", "01", "02", // date (YYYYMMDD)
		"15", "04", "05", // time (HHMMSS)
	)
)

func makeFilename(prefix, extention string) string {
	st := time.Now().Format(dateTimeFormat)
	return prefix + st + extention
}
