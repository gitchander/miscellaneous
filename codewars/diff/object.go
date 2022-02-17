package main

import (
	"fmt"
)

type Object interface {
	isObject()
}

func (Constant) isObject() {}
func (Variable) isObject() {}
func (Operator) isObject() {}
func (Function) isObject() {}

type Constant float64

func ct(c float64) Constant {
	return Constant(c)
}

type Variable string

// (op arg1 arg2)
type Operator struct {
	Sign string // {'+', '-', '*', '/', '^'}
	Arg1 Object
	Arg2 Object
}

func makeOperator(sign string, arg1, arg2 Object) Operator {
	return Operator{
		Sign: sign,
		Arg1: arg1,
		Arg2: arg2,
	}
}

func add(a1, a2 Object) Object {
	return makeOperator("+", a1, a2)
}

func sub(a1, a2 Object) Object {
	return makeOperator("-", a1, a2)
}

func mul(a1, a2 Object) Object {
	return makeOperator("*", a1, a2)
}

func div(a1, a2 Object) Object {
	return makeOperator("/", a1, a2)
}

func pow(a1, a2 Object) Object {
	return makeOperator("^", a1, a2)
}

// (func arg)
type Function struct {
	Name string
	Arg  Object
}

func fn(name string, arg Object) Function {
	return Function{
		Name: name,
		Arg:  arg,
	}
}

func formatObject(obj Object) string {
	switch t := obj.(type) {
	case Constant:
		return fmt.Sprint(t)
	case Variable:
		return string(t)
	case Operator:
		return fmt.Sprintf("(%s %s %s)", t.Sign,
			formatObject(t.Arg1),
			formatObject(t.Arg2),
		)
	case Function:
		return fmt.Sprintf("(%s %s)", t.Name, formatObject(t.Arg))
	}
	return ""
}
