package main

import (
	"fmt"
	"os"
)

// t=0: v0
// t=1: v1
// t=0.5: (v0+v1) / 2
func lerp(v0, v1 float64, t float64) float64 {
	return v0*(1-t) + v1*t
}

func MkdirIfNotExist(name string) error {
	fi, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(name, os.ModePerm)
		}
		return err
	}
	if !(fi.IsDir()) {
		return fmt.Errorf("name %q is not a directory", name)
	}
	return nil
}

func mod(x, y int) int {
	m := x % y
	if m < 0 {
		m += y
	}
	return m
}

// func cloneInts(a []int) []int {
// 	b := make([]int, len(a))
// 	copy(b, a)
// 	return b
// }

func cloneSlice[T any](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	return b
}
