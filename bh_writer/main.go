package main

import (
	"fmt"
	"io"
)

// BlackHole writer
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

func Merge(bs ...BlackHole) BlackHole {
	var mass int
	for _, b := range bs {
		mass += b.mass
	}
	return BlackHole{
		mass: mass,
	}
}

func main() {
	var b BlackHole
	data := []byte("Hello, black hole writer!")
	b.Write(data)
	fmt.Println("mass of black hole equal", b.Mass())
}
