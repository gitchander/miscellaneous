package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseIntsN(s string, n int) ([]int, error) {
	fields := strings.Fields(s)
	if len(fields) < n {
		return nil, fmt.Errorf("number of fields %d is less than %d", len(fields), n)
	}
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		x, err := strconv.Atoi(fields[i])
		if err != nil {
			return nil, err
		}
		xs[i] = x
	}
	return xs, nil
}

func LineParseInt(lr *LineReader) (int, error) {
	line, err := lr.ReadLine()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(line)
}

// Read line and parse ints.
func LineParseIntsN(lr *LineReader, n int) ([]int, error) {
	line, err := lr.ReadLine()
	if err != nil {
		return nil, err
	}
	return ParseIntsN(line, n)
}
