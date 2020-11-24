package main

import (
	"fmt"
	"log"

	rat "github.com/gitchander/miscellaneous/rational"
)

func main() {
	set()
	mul()
	parseTest()
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
