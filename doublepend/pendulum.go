package main

type Pendulum struct {
	Mass     float64 `json:"mass"`
	Length   float64 `json:"length"`
	Theta    float64 `json:"theta"`
	Velocity float64 `json:"velocity"`
}
