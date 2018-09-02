package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func SavePlot(points plotter.XYs) error {
	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = "Visualize Path"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p, "Path", points)
	if err != nil {
		return err
	}

	return p.Save(10*vg.Inch, 5*vg.Inch, "path.png")
}

func RandomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}

		pts[i].Y = (pts[i].X + 10*rand.Float64()) * 1000
	}

	return pts
}
