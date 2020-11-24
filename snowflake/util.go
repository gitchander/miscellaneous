package main

import (
	"errors"
	"fmt"
	"image"
	"strconv"
	"strings"
)

var errInvalidSize = errors.New("invalid size format")

func parseSize(s string) (image.Point, error) {
	var zeroPoint image.Point
	vs := strings.Split(s, "x")
	if len(vs) != 2 {
		return zeroPoint, errInvalidSize
	}
	x, err := strconv.Atoi(vs[0])
	if err != nil {
		return zeroPoint, fmt.Errorf("invalid size format: %s", err)
	}
	y, err := strconv.Atoi(vs[1])
	if err != nil {
		return zeroPoint, fmt.Errorf("invalid size format: %s", err)
	}
	p := image.Point{
		X: x,
		Y: y,
	}
	return p, nil
}
