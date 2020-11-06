package main

import (
	"math"
)

type DoublePendulum [2]Pendulum

func getDPCoords(dp *DoublePendulum) (x1, y1, x2, y2 float64) {

	var (
		p1 = &(dp[0])
		p2 = &(dp[1])
	)

	var (
		r1 = p1.Length
		r2 = p2.Length

		a1 = p1.Theta
		a2 = p2.Theta
	)

	sin1, cos1 := math.Sincos(a1)

	x1 = r1 * sin1
	y1 = r1 * cos1

	sin2, cos2 := math.Sincos(a2)

	x2 = x1 + r2*sin2
	y2 = y1 + r2*cos2

	return
}

func nextStep(dp *DoublePendulum, deltaTime float64) {

	var (
		p1 = &(dp[0])
		p2 = &(dp[1])
	)

	var (
		m1 = p1.Mass
		m2 = p2.Mass

		r1 = p1.Length
		r2 = p2.Length

		a1 = p1.Theta
		a2 = p2.Theta

		a1_v = p1.Velocity
		a2_v = p2.Velocity

		dt = deltaTime
	)

	//---------------------------------------------------

	a1_a, a2_a := calc(m1, m2, r1, r2, a1, a2, a1_v, a2_v)

	a1_v += a1_a * dt
	a2_v += a2_a * dt

	a1 += a1_v * dt
	a2 += a2_v * dt

	if false {
		a1_v *= 0.99
		a2_v *= 0.99
	}

	//---------------------------------------------------

	p1.Theta = a1
	p2.Theta = a2

	p1.Velocity = a1_v
	p2.Velocity = a2_v
}

// type Thetas struct {
// 	theta    float64 // theta
// 	d_theta  float64 // theta'
// 	dd_theta float64 // theta"
// }

// https://github.com/myphysicslab/myphysicslab/blob/master/src/sims/pendulum/DoublePendulumSim.js

// accelerations: a1_a, a2_a
func calc(m1, m2 float64, r1, r2 float64, a1, a2 float64, a1_v, a2_v float64) (a1_a, a2_a float64) {

	const g = 9.81

	den := (2*m1 + m2 - m2*math.Cos(2*(a1-a2)))

	tmp1 := -g * (2*m1 + m2) * math.Sin(a1)
	tmp2 := -g * m2 * math.Sin(a1-2*a2)
	tmp3 := -2 * math.Sin(a1-a2) * m2
	tmp4 := a2_v*a2_v*r2 + a1_v*a1_v*r1*math.Cos(a1-a2)

	a1_a = (tmp1 + tmp2 + tmp3*tmp4) / (r1 * den)

	tmp1 = 2 * math.Sin(a1-a2)
	tmp2 = a1_v * a1_v * r1 * (m1 + m2)
	tmp3 = g * (m1 + m2) * math.Cos(a1)
	tmp4 = a2_v * a2_v * r2 * m2 * math.Cos(a1-a2)

	a2_a = (tmp1 * (tmp2 + tmp3 + tmp4)) / (r2 * den)

	return
}
