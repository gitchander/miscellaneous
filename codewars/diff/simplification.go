package main

import (
	"math"
)

//------------------------------------------------------------------------------
// Simplify
//------------------------------------------------------------------------------

type OptFloat64 struct {
	Present bool
	Value   float64
}

func objectToConstant(obj Object) OptFloat64 {
	v, ok := obj.(Constant)
	return OptFloat64{
		Present: ok,
		Value:   float64(v),
	}
}

// simplification
func simplifyObject(obj Object) Object {
	switch t := obj.(type) {
	case Constant:
		return t
	case Variable:
		return t
	case Operator:
		return simplifyOperator(t)
	case Function:
		t.Arg = simplifyObject(t.Arg)
		return t
	default:
		return obj
	}
}

func simplifyOperator(op Operator) Object {

	op.Arg1 = simplifyObject(op.Arg1)
	op.Arg2 = simplifyObject(op.Arg2)

	var (
		c1 = objectToConstant(op.Arg1)
		c2 = objectToConstant(op.Arg2)
	)

	if c1.Present && c2.Present {
		switch op.Sign {
		case "+":
			return Constant(c1.Value + c2.Value)
		case "-":
			return Constant(c1.Value - c2.Value)
		case "*":
			return Constant(c1.Value * c2.Value)
		case "/":
			return Constant(c1.Value / c2.Value)
		case "^":
			return Constant(math.Pow(c1.Value, c2.Value))
		}
	}

	switch op.Sign {
	case "+":
		if c1.Present && (c1.Value == 0) {
			return op.Arg2 // 0 + a2 = a2
		}
		if c2.Present && (c2.Value == 0) {
			return op.Arg1 // a1 + 0 = a1
		}

	case "*":
		if c1.Present {
			switch c1.Value {
			case 0:
				return Constant(0) // 0 * a2 = 0
			case 1:
				return op.Arg2 // 1 * a2 = a2
			}
		}
		if c2.Present {
			switch c2.Value {
			case 0:
				return Constant(0) // a1 * 0 = 0
			case 1:
				return op.Arg1 // a1 * 1 = a1
			}
		}
	case "^":
		if c2.Present {
			switch c2.Value {
			case 0:
				return Constant(1) // a1 ^ 0 = 1
			case 1:
				return op.Arg1 // a1 ^ 1 = a1
			}
		}
	}

	return op
}
