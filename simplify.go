package main

import "fmt"

func PathSimplify(points []Point, ep float64) []Point {
	if len(points) <= 2 {
		return []Point{points[0]}
	}

	l := Line{Start: points[0], End: points[len(points)-1]}

	idx, maxDist := seekMostDistantPoint(l, points)
	fmt.Println(idx, maxDist)

	return nil
}

func seekMostDistantPoint(l Line, points []Point) (idx int, maxDist float64) {
	for i := 0; i < len(points); i++ {
		d := l.DistanceToPoint(points[i])
		if d > maxDist {
			maxDist = d
			idx = i
		}
	}

	return idx, maxDist
}
