package main

import (
	"fmt"

	"github.com/gotk3/gotk3/cairo"
)

type EventHandler interface {
	Resize(width, height int)
	Draw(c *cairo.Context)
	EventKey(keyCode uint16)
}

//------------------------------------------------------------------------------
type dummyEventHandler struct{}

var _ EventHandler = dummyEventHandler{}

func (dummyEventHandler) Resize(width, height int) {
	fmt.Printf("size: (%d,%d)\n", width, height)
}

func (dummyEventHandler) Draw(*cairo.Context) {

}

func (dummyEventHandler) EventKey(keyCode uint16) {

}

//------------------------------------------------------------------------------
type engineEventHandler struct {
	number int
	engine *Engine
}

var _ EventHandler = &engineEventHandler{}

func (p *engineEventHandler) Resize(width, height int) {
	p.engine._Resize(width, height)
}

func (p *engineEventHandler) Draw(c *cairo.Context) {
	p.engine._Draw(c)
}

func (p *engineEventHandler) EventKey(keyCode uint16) {
	switch keyCode {
	case 27: // 'r', 'R'
		{
			r := newRandNow()
			samples := randSamples(r, p.number)
			p.engine.SetSamples(samples)
		}
	case 65: // Space
		{
			p.engine.Pause()
		}
	case 39:
		{
			err := p.engine.SaveFileDP()
			checkError(err)
		}
	default:
		//fmt.Println("key code ", keyCode)
	}
}
