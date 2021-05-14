package main

import (
	"fmt"
	"log"

	rat "github.com/gitchander/miscellaneous/rational"
)

func main() {
	// sample1()
	// set()
	// mul()
	parseTest()
}

func sample1() {
	a := rat.Rat(2, 4)
	b := rat.Rat(3, 6)
	c := a.Add(b)
	fmt.Println(c)
}

func set() {
	a := rat.Rat(1, -5)
	fmt.Println(a)
}

func mul() {
	var (
		a = rat.Rat(5, 21)
		b = rat.Rat(7, -4)
	)

	c := a.Mul(b)

	fmt.Println(c)
}

func parseTest() {
	s := "3.14159265"
	r, err := rat.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}
