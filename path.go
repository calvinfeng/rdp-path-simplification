package main

import "math"

type Point struct {
	X float64
	Y float64
}

type Line struct {
	Start Point
	End   Point
}

// DistanceToPoint returns the perpendicular distance of a point to the line.
func (l Line) DistanceToPoint(pt Point) float64 {
	a, b, c := l.Coefficients()
	return math.Abs(a*pt.X+b*pt.Y+c) / math.Sqrt(a*a+b*b)
}

// Coefficients returns the three coefficients that define a line.
// A line can represent by the following equation.
//
// ax + by + c = 0
//
func (l Line) Coefficients() (a, b, c float64) {
	a = l.Start.Y - l.End.Y
	b = l.End.X - l.Start.X
	c = l.Start.X*l.End.Y - l.End.X*l.Start.Y

	return a, b, c
}
