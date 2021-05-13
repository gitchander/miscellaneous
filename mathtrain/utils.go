package main

import (
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func lerp(v0, v1 float64, t float64) float64 {
	return v0*(1-t) + v1*t
}

func percent(part, full float64) float64 {
	return part * 100 / full
}
