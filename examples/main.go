package main

import (
	"fmt"
	"rdp"

	"gonum.org/v1/plot/plotter"
)

const Threshold = 1

// ToPoints converts gonum plotter.XYs to point structs.
func ToPoints(xys plotter.XYs) []rdp.Point {
	points := make([]rdp.Point, 0, len(xys))
	for i := range xys {
		points = append(points, rdp.Point{X: xys[i].X, Y: xys[i].Y})
	}

	return points
}

// ToXYs converts point structs to gonum plotter.XYs.
func ToXYs(points []rdp.Point) plotter.XYs {
	xys := make(plotter.XYs, len(points))
	for i := range points {
		xys[i].X = points[i].X
		xys[i].Y = points[i].Y
	}

	return xys
}

func main() {
	xys := RandomXYs(200, 0.5)
	simXYs := ToXYs(rdp.SimplifyPath(ToPoints(xys), Threshold))
	fmt.Println("Saving plot...")
	if err := SavePlot(xys, simXYs); err != nil {
		panic(err)
	}
}
