package main

import (
	"fmt"
	"strings"
)

// https://en.wikipedia.org/wiki/Operation_(mathematics)

// Unary operation
// https://en.wikipedia.org/wiki/Unary_operation

type UnaryOperation interface {
	Do(x int) int
	Symbol() string
}

// Binary operation
// https://en.wikipedia.org/wiki/Binary_operation

type BinaryOperation interface {
	Do(x, y int) int
	Symbol() string // symbol or sign
}

//------------------------------------------------------------------------------
// Addition
// https://en.wikipedia.org/wiki/Addition

type Add struct{}

func (Add) Do(x, y int) int {
	return x + y
}

func (Add) Symbol() string {
	return "+"
}

//------------------------------------------------------------------------------
// Subtraction
// https://en.wikipedia.org/wiki/Subtraction

type Sub struct{}

func (Sub) Do(x, y int) int {
	return x - y
}

func (Sub) Symbol() string {
	return "-"
}

//------------------------------------------------------------------------------
// Multiplication
// https://en.wikipedia.org/wiki/Multiplication

type Mul struct{}

func (Mul) Do(x, y int) int {
	return x * y
}

func (Mul) Symbol() string {
	return "*"
}

//------------------------------------------------------------------------------
// Division
// https://en.wikipedia.org/wiki/Division_(mathematics)

type Div struct{}

func (Div) Do(x, y int) int {
	return x / y
}

func (Div) Symbol() string {
	return "/"
}

// ------------------------------------------------------------------------------
func parseOperations(s string) ([]BinaryOperation, error) {
	var bos []BinaryOperation
	m := make(map[string]struct{})
	vs := strings.Split(s, ",")
	for _, v := range vs {
		bo, err := makeBinaryOperation(v)
		if err != nil {
			return nil, err
		}
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			bos = append(bos, bo)
		}
	}
	return bos, nil
}

func makeBinaryOperation(name string) (BinaryOperation, error) {
	switch name {
	case "add":
		return Add{}, nil
	case "sub":
		return Sub{}, nil
	case "mul":
		return Mul{}, nil
	case "div":
		return Div{}, nil
	default:
		return nil, fmt.Errorf("invalid operation %q", name)
	}
}
