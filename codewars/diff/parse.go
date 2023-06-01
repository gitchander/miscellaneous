package main

import (
	"errors"
	"fmt"
	"strconv"
)

//------------------------------------------------------------------------------
// Parse
//------------------------------------------------------------------------------

var parseError = errors.New("parse error")

func parseObject(data []byte) (obj Object, rest []byte, err error) {

	data = trimLeftSpaces(data)

	if len(data) == 0 {
		return nil, data, parseError
	}
	b := data[0]

	switch {
	case (b == '-') || byteIsDigit(b):
		return parseConstant(data)
	case byteIsLetter(b):
		return parseVariable(data)
	case (b == '('):
		return parseFuncOp(data)
	default:
		return nil, data, fmt.Errorf("invalid object byte (%#U)", b)
	}
}

func parseConstant(data []byte) (c Constant, rest []byte, err error) {

	var n int
	for (n < len(data)) && byteIsConstant(data[n]) {
		n++
	}

	x, err := strconv.ParseFloat(string(data[:n]), 64)
	if err != nil {
		return Constant(0), data, err
	}

	c = Constant(x)

	return c, data[n:], nil
}

func parseVariable(data []byte) (v Variable, rest []byte, err error) {

	var n int
	for (n < len(data)) && byteIsLetter(data[n]) {
		n++
	}

	v = Variable(string(data[:n]))

	return v, data[n:], nil
}

func parseFuncOp(data []byte) (obj Object, rest []byte, err error) {

	// read '('
	if !hasPrefixByte(data, '(') {
		return nil, data, parseError
	}
	data = data[1:]

	// read operator sign or function name
	i := indexFunc(data, byteIsSpace)
	if i == -1 {
		return nil, data, parseError
	}
	name := string(data[:i])
	data = data[i:]

	arg1, data, err := parseObject(data)
	if err != nil {
		return nil, data, err
	}

	if hasPrefixByte(data, ')') {
		f := Function{
			Name: name,
			Arg:  arg1,
		}
		return f, data[1:], nil
	}

	arg2, data, err := parseObject(data)
	if err != nil {
		return nil, data, err
	}

	data = trimLeftSpaces(data)
	if hasPrefixByte(data, ')') {
		o := Operator{
			Sign: name,
			Arg1: arg1,
			Arg2: arg2,
		}
		return o, data[1:], nil
	}

	return nil, data, parseError
}

// ------------------------------------------------------------------------------
type ByteIsFunc func(byte) bool

func indexFunc(data []byte, f ByteIsFunc) int {
	for i, b := range data {
		if f(b) {
			return i
		}
	}
	return -1
}

func hasPrefixByte(data []byte, b byte) bool {
	return (len(data) > 0) && (data[0] == b)
}

func trimLeft(data []byte, f ByteIsFunc) []byte {
	for (len(data) > 0) && f(data[0]) {
		data = data[1:]
	}
	return data
}

// skipSpaces
func trimLeftSpaces(data []byte) []byte {
	return trimLeft(data, byteIsSpace)
}

func byteIsSpace(b byte) bool {
	switch b {
	case ' ', '\t', '\n', '\v', '\f', '\r':
		return true
	default:
		return false
	}
}

func byteIsDigit(b byte) bool {
	return ('0' <= b) && (b <= '9')
}

func byteIsLetter(b byte) bool {
	return ('a' <= b) && (b <= 'z')
}

func byteIsConstant(b byte) bool {
	if byteIsDigit(b) {
		return true
	}
	switch b {
	case '+', '-', '.', 'e':
		return true
	default:
		return false
	}
}
