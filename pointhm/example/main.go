package main

import (
	"fmt"

	"github.com/gitchander/miscellaneous/pointhm"
)

func main() {

	p := pointhm.NewPointHM()

	p.Set(2, -5, struct{}{})
	p.Set(1, 7, struct{}{})
	p.Set(5, 3, struct{}{})
	p.Set(7, -3, struct{}{})

	v, ok := p.Get(2, -5)
	fmt.Println("get:", v, ok)

	v, ok = p.Remove(1, 7)
	fmt.Println("remove:", v, ok)

	p.Walk(func(x, y int, v interface{}) {
		fmt.Printf("(%d,%d)\n", x, y)
	})
}
