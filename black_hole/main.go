package main

import (
	"fmt"
	"io"
)

type BlackHole struct {
	mass int
}

var _ io.Writer = &BlackHole{}

func (p *BlackHole) Write(data []byte) (n int, err error) {
	n = len(data)
	p.mass += n
	return n, nil
}

func (p BlackHole) Mass() int {
	return p.mass
}

func Merge(a, b BlackHole) BlackHole {
	return BlackHole{
		mass: a.mass + b.mass,
	}
}

func main() {
	var b BlackHole
	data := []byte("Hello, Black hole!")
	b.Write(data)
	fmt.Println("mass of black hole equal", b.Mass())
}
