package rational

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errInvalidFormat = errors.New("invalid format")
	errNotDotFormat  = errors.New("string isn't \"dot\" format")
	errNotDivFormat  = errors.New("string isn't \"div\" format")
)

func Parse(s string) (Rational, error) {

	if x, err := strconv.Atoi(s); err == nil {
		return RatInt(x), nil
	}

	if r, err := parseDiv(s); err == nil {
		return r, nil
	}

	if r, err := parseDot(s); err == nil {
		return r, nil
	}

	return zero, errInvalidFormat
}

// Variant with div: 1/2, 321/654
func parseDiv(s string) (Rational, error) {

	index := strings.IndexByte(s, '/')
	if index == -1 {
		return zero, errNotDivFormat
	}

	var (
		sP = s[:index]
		sQ = s[index+1:]
	)

	p, err := strconv.Atoi(sP)
	if err != nil {
		return zero, err
	}

	q, err := strconv.Atoi(sQ)
	if err != nil {
		return zero, err
	}

	return Rat(p, q), nil
}

// Variant with dot: 0.1, 3.14
func parseDot(s string) (Rational, error) {

	index := strings.IndexByte(s, '.')
	if index == -1 {
		return zero, errNotDotFormat
	}

	var (
		sA = s[:index]   // before dot
		sB = s[index+1:] // after dot
	)

	a, err := strconv.Atoi(sA)
	if err != nil {
		return zero, err
	}

	b, err := strconv.Atoi(sB)
	if err != nil {
		return zero, err
	}

	exp := powerInt(10, len(sB))

	var (
		p = a*exp + b
		q = exp
	)

	return Rat(p, q), nil
}
