package main

type Point2f struct {
	X, Y float64
}

type OptPoint2f struct {
	Present bool
	Value   Point2f
}
