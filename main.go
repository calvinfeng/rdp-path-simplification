package main

import (
	"gonum.org/v1/plot/plotter"
)

const Threshold = 1

// ToPoints converts gonum plotter.XYs to point structs.
func ToPoints(xys plotter.XYs) []Point {
	points := make([]Point, 0, len(xys))
	for i := range xys {
		points = append(points, Point{xys[i].X, xys[i].Y})
	}

	return points
}

// ToXYs converts point structs to gonum plotter.XYs.
func ToXYs(points []Point) plotter.XYs {
	xys := make(plotter.XYs, len(points))
	for i := range points {
		xys[i].X = points[i].X
		xys[i].Y = points[i].Y
	}

	return xys
}

func main() {
	xys := RandomXYs(200, 0.5)
	simXYs := ToXYs(SimplifyPath(ToPoints(xys), Threshold))
	if err := SavePlot(xys, simXYs); err != nil {
		panic(err)
	}
}
