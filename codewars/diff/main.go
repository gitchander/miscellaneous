package main

import (
	"fmt"
	"log"
)

// https://www.codewars.com/kata/584daf7215ac503d5a0001ae

func main() {
	testSamples()
}

func testSamples() {

	test("5", "0")
	test("x", "1")
	test("(+ x x)", "2")
	test("(- x x)", "0")
	test("(* x 2)", "2")
	test("(/ x 2)", "0.5")
	test("(^ x 2)", "(* 2 x)")
	test("(cos x)", "(* -1 (sin x))")
	test("(sin x)", "(cos x)")
	test("(tan x)", "(+ 1 (^ (tan x) 2))")
	test("(exp x)", "(exp x)")
	test("(ln x)", "(/ 1 x)")

	test("(+ x (+ x x))", "3")
	test("(- (+ x x) x)", "1")
	test("(- (+ x x) x)", "1")
	test("(/ 2 (+ 1 x))", "(/ -2 (^ (+ 1 x) 2))")
	test("(cos (+ x 1))", "(* -1 (sin (+ x 1)))")
	test("(cos (* 2 x))", "(* 2 (* -1 (sin (* 2 x))))")

	test("(sin (+ x 1))", "(cos (+ x 1))")
	test("(sin (* 2 x))", "(* 2 (cos (* 2 x)))")
	test("(tan (* 2 x))", "(* 2 (+ 1 (^ (tan (* 2 x)) 2)))")
	test("(exp (* 2 x))", "(* 2 (exp (* 2 x)))")

	testDD("(sin x)", "(* -1 (sin x))")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Diff(expression string) string {
	v, _, err := parseObject([]byte(expression))
	checkError(err)
	v = simplifyObject(v)
	v = derivativeObject(v)
	v = simplifyObject(v)
	return formatObject(v)
}

func test(a, b string) {
	s := Diff(a)
	if false {
		fmt.Println(s)
	}
	if s != b {
		log.Fatalf("invalid result: have %q, want %q", s, b)
	}
}

func testDD(a, b string) {
	s := Diff(Diff(a))
	if s != b {
		log.Fatalf("invalid result: have %q, want %q", s, b)
	}
}
