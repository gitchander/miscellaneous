package main

//------------------------------------------------------------------------------
// Derivative
//------------------------------------------------------------------------------

// https://www.rapidtables.com/math/calculus/derivative.html

func derivativeObject(obj Object) Object {
	switch t := obj.(type) {
	case Constant:
		return ct(0)
	case Variable:
		return ct(1)
	case Operator:
		return derivativeOperator(t)
	case Function:
		return derivativeFunction(t)
	default:
		return nil
	}
}

func derivativeOperator(op Operator) Object {

	var (
		da1 = derivativeObject(op.Arg1)
		da2 = derivativeObject(op.Arg2)
	)

	switch op.Sign {
	case "+":
		return add(da1, da2)
	case "-":
		return sub(da1, da2)
	case "*":
		return add(mul(da1, op.Arg2), mul(op.Arg1, da2))
	case "/":
		{
			var (
				numerator   = sub(mul(da1, op.Arg2), mul(op.Arg1, da2))
				denominator = pow(op.Arg2, ct(2))
			)
			return div(numerator, denominator)
		}
	case "^":
		{
			if a, ok := op.Arg1.(Constant); ok {
				return mul(pow(a, op.Arg2), fn("ln", a))
			}
			if a, ok := op.Arg2.(Constant); ok {
				return mul(a, pow(op.Arg1, a-1))
			}
			return derivativePow(op.Arg1, op.Arg2, da1, da2)
		}
	}

	return op
}

func derivativeFunction(f Function) Object {
	var df Object
	switch f.Name {
	case "sin":
		df = fn("cos", f.Arg)
	case "cos":
		df = mul(ct(-1), fn("sin", f.Arg))
	case "tan":
		df = derivativeTan(f.Arg)
	case "ln":
		df = div(ct(1), f.Arg)
	case "exp":
		df = fn("exp", f.Arg)
	}
	da := derivativeObject(f.Arg)
	return mul(da, df)
}

func derivativeTan(obj Object) Object {
	//return div(ct(1), pow(fn("cos", obj), ct(2)))
	return add(ct(1), pow(fn("tan", obj), ct(2)))
}

// https://www.youtube.com/watch?v=SUxcFxM65Ho
// (f(x)^g(x))'
func derivativePow(f, g, df, dg Object) Object {
	var (
		v1 = pow(f, g)
		v2 = div(mul(g, df), f)
		v3 = mul(dg, fn("ln", f))
	)
	return mul(v1, add(v2, v3))
}
