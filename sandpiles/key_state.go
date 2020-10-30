package main

import (
	"github.com/hajimehoshi/ebiten"
)

type KeyState struct {
	key ebiten.Key

	press     bool
	pressPrev bool
}

func NewKeyState(key ebiten.Key) *KeyState {
	return &KeyState{key: key}
}

func (ks *KeyState) Update() {
	ks.pressPrev = ks.press
	ks.press = ebiten.IsKeyPressed(ks.key)
}

func (ks *KeyState) KeyDown() bool {
	return ks.press && !ks.pressPrev
}

func (ks *KeyState) KeyUp() bool {
	return !ks.press && ks.pressPrev
}
