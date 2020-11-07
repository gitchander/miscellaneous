package main

import (
	"fmt"

	"github.com/gotk3/gotk3/cairo"
)

type Drawer interface {
	Resize(width, height int)
	Draw(c *cairo.Context)
}

type dummyDrawer struct{}

var _ Drawer = dummyDrawer{}

func (dummyDrawer) Resize(width, height int) {
	fmt.Printf("size: (%d,%d)\n", width, height)
}

func (dummyDrawer) Draw(*cairo.Context) {

}
