package cbig

type nat []Word

func (z nat) clear() {
	for i := range z {
		z[i] = 0
	}
}

func (z nat) norm() nat {
	i := len(z)
	for i > 0 && z[i-1] == 0 {
		i--
	}
	return z[0:i]
}

func (z nat) make(n int) nat {
	if n <= cap(z) {
		return z[:n] // reuse z
	}
	if n == 1 {
		// Most nats start small and stay that way; don't over-allocate.
		return make(nat, 1)
	}
	// Choosing a good value for e has significant performance impact
	// because it increases the chance that a value can be reused.
	const e = 4 // extra capacity
	return make(nat, n, n+e)
}

func (z nat) set(x nat) nat {
	z = z.make(len(x))
	copy(z, x)
	return z
}

func (z nat) setWord(x Word) nat {
	if x == 0 {
		return z[:0]
	}
	z = z.make(1)
	z[0] = x
	return z
}

func (z nat) add(x, y nat) nat {
	m := len(x)
	n := len(y)

	switch {
	case m < n:
		return z.add(y, x)
	case m == 0:
		// n == 0 because m >= n; result is 0
		return z[:0]
	case n == 0:
		// result is x
		return z.set(x)
	}
	// m > 0

	z = z.make(m + 1)
	c := addVV(z[0:n], x, y)
	if m > n {
		c = addVW(z[n:m], x[n:], c)
	}
	z[m] = c

	return z.norm()
}

func (z nat) sub(x, y nat) nat {
	m := len(x)
	n := len(y)

	switch {
	case m < n:
		panic("underflow")
	case m == 0:
		// n == 0 because m >= n; result is 0
		return z[:0]
	case n == 0:
		// result is x
		return z.set(x)
	}
	// m > 0

	z = z.make(m)
	c := subVV(z[0:n], x, y)
	if m > n {
		c = subVW(z[n:], x[n:], c)
	}
	if c != 0 {
		panic("underflow")
	}

	return z.norm()
}

// basicMul multiplies x and y and leaves the result in z.
// The (non-normalized) result is placed in z[0 : len(x) + len(y)].
func basicMul(z, x, y nat) {
	z[0 : len(x)+len(y)].clear() // initialize z
	for i, d := range y {
		if d != 0 {
			z[len(x)+i] = addMulVVW(z[i:i+len(x)], x, d)
		}
	}
}

// alias reports whether x and y share the same base array.
// Note: alias assumes that the capacity of underlying arrays
//
//	is never changed for nat values; i.e. that there are
//	no 3-operand slice expressions in this code (or worse,
//	reflect-based operations to the same effect).
func alias(x, y nat) bool {
	return cap(x) > 0 && cap(y) > 0 && &x[0:cap(x)][cap(x)-1] == &y[0:cap(y)][cap(y)-1]
}

func (z nat) mulAddWW(x nat, y, r Word) nat {
	m := len(x)
	if m == 0 || y == 0 {
		return z.setWord(r) // result is r
	}
	// m > 0

	z = z.make(m + 1)
	z[m] = mulAddVWW(z[0:m], x, y, r)

	return z.norm()
}

func (z nat) mul(x, y nat) nat {
	m := len(x)
	n := len(y)

	switch {
	case m < n:
		return z.mul(y, x)
	case m == 0 || n == 0:
		return z[:0]
	case n == 1:
		return z.mulAddWW(x, y[0], 0)
	}
	// m >= n > 1

	// determine if z can be reused
	if alias(z, x) || alias(z, y) {
		z = nil // z is an alias for x or y - cannot reuse
	}

	// use basic multiplication if the numbers are small

	z = z.make(m + n)
	basicMul(z, x, y)
	return z.norm()
}

func (x nat) cmp(y nat) (r int) {
	m := len(x)
	n := len(y)
	if m != n || m == 0 {
		switch {
		case m < n:
			r = -1
		case m > n:
			r = 1
		}
		return
	}

	i := m - 1
	for i > 0 && x[i] == y[i] {
		i--
	}

	switch {
	case x[i] < y[i]:
		r = -1
	case x[i] > y[i]:
		r = 1
	}
	return
}

// // divBasic performs word-by-word division of u by v.
// // The quotient is written in pre-allocated q.
// // The remainder overwrites input u.
// //
// // Precondition:
// // - q is large enough to hold the quotient u / v
// //   which has a maximum length of len(u)-len(v)+1.
// func (q nat) divBasic(u, v nat) {
// 	n := len(v)
// 	m := len(u) - n

// 	qhatvp := getNat(n + 1)
// 	qhatv := *qhatvp

// 	// D2.
// 	vn1 := v[n-1]
// 	for j := m; j >= 0; j-- {
// 		// D3.
// 		qhat := Word(_M)
// 		var ujn Word
// 		if j+n < len(u) {
// 			ujn = u[j+n]
// 		}
// 		if ujn != vn1 {
// 			var rhat Word
// 			qhat, rhat = divWW(ujn, u[j+n-1], vn1)

// 			// x1 | x2 = q̂v_{n-2}
// 			vn2 := v[n-2]
// 			x1, x2 := mulWW(qhat, vn2)
// 			// test if q̂v_{n-2} > br̂ + u_{j+n-2}
// 			ujn2 := u[j+n-2]
// 			for greaterThan(x1, x2, rhat, ujn2) {
// 				qhat--
// 				prevRhat := rhat
// 				rhat += vn1
// 				// v[n-1] >= 0, so this tests for overflow.
// 				if rhat < prevRhat {
// 					break
// 				}
// 				x1, x2 = mulWW(qhat, vn2)
// 			}
// 		}

// 		// D4.
// 		// Compute the remainder u - (q̂*v) << (_W*j).
// 		// The subtraction may overflow if q̂ estimate was off by one.
// 		qhatv[n] = mulAddVWW(qhatv[0:n], v, qhat, 0)
// 		qhl := len(qhatv)
// 		if j+qhl > len(u) && qhatv[n] == 0 {
// 			qhl--
// 		}
// 		c := subVV(u[j:j+qhl], u[j:], qhatv)
// 		if c != 0 {
// 			c := addVV(u[j:j+n], u[j:], v)
// 			// If n == qhl, the carry from subVV and the carry from addVV
// 			// cancel out and don't affect u[j+n].
// 			if n < qhl {
// 				u[j+n] += c
// 			}
// 			qhat--
// 		}

// 		if j == m && m == len(q) && qhat == 0 {
// 			continue
// 		}
// 		q[j] = qhat
// 	}

// 	putNat(qhatvp)
// }
