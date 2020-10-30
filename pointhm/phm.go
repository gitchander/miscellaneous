package pointhm

import (
	"container/list"
)

type PointHM struct {
	dx, dy int
	matrix [][]*list.List
}

func NewPointHM() *PointHM {

	dx := 5
	dy := 7

	matrix := make([][]*list.List, dx)
	for x := range matrix {
		ls := make([]*list.List, dy)
		for y := range ls {
			ls[y] = list.New()
		}
		matrix[x] = ls
	}

	return &PointHM{
		dx:     dx,
		dy:     dy,
		matrix: matrix,
	}
}

func (p *PointHM) getList(x, y int) *list.List {
	var (
		hX = mod(x, p.dx)
		hY = mod(y, p.dy)
	)
	return p.matrix[hX][hY]
}

func mod(x, y int) int {
	d := x % y
	if d < 0 {
		d += y
	}
	return d
}

func (p *PointHM) Set(x, y int, v interface{}) {

	l := p.getList(x, y)

	n := getNode(l, x, y)
	if n != nil {
		n.v = v
		return
	}

	n = &node{
		x: x,
		y: y,
		v: v,
	}

	l.PushBack(n)
}

func (p *PointHM) Get(x, y int) (v interface{}, ok bool) {
	l := p.getList(x, y)
	n := getNode(l, x, y)
	if n != nil {
		return n.v, true
	}
	return nil, false
}

func (p *PointHM) Remove(x, y int) (v interface{}, ok bool) {
	l := p.getList(x, y)
	for e := l.Front(); e != nil; e = e.Next() {
		n := e.Value.(*node)
		if (n.x == x) && (n.y == y) {
			l.Remove(e)
			return n.v, true
		}
	}
	return nil, false
}

func (p *PointHM) Walk(fn func(x, y int, v interface{})) {
	for _, ls := range p.matrix {
		for _, l := range ls {
			for e := l.Front(); e != nil; e = e.Next() {
				n := e.Value.(*node)
				fn(n.x, n.y, n.v)
			}
		}
	}
}

type node struct {
	x, y int
	v    interface{}
}

func getNode(l *list.List, x, y int) *node {
	for e := l.Front(); e != nil; e = e.Next() {
		n := e.Value.(*node)
		if (n.x == x) && (n.y == y) {
			return n
		}
	}
	return nil
}
